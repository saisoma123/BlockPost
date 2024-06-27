package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/keeper"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
	"github.com/stretchr/testify/require"
)

func setupQueryServer(t testing.TB) (keeper.Keeper, types.QueryServer, sdk.Context) {
	k, ctx := keepertest.BlockpostKeeper(t)
	return k, keeper.NewQueryServerImpl(k), ctx
}

func TestQueryServer_Message(t *testing.T) {
	k, queryServer, sdkCtx := setupQueryServer(t)

	// Add a message to query later
	creator := "cosmos1v9jxgu33kfsgr5"
	message := "Hello, blockchain!"

	// Ensure the message is added
	id, err := k.AddMessage(sdkCtx, creator, message)
	require.NoError(t, err)

	// Query the message by ID
	req := &types.QueryMessageRequest{Id: id}
	resp, err := queryServer.Message(sdk.WrapSDKContext(sdkCtx), req)
	require.NoError(t, err)
	require.Equal(t, message, resp.Message)
}

func TestQueryServer_Messages(t *testing.T) {
	k, queryServer, sdkCtx := setupQueryServer(t)

	// Add some messages to query later
	creator := "cosmos1v9jxgu33kfsgr5"
	messages := []string{"Hello, blockchain!", "Another message"}

	for _, msg := range messages {
		_, err := k.AddMessage(sdkCtx, creator, msg)
		require.NoError(t, err)
	}

	// Query all messages
	req := &types.QueryAllMessagesRequest{}
	resp, err := queryServer.Messages(sdk.WrapSDKContext(sdkCtx), req)
	require.NoError(t, err)
	require.Equal(t, len(messages), len(resp.Messages))
	for i, msg := range messages {
		require.Equal(t, msg, resp.Messages[i])
	}
}
