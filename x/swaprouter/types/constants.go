package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	// MultihopSwapFeeMultiplierForOsmoPools if a swap fees multiplier for trades consists of just two OSMO pools during a single transaction.
	MultihopSwapFeeMultiplierForOsmoPools = sdk.NewDecWithPrec(5, 1) // 0.5
)
