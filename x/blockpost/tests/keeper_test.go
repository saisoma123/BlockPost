package keeper_test

import (
	"fmt"
	"testing"

	testutil "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestAddMessage(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data
	creator := "Surya Somayyajula"
	message := "This is very sensitive information"

	// Call the addMessage method
	messageID, err := k.AddMessage(ctx, creator, message)
	require.NoError(t, err)
	require.NotEmpty(t, messageID)
	fmt.Println(messageID)
	// Retrieve the message and verify it was stored correctly
	storedMessage, err := k.GetMessage(ctx, messageID)
	require.NoError(t, err)
	require.Equal(t, message, storedMessage)
	fmt.Println(storedMessage)
}

func TestGetAllMessages(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data
	creator := "Surya Somayyajula"
	messages := []string{"Nightwing!", "Batman"}
	// Add messages

	for _, msg := range messages {
		_, err := k.AddMessage(ctx, creator, msg)
		require.NoError(t, err)
	}

	// Retrieve all messages
	retrievedMessages, err := k.GetAllMessages(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t, messages, retrievedMessages)

}
