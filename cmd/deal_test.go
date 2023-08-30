package cmd

import (
	"context"
	"testing"
	"time"

	"github.com/data-preservation-programs/singularity/cmd/cliutil"
	"github.com/data-preservation-programs/singularity/handler/deal"
	"github.com/data-preservation-programs/singularity/model"
	"github.com/data-preservation-programs/singularity/util/testutil"
	"github.com/gotidy/ptr"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func swapDealHandler(mockHandler deal.Handler) func() {
	actual := deal.Default
	deal.Default = mockHandler
	return func() {
		deal.Default = actual
	}
}

func TestSendDealHandler_TooFewArgs(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	out, stderr, err := runner.Run(ctx, "singularity deal send-manual")
	require.ErrorIs(t, err, cliutil.ErrIncorrectNArgs)
	require.Contains(t, out, "USAGE:")
	require.Empty(t, stderr)
}

func TestSendDealHandler(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockDeal)
	defer swapDealHandler(mockHandler)()
	mockHandler.On("SendManualHandler", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(&model.Deal{
		ID:               1,
		CreatedAt:        time.Time{},
		UpdatedAt:        time.Time{},
		State:            "proposed",
		Provider:         "f01",
		ProposalID:       "proposal_id",
		Label:            "label",
		PieceCID:         model.CID(testutil.TestCid),
		PieceSize:        1024,
		StartEpoch:       1001,
		EndEpoch:         1999,
		SectorStartEpoch: 1500,
		Price:            "0",
		Verified:         true,
		ClientID:         "client_id",
	}, nil)
	_, _, err := runner.Run(ctx, "singularity deal send-manual client provider piece_cid 1024")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --verbose deal send-manual client provider piece_cid 1024")
	require.NoError(t, err)

}

func TestListDealHandler(t *testing.T) {
	runner := Runner{}
	defer runner.Save(t)
	ctx := context.Background()
	mockHandler := new(MockDeal)
	defer swapDealHandler(mockHandler)()
	mockHandler.On("ListHandler", mock.Anything, mock.Anything, mock.Anything).Return([]model.Deal{
		{
			ID:               1,
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			DealID:           ptr.Of(uint64(100)),
			State:            "active",
			Provider:         "f01",
			ProposalID:       "proposal_id",
			Label:            "label",
			PieceCID:         model.CID(testutil.TestCid),
			PieceSize:        1024,
			StartEpoch:       1001,
			EndEpoch:         1999,
			SectorStartEpoch: 1500,
			Price:            "0",
			Verified:         true,
			ScheduleID:       ptr.Of(uint32(5)),
			ClientID:         "client_id",
		},
		{
			ID:               2,
			CreatedAt:        time.Time{},
			UpdatedAt:        time.Time{},
			State:            "proposed",
			Provider:         "f01",
			ProposalID:       "proposal_id_2",
			Label:            "label_2",
			PieceCID:         model.CID(testutil.TestCid),
			PieceSize:        1024,
			StartEpoch:       1011,
			EndEpoch:         2011,
			SectorStartEpoch: 1600,
			Price:            "0",
			Verified:         false,
			ScheduleID:       ptr.Of(uint32(5)),
			ClientID:         "client_id",
		},
	}, nil)
	_, _, err := runner.Run(ctx, "singularity deal list --preparation 1 --source source --schedule 5 --provider f01 --state active")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --verbose deal list --preparation 1 --source source --schedule 5 --provider f01 --state active")
	require.NoError(t, err)

	_, _, err = runner.Run(ctx, "singularity --json deal list --preparation 1 --source source --schedule 5 --provider f01 --state active")
	require.NoError(t, err)

}
