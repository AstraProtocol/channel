package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"

	"channel/x/channel/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Senderwithdrawhashlock(goCtx context.Context, msg *types.MsgSenderwithdrawhashlock) (*types.MsgSenderwithdrawhashlockResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	val, found := k.Keeper.GetFwdcommit(ctx, msg.Transferindex)
	if !found {
		return nil, errors.New("commitment is not existing")
	}

	if val.Sender != msg.To {
		return nil, fmt.Errorf("not matching receiver address! expected:", val.Sender)
		//return nil, errors.New("not matching receiver address!")
	}

	hash := sha256.Sum256([]byte(msg.Secret))
	if val.Hashcodehtlc != base64.StdEncoding.EncodeToString(hash[:]) {
		return nil, errors.New("Wrong hash !")
	}

	to, err := sdk.AccAddressFromBech32(msg.To)
	if err != nil {
		return nil, err
	}

	err = k.Keeper.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, to, sdk.Coins{*val.Coin})
	if err != nil {
		return nil, err
	}

	k.Keeper.RemoveCommitment(ctx, msg.Transferindex)

	return &types.MsgSenderwithdrawhashlockResponse{}, nil
}
