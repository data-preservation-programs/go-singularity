package deal

import (
	"context"
	"testing"

	"github.com/data-preservation-programs/singularity/database"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/replication"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockDealMaker struct {
	mock.Mock
}

func (m *MockDealMaker) MakeDeal(ctx context.Context, walletObj model.Wallet, car model.Car, dealConfig replication.DealConfig) (*model.Deal, error) {
	args := m.Called(ctx, walletObj, car, dealConfig)
	return args.Get(0).(*model.Deal), args.Error(1)
}

var proposal = Proposal{
	HTTPHeaders:   []string{"a=b"},
	URLTemplate:   "http://127.0.0.1/piece",
	RootCID:       "bafy2bzaceczlclcg4notjmrz4ayenf7fi4mngnqbgjs27r3resyhzwxjnviay",
	Verified:      true,
	IPNI:          true,
	KeepUnsealed:  true,
	StartDelay:    "24h",
	Duration:      "2400h",
	ClientAddress: "f01000",
	ProviderID:    "f01001",
	PieceCID:      "baga6ea4seaqdyupo27fj2fk2mtefzlxvrbf6kdi4twdpccdzbyqrbpsvfsh5ula",
	PieceSize:     "32GiB",
	FileSize:      1000,
}

func TestSendManualHandler_WalletNotFound(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f09999",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	_, err = SendManualHandler(db, ctx, mockDealMaker, proposal)
	require.ErrorContains(t, err, "client address not found")
}

func TestSendManualHandler_InvalidPieceCID(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.PieceCID = "bad"
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "invalid piece CID")
}

func TestSendManualHandler_InvalidPieceCID_NOTCOMMP(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.PieceCID = proposal.RootCID
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "piece CID must be commp")
}

func TestSendManualHandler_InvalidPieceSize(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.PieceSize = "aaa"
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "invalid piece size")
}

func TestSendManualHandler_InvalidPieceSize_NotPowerOfTwo(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.PieceSize = "31GiB"
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "piece size must be a power of 2")
}

func TestSendManualHandler_InvalidRootCID(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.RootCID = "xxxx"
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "invalid root CID")
}

func TestSendManualHandler_InvalidDuration(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.Duration = "xxxx"
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "invalid duration")
}

func TestSendManualHandler_InvalidStartDelay(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	badProposal := proposal
	badProposal.StartDelay = "xxxx"
	_, err = SendManualHandler(db, ctx, mockDealMaker, badProposal)
	require.ErrorContains(t, err, "invalid start delay")
}

func TestSendManualHandler(t *testing.T) {
	wallet := model.Wallet{
		ID:      "f01000",
		Address: "f10000",
	}

	db, err := database.OpenInMemory()
	require.NoError(t, err)
	err = db.Create(&wallet).Error
	require.NoError(t, err)
	ctx := context.Background()

	mockDealMaker := new(MockDealMaker)
	mockDealMaker.On("MakeDeal", ctx, wallet, mock.Anything, mock.Anything).Return(&model.Deal{}, nil)
	resp, err := SendManualHandler(db, ctx, mockDealMaker, proposal)
	require.NoError(t, err)
	require.NotNil(t, resp)
}
