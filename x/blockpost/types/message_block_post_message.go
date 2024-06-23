package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBlockPostMessage{}

func NewMsgBlockPostMessage(creator string, message string) *MsgBlockPostMessage {
	return &MsgBlockPostMessage{
		Creator: creator,
		Message: message,
	}
}

// Returns the module name
func (msg MsgBlockPostMessage) Route() string { return "blockpost" }

// Returns the message type within this module
func (msg MsgBlockPostMessage) Type() string { return "BlockPostMessage" }

// Error checking for potential invalid/missing sender addresses and messages
func (msg *MsgBlockPostMessage) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if msg.Creator == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "missing creator address (%s)", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.Creator); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid creator address: %s", err)
	}
	if msg.Message == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "missing message (%s)", err)
	}
	return nil
}

// Gets the bytes that need to be signed by the signers of this transaction
func (msg MsgBlockPostMessage) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// Gets all of the signers that need to sign this particular message
func (msg MsgBlockPostMessage) GetSigners() []sdk.AccAddress {
	creator, _ := sdk.AccAddressFromBech32(msg.Creator)
	return []sdk.AccAddress{creator}
}
