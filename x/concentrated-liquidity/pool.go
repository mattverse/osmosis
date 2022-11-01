package concentrated_liquidity

import (
	"errors"
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/v12/osmomath"
	cltypes "github.com/osmosis-labs/osmosis/v12/x/concentrated-liquidity/types"
	gammtypes "github.com/osmosis-labs/osmosis/v12/x/gamm/types"
)

// TODO: spec, tests, implementation
func (k Keeper) InitializePool(ctx sdk.Context, pool gammtypes.PoolI, creatorAddress sdk.AccAddress) error {
	panic("not implemented")
}

func (k Keeper) CreateNewConcentratedLiquidityPool(ctx sdk.Context, poolId uint64, poolAddress sdk.AccAddress, denom0, denom1 string, currSqrtPrice sdk.Dec, currTick sdk.Int) (Pool, error) {
	denom0, denom1, err := cltypes.OrderInitialPoolDenoms(denom0, denom1)
	if err != nil {
		return Pool{}, err
	}
	pool := Pool{
		Address:          poolAddress.String(),
		Id:               poolId,
		CurrentSqrtPrice: currSqrtPrice,
		CurrentTick:      currTick,
		Liquidity:        sdk.ZeroDec(),
		Token0:           denom0,
		Token1:           denom1,
	}

	k.setPoolById(ctx, poolId, pool)

	return pool, nil
}

// GetPool returns a pool with a given id.
func (k Keeper) GetPool(ctx sdk.Context, poolId uint64) (gammtypes.PoolI, error) {
	return nil, errors.New("not implemented")
}

// priceToTick takes a price and returns the corresponding tick index
func priceToTick(price sdk.Dec) sdk.Int {
	logOfPrice := osmomath.BigDecFromSDKDec(price).LogBase2()
	logInt := osmomath.NewDecWithPrec(10001, 4)
	tick := logOfPrice.Quo(logInt.LogBase2())
	return tick.SDKDec().TruncateInt()
}

func (p Pool) GetAddress() sdk.AccAddress {
	return sdk.AccAddress{}
}

func (p Pool) String() string {
	return ""
}

func (p Pool) GetId() uint64 {
	return 0
}
func (p Pool) GetSwapFee(ctx sdk.Context) sdk.Dec {
	return sdk.Dec{}
}
func (p Pool) GetExitFee(ctx sdk.Context) sdk.Dec {
	return sdk.Dec{}
}
func (p Pool) IsActive(ctx sdk.Context) bool {
	return true
}
func (p Pool) GetTotalShares() sdk.Int {
	return sdk.Int{}
}

func (k Keeper) SwapOutAmtGivenIn(ctx sdk.Context, tokenIn sdk.Coin, tokenOutDenom string, swapFee sdk.Dec, minPrice, maxPrice sdk.Dec, poolId uint64) (tokenOut sdk.Coin, err error) {
	tokenInCoin, newLiquidity, newCurrentTick, newCurrentSqrtPrice, err := k.CalcOutAmtGivenIn(ctx, tokenIn, tokenOutDenom, swapFee, minPrice, maxPrice, poolId)
	if err != nil {
		return sdk.Coin{}, err
	}

	k.applySwap(ctx, tokenInCoin, tokenOut, poolId, newLiquidity, newCurrentTick, newCurrentSqrtPrice)

	return tokenInCoin, nil
}

func (p Pool) SpotPrice(ctx sdk.Context, baseAssetDenom string, quoteAssetDenom string) (sdk.Dec, error) {
	return sdk.Dec{}, nil
}

type SwapState struct {
	amountSpecifiedRemaining sdk.Dec // remaining amount of tokens that need to be bought by the pool
	amountCalculated         sdk.Dec // amount out
	sqrtPrice                sdk.Dec // new current price when swap is done
	tick                     sdk.Int // new tick when swap is done
	liquidity                sdk.Dec
}

// this only works on a single directional trade, will implement bi directional trade in next milestone
func (k Keeper) CalcOutAmtGivenIn(ctx sdk.Context,
	tokenIn sdk.Coin,
	tokenOutDenom string,
	swapFee sdk.Dec,
	minPrice, maxPrice sdk.Dec,
	poolId uint64) (tokenOut sdk.Coin, newLiquidity sdk.Dec, currTick sdk.Int, currSqrtPrice sdk.Dec, err error) {
	p := k.getPoolbyId(ctx, poolId)
	asset0 := p.Token0
	asset1 := p.Token1
	tokenAmountInAfterFee := tokenIn.Amount.ToDec().Mul(sdk.OneDec().Sub(swapFee))

	zeroForOne := tokenIn.Denom == asset0

	// get current sqrt price from pool
	curSqrtPrice := p.CurrentSqrtPrice

	// validation
	if tokenIn.Denom != asset0 && tokenIn.Denom != asset1 {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("tokenIn (%s) does not match any asset in pool", tokenIn.Denom)
	}
	if tokenOutDenom != asset0 && tokenOutDenom != asset1 {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("tokenOutDenom (%s) does not match any asset in pool", tokenOutDenom)
	}
	if tokenIn.Denom == tokenOutDenom {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("tokenIn (%s) cannot be the same as tokenOut (%s)", tokenIn.Denom, tokenOutDenom)
	}
	if minPrice.GTE(curSqrtPrice.Power(2)) {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("minPrice (%s) must be less than current price (%s)", minPrice, curSqrtPrice.Power(2))
	}
	if maxPrice.LTE(curSqrtPrice.Power(2)) {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("maxPrice (%s) must be greater than current price (%s)", maxPrice, curSqrtPrice.Power(2))
	}

	// sqrtPrice of upper and lower user defined price range
	sqrtPLowerTick, err := minPrice.ApproxSqrt()
	if err != nil {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("issue calculating square root of minPrice")
	}

	sqrtPUpperTick, err := maxPrice.ApproxSqrt()
	if err != nil {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("issue calculating square root of maxPrice")
	}

	// TODO: How do we remove/generalize this? I am stumped.
	// the following coins represent the virtual amounts in the current price bucket
	amountETH := sdk.NewInt(1000000)
	amountUSDC := sdk.NewInt(5000000000)

	// find liquidity of assetA and assetB
	liq0 := liquidity0(amountETH, curSqrtPrice, sqrtPUpperTick)
	liq1 := liquidity1(amountUSDC, curSqrtPrice, sqrtPLowerTick)

	// utilize the smaller liquidity between assetA and assetB when performing the swap calculation
	liq := sdk.MinDec(liq0, liq1)

	swapState := SwapState{
		amountSpecifiedRemaining: tokenAmountInAfterFee,
		amountCalculated:         sdk.ZeroDec(),
		sqrtPrice:                curSqrtPrice,
		tick:                     priceToTick(curSqrtPrice.Power(2)),
		liquidity:                p.Liquidity,
	}

	// TODO: This should be GT 0 but some instances have very small remainder
	// need to look into fixing this
	for swapState.amountSpecifiedRemaining.GT(sdk.NewDecWithPrec(1, 6)) {
		nextTick, ok := k.NextInitializedTick(ctx, poolId, swapState.tick.Int64(), zeroForOne)
		if !ok {
			return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("there are no more ticks initialized to fill the swap")
		}

		nextSqrtPrice, err := k.tickToSqrtPrice(sdk.NewInt(nextTick))
		if err != nil {
			return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("could not convert next tick (%v) to nextSqrtPrice", sdk.NewInt(nextTick))
		}

		sqrtPrice, amountIn, amountOut := computeSwapStep(
			swapState.sqrtPrice,
			nextSqrtPrice,
			liq,
			swapState.amountSpecifiedRemaining,
			zeroForOne,
		)
		fmt.Println("===sqrt price")
		fmt.Println(sqrtPrice.String())
		fmt.Println("===next sqrt price")
		fmt.Println(nextSqrtPrice.String())
		swapState.amountSpecifiedRemaining = swapState.amountSpecifiedRemaining.Sub(amountIn)
		swapState.amountCalculated = swapState.amountCalculated.Add(amountOut)

		// if we have moved to the next tick,
		if nextSqrtPrice.Equal(sqrtPrice) {
			liquidityDelta, err := k.crossTick(ctx, p.Id, nextTick)

			if err != nil {
				return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), err
			}
			if zeroForOne {
				liquidityDelta = liquidityDelta.Neg()
			}
			swapState.liquidity = swapState.liquidity.Add(liquidityDelta.ToDec())
			if swapState.liquidity.LTE(sdk.ZeroDec()) || swapState.liquidity.IsNil() {
				return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), err
			}
			if zeroForOne {
				swapState.tick = sdk.NewInt(nextTick - 1)
			} else {
				swapState.tick = sdk.NewInt(nextTick)
			}
		} else {
			swapState.tick = priceToTick(sqrtPrice.Power(2))
		}
	}
	tokenOut = sdk.NewCoin(tokenOutDenom, swapState.amountCalculated.RoundInt())

	return tokenOut, swapState.liquidity, swapState.tick, swapState.sqrtPrice, nil
}

func (k *Keeper) SwapInAmtGivenOut(ctx sdk.Context, tokenOut sdk.Coin, tokenInDenom string, swapFee sdk.Dec, minPrice, maxPrice sdk.Dec, poolId uint64) (tokenIn sdk.Coin, err error) {
	tokenInCoin, newLiquidity, newCurrentTick, newCurrentSqrtPrice, err := k.CalcInAmtGivenOut(ctx, tokenOut, tokenInDenom, swapFee, sdk.ZeroDec(), sdk.NewDec(9999999999), poolId)
	if err != nil {
		return sdk.Coin{}, err
	}

	k.applySwap(ctx, tokenInCoin, tokenOut, poolId, newLiquidity, newCurrentTick, newCurrentSqrtPrice)

	return tokenInCoin, nil
}

func (k Keeper) CalcInAmtGivenOut(ctx sdk.Context, tokenOut sdk.Coin, tokenInDenom string, swapFee sdk.Dec, minPrice, maxPrice sdk.Dec, poolId uint64) (sdk.Coin, sdk.Dec, sdk.Int, sdk.Dec, error) {
	tokenOutAmt := tokenOut.Amount.ToDec()
	p := k.getPoolbyId(ctx, poolId)
	asset0 := p.Token0
	asset1 := p.Token1
	zeroForOne := tokenOut.Denom == asset0

	// get current sqrt price from pool
	curSqrtPrice := p.CurrentSqrtPrice

	// validation
	if tokenOut.Denom != asset0 && tokenOut.Denom != asset1 {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("tokenOut denom (%s) does not match any asset in pool", tokenOut.Denom)
	}
	if tokenInDenom != asset0 && tokenInDenom != asset1 {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("tokenInDenom (%s) does not match any asset in pool", tokenInDenom)
	}
	if tokenOut.Denom == tokenInDenom {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("tokenOut (%s) cannot be the same as tokenIn (%s)", tokenOut.Denom, tokenInDenom)
	}
	if minPrice.GTE(curSqrtPrice.Power(2)) {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("minPrice (%s) must be less than current price (%s)", minPrice, curSqrtPrice.Power(2))
	}
	if maxPrice.LTE(curSqrtPrice.Power(2)) {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("maxPrice (%s) must be greater than current price (%s)", maxPrice, curSqrtPrice.Power(2))
	}

	// sqrtPrice of upper and lower user defined price range
	sqrtPLowerTick, err := minPrice.ApproxSqrt()
	if err != nil {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("issue calculating square root of minPrice")
	}

	sqrtPUpperTick, err := maxPrice.ApproxSqrt()
	if err != nil {
		return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("issue calculating square root of maxPrice")
	}

	// TODO: How do we remove/generalize this? I am stumped.
	amountETH := sdk.NewInt(1000000)
	amountUSDC := sdk.NewInt(5000000000)

	// find liquidity of assetA and assetB
	liq0 := liquidity0(amountETH, curSqrtPrice, sqrtPUpperTick)
	liq1 := liquidity1(amountUSDC, curSqrtPrice, sqrtPLowerTick)

	// utilize the smaller liquidity between assetA and assetB when performing the swap calculation
	liq := sdk.MinDec(liq0, liq1)

	swapState := SwapState{
		amountSpecifiedRemaining: tokenOutAmt,
		amountCalculated:         sdk.ZeroDec(),
		sqrtPrice:                curSqrtPrice,
		tick:                     priceToTick(curSqrtPrice.Power(2)),
	}

	// TODO: This should be GT 0 but some instances have very small remainder
	// need to look into fixing this
	for swapState.amountSpecifiedRemaining.GT(sdk.NewDecWithPrec(1, 6)) {
		nextTick, ok := k.NextInitializedTick(ctx, poolId, swapState.tick.Int64(), zeroForOne)

		// TODO: we can enable this error checking once we fix tick initialization
		if !ok {
			return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("there are no more ticks initialized to fill the swap")
		}
		nextSqrtPrice, err := k.tickToSqrtPrice(sdk.NewInt(nextTick))
		if err != nil {
			return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), err
		}

		// TODO: In and out get flipped based on if we are calculating for in or out, need to fix this
		sqrtPrice, amountIn, amountOut := computeSwapStep(
			swapState.sqrtPrice,
			nextSqrtPrice,
			liq,
			swapState.amountSpecifiedRemaining,
			zeroForOne,
		)

		swapState.amountSpecifiedRemaining = swapState.amountSpecifiedRemaining.Sub(amountIn)
		swapState.amountCalculated = swapState.amountCalculated.Add(amountOut.Quo(sdk.OneDec().Sub(swapFee)))

		if swapState.sqrtPrice.Equal(sqrtPrice) {
			liquidityDelta, err := k.crossTick(ctx, p.Id, nextTick)
			if err != nil {
				return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), err
			}
			if !zeroForOne {
				liquidityDelta = liquidityDelta.Neg()
			}
			swapState.liquidity = swapState.liquidity.Add(liquidityDelta.ToDec())
			if swapState.liquidity.LTE(sdk.ZeroDec()) || swapState.liquidity.IsNil() {
				return sdk.Coin{}, sdk.ZeroDec(), sdk.ZeroInt(), sdk.ZeroDec(), fmt.Errorf("no liquidity available, cannot swap")
			}
			if !zeroForOne {
				swapState.tick = sdk.NewInt(nextTick - 1)
			} else {
				swapState.tick = sdk.NewInt(nextTick)
			}
		} else {
			swapState.tick = priceToTick(sqrtPrice.Power(2))
		}
	}

	return sdk.NewCoin(tokenInDenom, swapState.amountCalculated.RoundInt()), swapState.liquidity, swapState.tick, swapState.sqrtPrice, nil
}

// ApplySwap.
func (k *Keeper) applySwap(ctx sdk.Context, tokenIn sdk.Coin, tokenOut sdk.Coin, poolId uint64, newLiquidity sdk.Dec, newCurrentTick sdk.Int, newCurrentSqrtPrice sdk.Dec) {
	// Fixed gas consumption per swap to prevent spam
	ctx.GasMeter().ConsumeGas(gammtypes.BalancerGasFeeForSwap, "cl pool swap computation")
	pool := k.getPoolbyId(ctx, poolId)

	pool.Liquidity = newLiquidity
	pool.CurrentTick = newCurrentTick
	pool.CurrentSqrtPrice = newCurrentSqrtPrice
	k.setPoolById(ctx, pool.Id, pool)
}
