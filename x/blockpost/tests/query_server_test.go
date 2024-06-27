package keeper_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
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

func TestQueryServer_EmptyRequest(t *testing.T) {
	_, queryServer, sdkCtx := setupQueryServer(t)

	// Query with an empty request
	req := &types.QueryMessageRequest{}
	_, err := queryServer.Message(sdk.WrapSDKContext(sdkCtx), req)
	require.Error(t, err)
}

func TestQueryServer_NonExistentMessage(t *testing.T) {
	_, queryServer, sdkCtx := setupQueryServer(t)

	// Query a non-existent message ID
	req := &types.QueryMessageRequest{Id: "nonExistentID"}
	_, err := queryServer.Message(sdk.WrapSDKContext(sdkCtx), req)
	require.Error(t, err)
}

func TestQueryServer_ManyMessages(t *testing.T) {
	k, queryServer, sdkCtx := setupQueryServer(t)

	// Add many messages to test pagination
	creator := "cosmos1v9jxgu33kfsgr5"
	numMessages := 50
	messages := make([]string, numMessages)
	for i := 0; i < numMessages; i++ {
		msg := fmt.Sprintf("Message #%d", i)
		messages[i] = msg
		_, err := k.AddMessage(sdkCtx, creator, msg)
		require.NoError(t, err)
	}

	// Query all messages without pagination
	req := &types.QueryAllMessagesRequest{Pagination: &query.PageRequest{
		Limit:  10,
		Offset: 0,
	}}
	resp, err := queryServer.Messages(sdk.WrapSDKContext(sdkCtx), req)
	require.NoError(t, err)
	require.Equal(t, len(messages), len(resp.Messages))
	require.ElementsMatch(t, messages, resp.Messages)
}

func TestQueryServer_ManyMessagesWithDifferentCreators(t *testing.T) {
	k, queryServer, sdkCtx := setupQueryServer(t)

	// Add many messages from different creators
	creators := []string{
		"cosmos1v9jxgu33kfsgr5", "cosmos1d9v5e4k8g7e8u4k9f7g6u5k4", "cosmos1k2n3v4u5t6g7e8v9j0f1g2h3",
		"cosmos1e7u4k9f7g6u5k4d9v5e4k8g7", "cosmos1j0f1g2h3k2n3v4u5t6g7e8v9",
		"cosmos1f7g6u5k4d9v5e4k8g7e8u4k9", "cosmos1v4u5t6g7e8v9j0f1g2h3k2n3",
		"cosmos1u5k4d9v5e4k8g7e8v9j0f1g2", "cosmos1k9f7g6u5t6g7e8v9j0f1g2h3",
		"cosmos1g7e8u4k9f7g6u5k4d9v5e4k8", "cosmos1h3k2n3v4u5t6g7e8v9j0f1g2",
		"cosmos1g2h3k2n3v4u5t6g7e8v9j0f1", "cosmos1v9jxgu33kfsgr5d9v5e4k8g7",
		"cosmos1d9v5e4k8g7e8u4k9f7g6u5k4", "cosmos1k2n3v4u5t6g7e8v9j0f1g2h3",
	}
	numMessages := len(creators) * 10 // 10 messages per creator
	messages := make([]string, numMessages)
	for i, creator := range creators {
		for j := 0; j < 10; j++ {
			msg := fmt.Sprintf("Message #%d from %s", j, creator)
			messages[i*10+j] = msg
			_, err := k.AddMessage(sdkCtx, creator, msg)
			require.NoError(t, err)
		}
	}

	req := &types.QueryAllMessagesRequest{Pagination: &query.PageRequest{
		Limit:  10,
		Offset: 0,
	}}
	resp, err := queryServer.Messages(sdk.WrapSDKContext(sdkCtx), req)
	require.NoError(t, err)
	require.Equal(t, len(messages), len(resp.Messages))
	require.ElementsMatch(t, messages, resp.Messages)
}
