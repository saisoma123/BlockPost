package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlockpostKeeper(t)
	return keeper.NewMsgServerImpl(k), ctx
}

func TestMsgServer_BlockPostMessageStore(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)

	tests := []struct {
		name    string
		creator string
		message string
		wantErr bool
	}{
		{"ValidMessage", "cosmos1v9jxgu33kfsgr5", "Hello, blockchain!", false},
		{"EmptyCreator", "", "Hello, blockchain!", true},
		{"EmptyMessage", "cosmos1v9jxgu33kfsgr5", "", true},
	}

	for i, tt := range tests {
		// Print the test case name
		t.Logf("Running test case: %s", tt.name)

		// Create a new MsgBlockPostMessage based on the test case
		msg := &types.MsgBlockPostMessage{
			Creator: tt.creator,
			Message: tt.message,
		}

		// Call the BlockPostMessage method
		_, err := msgServer.BlockPostMessage(ctx, msg)

		// Check if the error status matches the expected result
		if (err != nil) != tt.wantErr {
			t.Errorf("Test case %d (%s) - BlockPostMessage() error = %v, wantErr %v", i, tt.name, err, tt.wantErr)
		}
	}
}
