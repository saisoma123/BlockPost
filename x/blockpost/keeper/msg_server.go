package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/saisoma123/BlockPost/x/blockpost/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return msgServer{Keeper: keeper}
}

// Calls the addMessage keeper function and validates transactions
func (m msgServer) BlockPostMessage(ctx context.Context, msg *types.MsgBlockPostMessage) (*types.MsgBlockPostMessageResponse, error) {
	sdkCtx := sdk.UnwrapSDKContext(ctx)

	// Validates incoming message through creator and message strings
	if err := msg.ValidateBasic(); err != nil {
		fmt.Println("Validation failed!")
		return nil, err
	}

	// Adds message to KVStore and checks for any errors in processing
	messageID, err := m.Keeper.AddMessage(sdk.UnwrapSDKContext(ctx), msg.Creator, msg.Message)
	if err != nil {
		return nil, errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "failed to add message")
	}

	// Emits event for logging history
	sdkCtx.EventManager().EmitEvents(sdk.Events{
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, types.ModuleName),
			sdk.NewAttribute(sdk.AttributeKeyAction, "BlockPostMessage"),
			sdk.NewAttribute(sdk.AttributeKeySender, msg.Creator),
			sdk.NewAttribute("message_id", messageID),
		),
	})

	return &types.MsgBlockPostMessageResponse{MessageId: messageID}, nil
}

var _ types.MsgServer = msgServer{}
