package concentrated_liquidity_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/osmosis-labs/osmosis/osmoutils/accum"
	concentratedliquidity "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity"
	"github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/internal/math"
	clmodel "github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/model"
	"github.com/osmosis-labs/osmosis/v14/x/concentrated-liquidity/types"
)

var (
	oneEth = sdk.NewDecCoin(ETH, sdk.OneInt())
)

func (s *KeeperTestSuite) TestInitializeFeeAccumulatorPosition() {
	defaultPoolId := uint64(1)
	defaultLiquidityDelta := sdk.MustNewDecFromStr("10.0")
	type initFeeAccumTest struct {
		setPoolAccumulator  bool
		setExistingPosition bool
		expectedPass        bool
	}
	tests := map[string]initFeeAccumTest{
		"existing accumulator, new position": {
			setPoolAccumulator:  true,
			setExistingPosition: false,
			expectedPass:        true,
		},
		"existing accumulator, try overriding existing position": {
			setPoolAccumulator:  true,
			setExistingPosition: true,
			expectedPass:        true,
		},
		"error: non-existing accumulator": {
			setPoolAccumulator: false,
			expectedPass:       false,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()
			clKeeper := s.App.ConcentratedLiquidityKeeper
			if tc.setPoolAccumulator {
				err := clKeeper.CreateFeeAccumulator(s.Ctx, defaultPoolId)
				s.Require().NoError(err)
			}
			if tc.setExistingPosition {
				// initialize with default liquidity delta * 2 to see if sut correctly initializes
				err := clKeeper.InitializeFeeAccumulatorPosition(s.Ctx, defaultPoolId, s.TestAccs[0], defaultLiquidityDelta.Add(defaultLiquidityDelta))
				s.Require().NoError(err)
			}

			// system under test
			err := clKeeper.InitializeFeeAccumulatorPosition(s.Ctx, defaultPoolId, s.TestAccs[0], defaultLiquidityDelta)
			if tc.expectedPass {
				s.Require().NoError(err)

				// get fee accum and see if position size has been properly initialized
				poolFeeAccumulator, err := clKeeper.GetFeeAccumulator(s.Ctx, defaultPoolId)
				s.Require().NoError(err)
				positionSize, err := poolFeeAccumulator.GetPositionSize(string(s.TestAccs[0].String()))
				s.Require().NoError(err)
				// position should have been properly initialzied to liquidityDelta provided
				s.Require().Equal(positionSize, defaultLiquidityDelta)
			} else {
				s.Require().Error(err)
			}
		})
	}
}

func (s *KeeperTestSuite) TestGetFeeGrowthOutside() {
	type feeGrowthOutsideTest struct {
		poolSetup bool

		lowerTick                 int64
		upperTick                 int64
		currentTick               int64
		lowerTickFeeGrowthOutside sdk.DecCoins
		upperTickFeeGrowthOutside sdk.DecCoins
		globalFeeGrowth           sdk.DecCoin

		expectedFeeGrowthOutside sdk.DecCoins
		invalidTick              bool
		expectedError            bool
	}

	defaultAccumCoins := sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(10)))
	defaultPoolId := uint64(1)
	defaultInitialLiquidity := sdk.OneDec()
	// defaultLiquidityForLowerTick := sdk.MustNewDecFromStr("10.0")
	// defaultLiquidityForUpperTick := sdk.MustNewDecFromStr("20.0")

	defaultUpperTickIndex := int64(5)
	defaultLowerTickIndex := int64(3)

	tests := map[string]feeGrowthOutsideTest{
		"single swap left -> right: 2 ticks, one share - current price > upper tick": {
			poolSetup:                 true,
			lowerTick:                 0,
			upperTick:                 1,
			currentTick:               2,
			lowerTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(0))),
			upperTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(10))),
			globalFeeGrowth:           sdk.NewDecCoin(ETH, sdk.NewInt(10)),
			expectedFeeGrowthOutside:  defaultAccumCoins,
			expectedError:             false,
		},
		"single swap left -> right: 3 ticks, two shares - current price > upper tick": {
			poolSetup:                 true,
			lowerTick:                 0,
			upperTick:                 2,
			currentTick:               3,
			lowerTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(0))),
			upperTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(10))),
			globalFeeGrowth:           sdk.NewDecCoin(ETH, sdk.NewInt(10)),
			expectedFeeGrowthOutside:  defaultAccumCoins,
			expectedError:             false,
		},
		"single swap left -> right: 2 ticks, one share - current price == upper tick": {
			poolSetup:                 true,
			lowerTick:                 0,
			upperTick:                 1,
			currentTick:               0,
			lowerTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(0))),
			upperTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(10))),
			globalFeeGrowth:           sdk.NewDecCoin(ETH, sdk.NewInt(10)),
			expectedFeeGrowthOutside:  defaultAccumCoins,
			expectedError:             false,
		},
		// imagine single swap over entire position
		// crossing right > left and stopping at lower tick
		// In this case, all fees must have been accrued inside the tick
		// Since we track fees accrued below a tick, both upper and lower position
		// ticks are zero
		"single swap right -> left: 2 ticks, one share - current price == lower tick": {
			poolSetup:                 true,
			lowerTick:                 0,
			upperTick:                 1,
			currentTick:               0,
			lowerTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(0))),
			upperTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(10))),
			globalFeeGrowth:           sdk.NewDecCoin(ETH, sdk.NewInt(10)),
			expectedFeeGrowthOutside:  defaultAccumCoins,
			expectedError:             false,
		},
		"single swap right -> left: 2 ticks, one share - current price < lower tick": {
			poolSetup:                 true,
			lowerTick:                 0,
			upperTick:                 1,
			currentTick:               -1,
			lowerTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(0))),
			upperTickFeeGrowthOutside: sdk.NewDecCoins(sdk.NewDecCoin(ETH, sdk.NewInt(10))),
			globalFeeGrowth:           sdk.NewDecCoin(ETH, sdk.NewInt(10)),
			expectedFeeGrowthOutside:  defaultAccumCoins,
			expectedError:             false,
		},
		"error: pool has not been setup": {
			poolSetup:     false,
			expectedError: true,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()

			// if pool set up true, set up default pool
			var pool types.ConcentratedPoolExtension
			if tc.poolSetup {
				pool = s.PrepareConcentratedPool()

				s.initializeTick(s.Ctx, tc.lowerTick, defaultInitialLiquidity, tc.lowerTickFeeGrowthOutside, false)
				s.initializeTick(s.Ctx, tc.upperTick, defaultInitialLiquidity, tc.upperTickFeeGrowthOutside, true)
				pool.SetCurrentTick(sdk.NewInt(tc.currentTick))
				s.App.ConcentratedLiquidityKeeper.SetPool(s.Ctx, pool)
				// poolId = pool.GetId()
				err := s.App.ConcentratedLiquidityKeeper.ChargeFee(s.Ctx, validPoolId, tc.globalFeeGrowth)
				s.Require().NoError(err)
			}

			// system under test
			feeGrowthOutside, err := s.App.ConcentratedLiquidityKeeper.GetFeeGrowthOutside(s.Ctx, defaultPoolId, defaultLowerTickIndex, defaultUpperTickIndex)
			if tc.expectedError {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				// check if returned fee growth outside has correct value
				s.Require().Equal(feeGrowthOutside, tc.expectedFeeGrowthOutside)
			}
		})
	}
}

func (s *KeeperTestSuite) TestCalculateFeeGrowth() {
	defaultGeeFrowthGlobal := sdk.NewDecCoins(sdk.NewDecCoin("uosmo", sdk.NewInt(10)))
	defaultGeeFrowthOutside := sdk.NewDecCoins(sdk.NewDecCoin("uosmo", sdk.NewInt(3)))

	defaultSmallerTargetTick := int64(1)
	defaultCurrentTick := int64(2)
	defaultLargerTargetTick := int64(3)

	type calcFeeGrowthTest struct {
		isUpperTick                bool
		isCurrentTickGTETargetTick bool
		expectedFeeGrowth          sdk.DecCoins
	}

	tests := map[string]calcFeeGrowthTest{
		"current Tick is greater than the upper tick": {
			isUpperTick:                true,
			isCurrentTickGTETargetTick: false,
			expectedFeeGrowth:          defaultGeeFrowthOutside,
		},
		"current Tick is less than the upper tick": {
			isUpperTick:                true,
			isCurrentTickGTETargetTick: true,
			expectedFeeGrowth:          defaultGeeFrowthGlobal.Sub(defaultGeeFrowthOutside),
		},
		"current Tick is less than the lower tick": {
			isUpperTick:                false,
			isCurrentTickGTETargetTick: false,
			expectedFeeGrowth:          defaultGeeFrowthGlobal.Sub(defaultGeeFrowthOutside),
		},
		"current Tick is greater than the lower tick": {
			isUpperTick:                false,
			isCurrentTickGTETargetTick: true,
			expectedFeeGrowth:          defaultGeeFrowthOutside,
		},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			s.SetupTest()
			var targetTick int64
			if tc.isCurrentTickGTETargetTick {
				targetTick = defaultSmallerTargetTick
			} else {
				targetTick = defaultLargerTargetTick
			}
			feeGrowth := concentratedliquidity.CalculateFeeGrowth(
				targetTick,
				defaultGeeFrowthOutside,
				defaultCurrentTick,
				defaultGeeFrowthGlobal,
				tc.isUpperTick,
			)
			s.Require().Equal(feeGrowth, tc.expectedFeeGrowth)
		})
	}

}

func (suite *KeeperTestSuite) TestGetInitialFeeGrowthOutsideForTick() {
	const (
		validPoolId = 1
	)

	var (
		initialPoolTick = math.PriceToTick(DefaultAmt1.ToDec().Quo(DefaultAmt0.ToDec())).Int64()
	)

	tests := map[string]struct {
		poolId                   uint64
		tick                     int64
		initialGlobalFeeGrowth   sdk.DecCoin
		shouldAvoidCreatingAccum bool

		expectedInitialFeeGrowthOutside sdk.DecCoins
		expectError                     error
	}{
		"current tick > tick -> fee growth global": {
			poolId:                 validPoolId,
			tick:                   initialPoolTick - 1,
			initialGlobalFeeGrowth: oneEth,

			expectedInitialFeeGrowthOutside: sdk.NewDecCoins(oneEth),
		},
		"current tick == tick -> fee growth global": {
			poolId:                 validPoolId,
			tick:                   initialPoolTick,
			initialGlobalFeeGrowth: oneEth,

			expectedInitialFeeGrowthOutside: sdk.NewDecCoins(oneEth),
		},
		"current tick < tick -> empty coins": {
			poolId:                 validPoolId,
			tick:                   initialPoolTick + 1,
			initialGlobalFeeGrowth: oneEth,

			expectedInitialFeeGrowthOutside: concentratedliquidity.EmptyCoins,
		},
		"pool does not exist": {
			poolId:                 validPoolId + 1,
			tick:                   initialPoolTick - 1,
			initialGlobalFeeGrowth: oneEth,

			expectError: types.PoolNotFoundError{PoolId: validPoolId + 1},
		},
		"accumulator does not exist": {
			poolId:                   validPoolId,
			tick:                     0,
			initialGlobalFeeGrowth:   oneEth,
			shouldAvoidCreatingAccum: true,

			expectError: accum.AccumDoesNotExistError{AccumName: concentratedliquidity.GetFeeAccumulatorName(validPoolId)},
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			suite.SetupTest()
			ctx := suite.Ctx
			clKeeper := suite.App.ConcentratedLiquidityKeeper

			pool, err := clmodel.NewConcentratedLiquidityPool(validPoolId, USDC, ETH, DefaultTickSpacing)
			suite.Require().NoError(err)

			err = clKeeper.SetPool(ctx, &pool)
			suite.Require().NoError(err)

			if !tc.shouldAvoidCreatingAccum {
				err = clKeeper.CreateFeeAccumulator(ctx, validPoolId)
				suite.Require().NoError(err)

				// Setup test position to make sure that tick is initialized
				suite.SetupPosition(validPoolId)

				err = clKeeper.ChargeFee(ctx, validPoolId, tc.initialGlobalFeeGrowth)
				suite.Require().NoError(err)
			}

			// System under test.
			initialFeeGrowthOutside, err := clKeeper.GetInitialFeeGrowthOutsideForTick(ctx, tc.poolId, tc.tick)

			if tc.expectError != nil {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expectError)
				return
			}
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expectedInitialFeeGrowthOutside, initialFeeGrowthOutside)
		})
	}
}

func (suite *KeeperTestSuite) TestChargeFee() {
	// setup once at the beginning.
	suite.SetupTest()

	ctx := suite.Ctx
	clKeeper := suite.App.ConcentratedLiquidityKeeper

	// create fee accumulators with ids 1 and 2 but not 3.
	err := clKeeper.CreateFeeAccumulator(ctx, 1)
	suite.Require().NoError(err)
	err = clKeeper.CreateFeeAccumulator(ctx, 2)
	suite.Require().NoError(err)

	tests := map[string]struct {
		poolId    uint64
		feeUpdate sdk.DecCoin

		expectedGlobalGrowth sdk.DecCoins
		expectError          error
	}{
		"pool id 1 - one eth": {
			poolId:    1,
			feeUpdate: oneEth,

			expectedGlobalGrowth: sdk.NewDecCoins(oneEth),
		},
		"pool id 1 - 2 usdc": {
			poolId:    1,
			feeUpdate: sdk.NewDecCoin(USDC, sdk.NewInt(2)),

			expectedGlobalGrowth: sdk.NewDecCoins(oneEth).Add(sdk.NewDecCoin(USDC, sdk.NewInt(2))),
		},
		"pool id 2 - 1 usdc": {
			poolId:    2,
			feeUpdate: oneEth,

			expectedGlobalGrowth: sdk.NewDecCoins(oneEth),
		},
		"accumulator does not exist": {
			poolId:    3,
			feeUpdate: oneEth,

			expectError: accum.AccumDoesNotExistError{AccumName: concentratedliquidity.GetFeeAccumulatorName(3)},
		},
	}

	for name, tc := range tests {
		tc := tc
		suite.Run(name, func() {
			// System under test.
			err := clKeeper.ChargeFee(ctx, tc.poolId, tc.feeUpdate)

			if tc.expectError != nil {
				suite.Require().Error(err)
				suite.Require().ErrorIs(err, tc.expectError)
				return
			}
			suite.Require().NoError(err)

			feeAcumulator, err := clKeeper.GetFeeAccumulator(ctx, tc.poolId)
			suite.Require().NoError(err)
			suite.Require().Equal(tc.expectedGlobalGrowth, feeAcumulator.GetValue())
		})
	}
}

func (s *KeeperTestSuite) initializeTick(ctx sdk.Context, tickIndex int64, initialLiquidity sdk.Dec, feeGrowthOutside sdk.DecCoins, isLower bool) {
	err := s.App.ConcentratedLiquidityKeeper.InitOrUpdateTick(ctx, validPoolId, tickIndex, initialLiquidity, isLower)
	s.Require().NoError(err)

	tickInfo, err := s.App.ConcentratedLiquidityKeeper.GetTickInfo(ctx, validPoolId, tickIndex)
	s.Require().NoError(err)

	tickInfo.FeeGrowthOutside = feeGrowthOutside

	s.App.ConcentratedLiquidityKeeper.SetTickInfo(ctx, validPoolId, tickIndex, tickInfo)
}
