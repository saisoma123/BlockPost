package keeper_test

import (
	"fmt"
	"testing"

	testutil "github.com/saisoma123/BlockPost/testutil/keeper"
	"github.com/stretchr/testify/require"
)

func TestBasicAddMessage(t *testing.T) {
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

func TestSpamMessages(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data, creates many messages and stores in arrays
	creator := "cosmos1v9jxgu33kfsgr5"
	numMessages := 10000 // Number of messages to add
	messages := make([]string, numMessages)
	for i := 0; i < numMessages; i++ {
		messages[i] = fmt.Sprintf("Message #%d", i)
	}

	// Add messages
	for _, msg := range messages {
		_, err := k.AddMessage(ctx, creator, msg)
		require.NoError(t, err)
	}

	// Retrieve all messages
	retrievedMessages, err := k.GetAllMessages(ctx)
	require.NoError(t, err)

	// Ensure all messages are retrieved
	require.ElementsMatch(t, messages, retrievedMessages)
	require.Equal(t, len(messages), len(retrievedMessages))

	fmt.Println("All messages added and retrieved successfully")
}

func TestGetNonExistentMessage(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data
	messageID := "nonExistentID"

	// Try to retrieve a non-existent message
	_, err := k.GetMessage(ctx, messageID)
	require.Error(t, err)
}

func TestGetAllMessagesEmptyStore(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Retrieve all messages from an empty store
	retrievedMessages, err := k.GetAllMessages(ctx)
	require.NoError(t, err)
	require.Empty(t, retrievedMessages)
}

func TestAddMultipleMessages(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data
	creator := "cosmos1v9jxgu33kfsgr5"
	messages := []string{
		"This is the first message",
		"This is the second message",
		"This is the third message",
	}

	// Add messages
	var messageIDs []string
	for _, msg := range messages {
		messageID, err := k.AddMessage(ctx, creator, msg)
		require.NoError(t, err)
		messageIDs = append(messageIDs, messageID)
	}

	// Retrieve and verify each message
	for i, messageID := range messageIDs {
		storedMessage, err := k.GetMessage(ctx, messageID)
		require.NoError(t, err)
		require.Equal(t, messages[i], storedMessage)
	}
}
