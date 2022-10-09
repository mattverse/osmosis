package keeper

import (
	"context"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/osmosis-labs/osmosis/v12/x/validator-preference/types"
)

type msgServer struct {
	keeper *Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper *Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (server msgServer) SetValidatorSetPreference(goCtx context.Context, msg *types.MsgSetValidatorSetPreference) (*types.MsgSetValidatorSetPreferenceResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// new user preference
	preferences := msg.Preferences

	delegatorAddr, err := sdk.AccAddressFromBech32(msg.Delegator)
	if err != nil {
		return nil, err
	}

	// check that the user account has funds to prevent spam create
	userAccBalance := server.keeper.bankKeeper.GetAllBalances(ctx, delegatorAddr)
	if len(userAccBalance) == 0 {
		return nil, fmt.Errorf("The user doesnot have any tokens")
	}

	// check if a user already have a validator-set created
	existingValidators, found := server.keeper.GetValidatorSetPreference(ctx, msg.Delegator)
	if found {
		// check if the new preferences is the same as the existing preferences
		if reflect.DeepEqual(preferences, existingValidators.Preferences) {
			return nil, fmt.Errorf("The preferences are the same")
		}
	}

	// checks that all the validators exist on chain
	err = server.keeper.ValidatePreferences(ctx, preferences)
	if err != nil {
		return nil, err
	}

	// create/update the validator-set based on what user provides
	setMsg := types.ValidatorSetPreferences{
		Preferences: msg.Preferences,
	}

	server.keeper.SetValidatorSetPreferences(ctx, msg.Delegator, setMsg)
	return &types.MsgSetValidatorSetPreferenceResponse{}, nil
}

func (server msgServer) DelegateToValidatorSet(goCtx context.Context, msg *types.MsgDelegateToValidatorSet) (*types.MsgDelegateToValidatorSetResponse, error) {
	return &types.MsgDelegateToValidatorSetResponse{}, nil
}

func (server msgServer) UndelegateFromValidatorSet(goCtx context.Context, msg *types.MsgUndelegateFromValidatorSet) (*types.MsgUndelegateFromValidatorSetResponse, error) {
	return &types.MsgUndelegateFromValidatorSetResponse{}, nil
}

func (server msgServer) WithdrawDelegationRewards(goCtx context.Context, msg *types.MsgWithdrawDelegationRewards) (*types.MsgWithdrawDelegationRewardsResponse, error) {
	return &types.MsgWithdrawDelegationRewardsResponse{}, nil
}
