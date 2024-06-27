package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
	"github.com/stretchr/testify/require"
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
		{"LongMessage", "cosmos1v9jxgu33kfsgr5", string(make([]byte, 1000)), false},
		{"SpecialCharacters", "cosmos1v9jxgu33kfsgr5", "Hello, @blockchain!", false},
		{"SQLInjection", "cosmos1v9jxgu33kfsgr5", "DROP TABLE users;", false},
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

func TestMsgServer_BlockPostMessageStore_DuplicateMessages(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)

	msg := &types.MsgBlockPostMessage{
		Creator: "cosmos1v9jxgu33kfsgr5",
		Message: "Hello, blockchain!",
	}

	// Add the same message twice
	_, err := msgServer.BlockPostMessage(ctx, msg)
	require.NoError(t, err, "Expected no error for the first message")

	_, err = msgServer.BlockPostMessage(ctx, msg)
	require.NoError(t, err, "Expected no error for the duplicate message")
}

func TestMsgServer_BlockPostMessageStore_InvalidTypes(t *testing.T) {
	msgServer, ctx := setupMsgServer(t)

	tests := []struct {
		name string
		msg  *types.MsgBlockPostMessage
	}{
		{"InvalidAddress", &types.MsgBlockPostMessage{Creator: "invalid_address", Message: "Hello, blockchain!"}},
		{"EmptyMessage", &types.MsgBlockPostMessage{Creator: "cosmos1v9jxgu33kfsgr5", Message: ""}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := msgServer.BlockPostMessage(ctx, tt.msg)
			require.Error(t, err, "Expected error for test case: %s", tt.name)
		})
	}
}
