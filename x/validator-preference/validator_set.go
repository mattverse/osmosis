package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/osmosis-labs/osmosis/v12/x/validator-preference/types"
)

// GetValAddrAndVal checks if the validator address is valid and the validator provided exists on chain.
func (k Keeper) GetValAddrAndVal(ctx sdk.Context, valOperAddress string) (sdk.ValAddress, stakingtypes.Validator, error) {
	valAddr, err := sdk.ValAddressFromBech32(valOperAddress)
	if err != nil {
		return nil, stakingtypes.Validator{}, fmt.Errorf("validator address not formatted")
	}

	validator, found := k.stakingKeeper.GetValidator(ctx, valAddr)
	if !found {
		return nil, stakingtypes.Validator{}, fmt.Errorf("validator not found %s", validator)
	}

	return valAddr, validator, nil
}

func (k Keeper) ValidatePreferences(ctx sdk.Context, preferences []types.ValidatorPreference) error {
	for _, val := range preferences {
		_, _, err := k.GetValAddrAndVal(ctx, val.ValOperAddress)
		if err != nil {
			return err
		}
	}
	return nil
}
