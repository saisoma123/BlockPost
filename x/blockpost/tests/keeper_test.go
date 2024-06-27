package keeper_test

import (
	"testing"

	testutil "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestAddMessage(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data
	creator := "cosmos1v9jxgu33kfsgr5"
	message := "Hello, blockchain!"

	// Call the addMessage method
	messageID, err := k.AddMessage(ctx, creator, message)
	require.NoError(t, err)
	require.NotEmpty(t, messageID)

	// Retrieve the message and verify it was stored correctly
	storedMessage, err := k.GetMessage(ctx, messageID)
	require.NoError(t, err)
	require.Equal(t, message, storedMessage)
}
