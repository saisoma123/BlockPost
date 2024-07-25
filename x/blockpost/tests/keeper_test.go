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

func TestAddMessagesWithDifferentCreators(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data
	creators := []string{
		"cosmos1v9jxgu33kfsgr5",
		"cosmos1d9v5e4k8g7e8u4k9f7g6u5k4",
		"cosmos1a1b2c3d4e5f6g7h8i9j0k1l2",
		"cosmos1x1y2z3a4b5c6d7e8f9g0h1i2",
		"cosmos1p1q2r3s4t5u6v7w8x9y0z1a2",
		"cosmos1g1h2i3j4k5l6m7n8o9p0q1r2",
		"cosmos1m1n2o3p4q5r6s7t8u9v0w1x2",
		"cosmos1t1u2v3w4x5y6z7a8b9c0d1e2",
		"cosmos1b1c2d3e4f5g6h7i8j9k0l1m2",
		"cosmos1z1y2x3w4v5u6t7s8r9q0p1o2",
		"cosmos1q1w2e3r4t5y6u7i8o9p0a1s2",
		"cosmos1a2b3c4d5e6f7g8h9i0j1k2l3",
		"cosmos1u2v3w4x5y6z7a8b9c0d1e2f3",
		"cosmos1m2n3o4p5q6r7s8t9u0v1w2x3",
		"cosmos1g2h3i4j5k6l7m8n9o0p1q2r3",
	}

	messages := []string{
		"Message from first creator",
		"Message from second creator",
		"Message from third creator",
		"Message from fourth creator",
		"Message from fifth creator",
		"Message from sixth creator",
		"Message from seventh creator",
		"Message from eighth creator",
		"Message from ninth creator",
		"Message from tenth creator",
		"Message from eleventh creator",
		"Message from twelfth creator",
		"Message from thirteenth creator",
		"Message from fourteenth creator",
		"Message from fifteenth creator",
	}

	// Add messages
	for i, creator := range creators {
		_, err := k.AddMessage(ctx, creator, messages[i])
		require.NoError(t, err)
	}

	// Retrieve all messages
	retrievedMessages, err := k.GetAllMessages(ctx)
	require.NoError(t, err)
	require.ElementsMatch(t, messages, retrievedMessages)
}

func TestAddAndGetMessagesFromMultipleCreators(t *testing.T) {
	k, ctx := testutil.BlockpostKeeper(t)

	// Test data: 15 creators with 15 messages each
	creators := []string{
		"cosmos1v9jxgu33kfsgr5", "cosmos1d9v5e4k8g7e8u4k9f7g6u5k4", "Bob", "Alice",
		"cosmos1a4f3g6k8u9e7w4k6g8t7h9", "cosmos1p3d7f4g9k8s7w3k5h9u6t8", "cosmos1m3f6g4k8j7l9w2k4h8t9r6",
		"cosmos1u5g4j7k9l8p3k6t5m9f8r7", "cosmos1w2k8g3f7j9l4p5k6t8u9r6", "cosmos1z7f4k8j9l6p3k5t4m8g9r2",
		"cosmos1n3f6g8k7j9p5k4t8m6r2", "cosmos1t3f4k8j7l9p5k6g8u9r4", "cosmos1x8f3g7k9j4p6k5t7m9u2",
		"cosmos1y7f4k3g8j9l6p5k8t4m7r", "cosmos1q5f3g6k7j9p4k8t5m6u7",
	}
	numMessages := 15
	messages := make([][]string, len(creators))
	for i := range creators {
		messages[i] = make([]string, numMessages)
		for j := 0; j < numMessages; j++ {
			messages[i][j] = fmt.Sprintf("Message #%d from %s", j, creators[i])
		}
	}

	messageIDs := make([][]string, len(creators))
	for i := range creators {
		messageIDs[i] = make([]string, numMessages)
	}

	// Add messages
	for i, creator := range creators {
		for j, msg := range messages[i] {
			messageID, err := k.AddMessage(ctx, creator, msg)
			require.NoError(t, err)
			messageIDs[i][j] = messageID
		}
	}

	// Retrieve and verify messages
	for i := range creators {
		for j, messageID := range messageIDs[i] {
			storedMessage, err := k.GetMessage(ctx, messageID)
			require.NoError(t, err)
			require.Equal(t, messages[i][j], storedMessage)
		}
	}
}
