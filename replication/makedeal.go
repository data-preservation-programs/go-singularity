package replication

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication/internal/proposal110"
	"github.com/data-preservation-programs/singularity/replication/internal/proposal120"
	"github.com/data-preservation-programs/singularity/util"
	"github.com/filecoin-project/go-address"
	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/big"
	"github.com/google/uuid"
	"github.com/ipfs/go-cid"
	"github.com/ipfs/go-log/v2"
	"github.com/jsign/go-filsigner/wallet"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/peerstore"
	"github.com/libp2p/go-libp2p/core/protocol"
	"github.com/multiformats/go-multiaddr"
	"github.com/pkg/errors"
	"github.com/ybbus/jsonrpc/v3"
	"golang.org/x/exp/slices"
	"strings"
	"time"
)

const (
	StorageProposalV120 = "/fil/storage/mk/1.2.0"
	StorageProposalV111 = "/fil/storage/mk/1.1.1"
)

func TimeToEpoch(t time.Time) abi.ChainEpoch {
	return abi.ChainEpoch((t.Unix() - 1598306400) / 30)
}

func EpochToTime(epoch abi.ChainEpoch) time.Time {
	return time.Unix(int64(epoch*30+1598306400), 0)
}

// nolint: tagliatelle
type MinerInfo struct {
	PeerIDEncoded           string `json:"PeerID"`
	PeerID                  peer.ID
	MultiaddrsBase64Encoded []string `json:"Multiaddrs"`
	Multiaddrs              []multiaddr.Multiaddr
}

type DealProviderCollateralBound struct {
	Min string
	Max string
}

type DealMaker struct {
	lotusClient jsonrpc.RPCClient
	host        host.Host
	logger      *log.ZapEventLogger
}

func NewDealMaker(lotusURL string, lotusToken string, libp2p host.Host) (*DealMaker, error) {
	lotusClient := util.NewLotusClient(lotusURL, lotusToken)

	return &DealMaker{
		lotusClient: lotusClient,
		host:        libp2p,
		logger:      log.Logger("deal_maker"),
	}, nil
}

func (d DealMaker) GetProviderInfo(ctx context.Context, provider string) (MinerInfo, error) {
	logger := log.Logger("deal_maker")

	logger.With("provider", provider).Debug("Getting miner info")
	minerInfo := new(MinerInfo)
	err := d.lotusClient.CallFor(ctx, minerInfo, "Filecoin.StateMinerInfo", provider, nil)
	if err != nil {
		return MinerInfo{}, errors.Wrap(err, "failed to get miner info")
	}

	logger.With("provider", provider, "minerinfo", minerInfo).Debug("Got miner info")
	minerInfo.Multiaddrs = make([]multiaddr.Multiaddr, len(minerInfo.MultiaddrsBase64Encoded))
	for i, addr := range minerInfo.MultiaddrsBase64Encoded {
		decoded, err := base64.StdEncoding.DecodeString(addr)
		if err != nil {
			return MinerInfo{}, errors.Wrap(err, "failed to decode multiaddr")
		}
		minerInfo.Multiaddrs[i], err = multiaddr.NewMultiaddrBytes(decoded)
		if err != nil {
			return MinerInfo{}, errors.Wrap(err, "failed to create multiaddr")
		}
	}
	minerInfo.PeerID, err = peer.Decode(minerInfo.PeerIDEncoded)
	if err != nil {
		return MinerInfo{}, errors.Wrap(err, "failed to decode peer id")
	}
	return *minerInfo, nil
}

func (d DealMaker) GetProtocols(ctx context.Context, minerInfo peer.AddrInfo) ([]protocol.ID, error) {
	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return nil, errors.Wrap(err, "failed to connect to miner")
	}

	protocols, err := d.host.Peerstore().GetProtocols(minerInfo.ID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get protocols")
	}

	return protocols, nil
}

func (d DealMaker) getMinCollateral(ctx context.Context, pieceSize int64, verified bool) (big.Int, error) {
	bound := new(DealProviderCollateralBound)
	err := d.lotusClient.CallFor(ctx, bound, "Filecoin.StateDealProviderCollateralBounds", pieceSize, verified, nil)
	if err != nil {
		return big.Int{}, errors.Wrap(err, "failed to get deal provider collateral bounds")
	}

	return big.FromString(bound.Min)
}

func (d DealMaker) makeDeal120(ctx context.Context,
	deal proposal110.ClientDealProposal,
	dealID uuid.UUID,
	car model.Car, schedule model.Schedule,
	minerInfo peer.AddrInfo) (*proposal120.DealResponse, error) {
	transfer := proposal120.Transfer{
		Size: uint64(car.FileSize),
	}
	url := strings.Replace(schedule.URLTemplate, "{PIECE_CID}", cid.MustParse(car.PieceCID).String(), 1)
	isOnline := url != ""
	if isOnline {
		transferParams := &proposal120.HttpRequest{URL: url}
		if len(schedule.HTTPHeaders) > 0 {
			transferParams.Headers = make(map[string]string)
			for _, header := range schedule.HTTPHeaders {
				sp := strings.Split(header, "=")
				if len(sp) != 2 {
					return nil, errors.Errorf("invalid http header %s", header)
				}
				transferParams.Headers[sp[0]] = sp[1]
			}
		}
		paramsBytes, err := json.Marshal(transferParams)
		if err != nil {
			return nil, errors.Wrap(err, "failed to serialize transfer params")
		}
		transfer.Type = "http"
		transfer.Params = paramsBytes
	}

	dealParams := proposal120.DealParams{
		DealUUID:           dealID,
		ClientDealProposal: deal,
		DealDataRoot:       cid.MustParse(car.RootCID),
		IsOffline:          !isOnline,
		Transfer:           transfer,
		RemoveUnsealedCopy: !schedule.KeepUnsealed,
		SkipIPNIAnnounce:   !schedule.AnnounceToIPNI,
	}

	stream, err := d.host.NewStream(ctx, minerInfo.ID, StorageProposalV120)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stream")
	}
	defer stream.Close()
	if deadline, ok := ctx.Deadline(); ok {
		err := stream.SetDeadline(deadline)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set stream deadline")
		}
		//nolint:errcheck
		defer stream.SetDeadline(time.Time{})
	}

	err = cborutil.WriteCborRPC(stream, &dealParams)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write deal params")
	}

	var resp proposal120.DealResponse
	err = cborutil.ReadCborRPC(stream, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read deal response")
	}

	return &resp, nil
}

func (d DealMaker) makeDeal111(ctx context.Context,
	deal proposal110.ClientDealProposal,
	car model.Car, schedule model.Schedule,
	minerInfo peer.AddrInfo) (*proposal110.SignedResponse, error) {
	proposal := proposal110.Proposal{
		FastRetrieval: schedule.KeepUnsealed,
		DealProposal:  &deal,
		Piece: &proposal110.DataRef{
			TransferType: proposal110.TTManual,
			Root:         cid.MustParse(car.RootCID),
			PieceCid:     &deal.Proposal.PieceCID,
			PieceSize:    deal.Proposal.PieceSize.Unpadded(),
		},
	}

	stream, err := d.host.NewStream(ctx, minerInfo.ID, StorageProposalV111)
	if err != nil {
		return nil, errors.Wrap(err, "failed to open stream")
	}
	defer stream.Close()
	if deadline, ok := ctx.Deadline(); ok {
		err = stream.SetDeadline(deadline)
		if err != nil {
			return nil, errors.Wrap(err, "failed to set stream deadline")
		}
		//nolint:errcheck
		defer stream.SetDeadline(time.Time{})
	}

	var resp proposal110.SignedResponse
	err = cborutil.WriteCborRPC(stream, &proposal)
	if err != nil {
		return nil, errors.Wrap(err, "failed to write deal params")
	}

	err = cborutil.ReadCborRPC(stream, &resp)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read deal response")
	}

	return &resp, nil
}

func (d DealMaker) MakeDeal(ctx context.Context, now time.Time, walletObj model.Wallet,
	car model.Car, schedule model.Schedule, minerInfo peer.AddrInfo) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.host.Peerstore().AddAddrs(minerInfo.ID, minerInfo.Addrs, peerstore.TempAddrTTL)
	if err := d.host.Connect(ctx, minerInfo); err != nil {
		return "", errors.Wrap(err, "failed to connect to miner")
	}

	protocols, err := d.host.Peerstore().GetProtocols(minerInfo.ID)
	if err != nil {
		return "", errors.Wrap(err, "failed to get supported protocol")
	}

	addr, err := address.NewFromString(walletObj.Address)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse wallet address")
	}
	pvd, err := address.NewFromString(schedule.Provider)
	if err != nil {
		return "", errors.Wrap(err, "failed to parse provider address")
	}
	label, err := proposal110.NewLabelFromString(cid.MustParse(car.RootCID).String())
	if err != nil {
		return "", errors.Wrap(err, "failed to parse label")
	}

	startEpoch := TimeToEpoch(now.Add(schedule.StartDelay))
	endEpoch := TimeToEpoch(now.Add(schedule.StartDelay + schedule.Duration))
	price := big.NewInt(int64(schedule.Price * 1e18 / schedule.Duration.Minutes() * 2 * float64(car.PieceSize) / (1 << 38)))
	verified := schedule.Verified
	pieceCID := cid.MustParse(car.PieceCID)
	pieceSize := abi.PaddedPieceSize(car.PieceSize)
	collateral, err := d.getMinCollateral(ctx, car.PieceSize, verified)
	if err != nil {
		return "", errors.Wrap(err, "failed to get min collateral")
	}
	proposal := proposal110.DealProposal{
		PieceCID:             pieceCID,
		PieceSize:            pieceSize,
		VerifiedDeal:         verified,
		Client:               addr,
		Provider:             pvd,
		Label:                label,
		StartEpoch:           startEpoch,
		EndEpoch:             endEpoch,
		StoragePricePerEpoch: price,
		ProviderCollateral:   collateral,
	}
	proposalBytes, err := cborutil.Dump(&proposal)
	if err != nil {
		return "", errors.Wrap(err, "failed to serialize deal proposal")
	}

	signature, err := wallet.WalletSign(walletObj.PrivateKey, proposalBytes)
	if err != nil {
		return "", errors.Wrap(err, "failed to sign deal proposal")
	}

	deal := proposal110.ClientDealProposal{
		Proposal:        proposal,
		ClientSignature: *signature,
	}

	if slices.Contains(protocols, StorageProposalV120) {
		dealID := uuid.New()
		resp, err := d.makeDeal120(ctx, deal, dealID, car, schedule, minerInfo)
		if err != nil {
			return "", errors.Wrap(err, "failed to make deal")
		}
		if resp.Accepted {
			return dealID.String(), nil
		}

		return "", errors.Errorf("deal rejected: %s", resp.Message)
	} else if slices.Contains(protocols, StorageProposalV111) {
		resp, err := d.makeDeal111(ctx, deal, car, schedule, minerInfo)
		if err != nil {
			return "", errors.Wrap(err, "failed to make deal")
		}

		return resp.Response.Proposal.String(), nil
	}

	return "", errors.New("no supported protocol found")
}
