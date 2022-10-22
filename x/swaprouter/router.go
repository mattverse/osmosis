package swaprouter

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v12/x/swaprouter/types"
)

// TODO: spec and tests
func (k Keeper) RouteExactAmountIn(
	ctx sdk.Context,
	sender sdk.AccAddress,
	routes []types.SwapAmountInRoute,
	tokenIn sdk.Coin,
	tokenOutMinAmount sdk.Int) (tokenOutAmount sdk.Int, err error) {
	isGamm := true

	if isGamm {
		return k.gammKeeper.MultihopSwapExactAmountIn(ctx, sender, routes, tokenIn, tokenOutMinAmount)
	}

	return k.concentratedKeeper.MultihopSwapExactAmountIn(ctx, sender, routes, tokenIn, tokenOutMinAmount)
}

// TODO: spec and tests
func (k Keeper) RouteExactAmountOut(ctx sdk.Context,
	sender sdk.AccAddress,
	routes []types.SwapAmountOutRoute,
	tokenInMaxAmount sdk.Int,
	tokenOut sdk.Coin) (tokenInAmount sdk.Int, err error) {
	isGamm := true

	if isGamm {
		return k.gammKeeper.MultihopSwapExactAmountOut(ctx, sender, routes, tokenInMaxAmount, tokenOut)
	}

	return k.concentratedKeeper.MultihopSwapExactAmountOut(ctx, sender, routes, tokenInMaxAmount, tokenOut)
}

func (k Keeper) RouteMulihopSwapExactAmountIn(
	ctx sdk.Context,
	sender sdk.AccAddress,
	routes []types.SwapAmountInRoute,
	tokenIn sdk.Coin,
	tokenOutMinAmount sdk.Int,
	isGamm bool,
) (tokenOutAmount sdk.Int, err error) {
	for i, route := range routes {
		swapFeeMultiplier := sdk.OneDec()
		if types.SwapAmountInRoutes(routes).IsOsmoRoutedMultihop() {
			swapFeeMultiplier = types.MultihopSwapFeeMultiplierForOsmoPools.Clone()
		}

		// To prevent the multihop swap from being interrupted prematurely, we keep
		// the minimum expected output at a very low number until the last pool
		_outMinAmount := sdk.NewInt(1)
		if len(routes)-1 == i {
			_outMinAmount = tokenOutMinAmount
		}

		tokenOutAmount, err := k.gammKeeper.SwapExactAmountIn(ctx, sender, route.PoolId, tokenIn, route.TokenOutDenom, _outMinAmount)
		if err != nil {
			return sdk.Int{}, err
		}

		// Chain output of current pool as the input for the next routed pool
		tokenIn = sdk.NewCoin(route.TokenOutDenom, tokenOutAmount)
	}
	return tokenOutAmount, err
}
