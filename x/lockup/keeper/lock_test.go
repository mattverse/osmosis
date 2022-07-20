package keeper_test

import (
	"fmt"
	"time"

	"github.com/osmosis-labs/osmosis/v7/x/lockup/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (suite *KeeperTestSuite) TestBeginUnlocking() { // test for all unlockable coins
	suite.SetupTest()

	// initial check
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// check locks
	locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 1)
	suite.Require().Equal(locks[0].EndTime, time.Time{})
	suite.Require().Equal(locks[0].IsUnlocking(), false)

	// begin unlock
	locks, err = suite.App.LockupKeeper.BeginUnlockAllNotUnlockings(suite.Ctx, addr1)
	unlockedCoins := suite.App.LockupKeeper.GetCoinsFromLocks(locks)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 1)
	suite.Require().Equal(unlockedCoins, coins)
	suite.Require().Equal(locks[0].ID, uint64(1))

	// check locks
	locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 1)
	suite.Require().NotEqual(locks[0].EndTime, time.Time{})
	suite.Require().NotEqual(locks[0].IsUnlocking(), false)
}

func (suite *KeeperTestSuite) TestBeginUnlockPeriodLock() {
	suite.SetupTest()

	// initial check
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// check locks
	locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 1)
	suite.Require().Equal(locks[0].EndTime, time.Time{})
	suite.Require().Equal(locks[0].IsUnlocking(), false)

	// begin unlock
	err = suite.App.LockupKeeper.BeginUnlock(suite.Ctx, locks[0].ID, nil)
	suite.Require().NoError(err)

	// check locks
	locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().NotEqual(locks[0].EndTime, time.Time{})
	suite.Require().NotEqual(locks[0].IsUnlocking(), false)
}

func (suite *KeeperTestSuite) TestGetPeriodLocks() {
	suite.SetupTest()

	// initial check
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// check locks
	locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 1)
}

func (suite *KeeperTestSuite) TestUnlockPeriodLockByID() {
	suite.SetupTest()
	now := suite.Ctx.BlockTime()
	lockKeeper := suite.App.LockupKeeper

	// initial check
	locks, err := lockKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// unlock lock just now
	lock, err := lockKeeper.GetLockByID(suite.Ctx, 1)
	suite.Require().NoError(err)
	err = lockKeeper.UnlockMaturedLock(suite.Ctx, lock.ID)
	suite.Require().Error(err)

	// move start time to 1 second in the future.
	suite.Ctx = suite.Ctx.WithBlockTime(now.Add(time.Second))

	// Try to finish unlocking a lock, before starting unlock.
	lock, err = lockKeeper.GetLockByID(suite.Ctx, 1)
	suite.Require().NoError(err)
	err = lockKeeper.UnlockMaturedLock(suite.Ctx, lock.ID)
	suite.Require().Error(err)

	// begin unlock
	lock, err = lockKeeper.GetLockByID(suite.Ctx, 1)
	suite.Require().NoError(err)
	err = lockKeeper.BeginUnlock(suite.Ctx, lock.ID, nil)
	suite.Require().NoError(err)

	// unlock 1s after begin unlock
	lock, err = lockKeeper.GetLockByID(suite.Ctx, 1)
	suite.Require().NoError(err)
	err = lockKeeper.UnlockMaturedLock(suite.Ctx.WithBlockTime(now.Add(time.Second*2)), lock.ID)
	suite.Require().NoError(err)

	// check locks
	locks, err = lockKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)
}

func (suite *KeeperTestSuite) TestUnlock() {
	// test for coin unlocking
	suite.SetupTest()
	now := suite.Ctx.BlockTime()

	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}

	// lock with balance
	suite.FundAcc(addr1, coins)
	lock, err := suite.App.LockupKeeper.CreateLock(suite.Ctx, addr1, coins, time.Second)
	suite.Require().NoError(err)

	// begin unlock with lock object
	err = suite.App.LockupKeeper.BeginUnlock(suite.Ctx, lock.ID, nil)
	suite.Require().NoError(err)

	lockPtr, err := suite.App.LockupKeeper.GetLockByID(suite.Ctx, lock.ID)
	suite.Require().NoError(err)

	// unlock with lock object
	err = suite.App.LockupKeeper.UnlockMaturedLock(suite.Ctx.WithBlockTime(now.Add(time.Second)), lockPtr.ID)
	suite.Require().NoError(err)
}

func (suite *KeeperTestSuite) TestPartialUnlock() {
	suite.SetupTest()
	now := suite.Ctx.BlockTime()

	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}

	// lock with balance
	suite.FundAcc(addr1, coins)
	lock, err := suite.App.LockupKeeper.CreateLock(suite.Ctx, addr1, coins, time.Second)
	suite.Require().NoError(err)

	// check unlocking coins
	unlockings := suite.App.LockupKeeper.GetAccountUnlockingCoins(suite.Ctx, addr1)
	suite.Require().Equal(len(unlockings), 0)

	// check locked coins
	locked := suite.App.LockupKeeper.GetAccountLockedCoins(suite.Ctx, addr1)
	suite.Require().Equal(len(locked), 1)
	suite.Require().Equal(locked[0].Amount.Int64(), int64(10))

	// test exceeding coins
	exceedingCoins := sdk.Coins{sdk.NewInt64Coin("stake", 15)}
	err = suite.App.LockupKeeper.BeginUnlock(suite.Ctx, lock.ID, exceedingCoins)
	suite.Require().Error(err)

	// test invalid coins
	invalidCoins := sdk.Coins{sdk.NewInt64Coin("unknown", 1)}
	err = suite.App.LockupKeeper.BeginUnlock(suite.Ctx, lock.ID, invalidCoins)
	suite.Require().Error(err)

	// begin unlock partial amount
	partialCoins := sdk.Coins{sdk.NewInt64Coin("stake", 1)}
	err = suite.App.LockupKeeper.BeginUnlock(suite.Ctx, lock.ID, partialCoins)
	suite.Require().NoError(err)

	// check unlocking coins
	unlockings = suite.App.LockupKeeper.GetAccountUnlockingCoins(suite.Ctx, addr1)
	suite.Require().Equal(len(unlockings), 1)
	suite.Require().Equal(unlockings[0].Amount.Int64(), int64(1))

	// check locked coins
	locked = suite.App.LockupKeeper.GetAccountLockedCoins(suite.Ctx, addr1)
	suite.Require().Equal(len(locked), 1)
	suite.Require().Equal(locked[0].Amount.Int64(), int64(10))

	// check locked coins after the unlocking period
	locked = suite.App.LockupKeeper.GetAccountLockedCoins(suite.Ctx.WithBlockTime(now.Add(time.Second)), addr1)
	suite.Require().Equal(len(locked), 1)
	suite.Require().Equal(locked[0].Amount.Int64(), int64(9))

	// Finish unlocking partial unlock
	partialUnlock := suite.App.LockupKeeper.GetAccountPeriodLocks(suite.Ctx, addr1)[1]
	err = suite.App.LockupKeeper.UnlockMaturedLock(suite.Ctx.WithBlockTime(now.Add(time.Second)), partialUnlock.ID)
	suite.Require().NoError(err)

	// check unlocking coins
	unlockings = suite.App.LockupKeeper.GetAccountUnlockingCoins(suite.Ctx, addr1)
	suite.Require().Equal(len(unlockings), 0)

	// check locked coins
	locked = suite.App.LockupKeeper.GetAccountLockedCoins(suite.Ctx, addr1)
	suite.Require().Equal(len(locked), 1)
	suite.Require().Equal(locked[0].Amount.Int64(), int64(9))
}

func (suite *KeeperTestSuite) TestModuleLockedCoins() {
	suite.SetupTest()

	// initial check
	lockedCoins := suite.App.LockupKeeper.GetModuleLockedCoins(suite.Ctx)
	suite.Require().Equal(lockedCoins, sdk.Coins(nil))

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// final check
	lockedCoins = suite.App.LockupKeeper.GetModuleLockedCoins(suite.Ctx)
	suite.Require().Equal(lockedCoins, coins)
}

func (suite *KeeperTestSuite) TestLocksPastTimeDenom() {
	suite.SetupTest()

	now := time.Now()
	suite.Ctx = suite.Ctx.WithBlockTime(now)

	// initial check
	locks := suite.App.LockupKeeper.GetLocksPastTimeDenom(suite.Ctx, "stake", now)
	suite.Require().Len(locks, 0)

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// final check
	locks = suite.App.LockupKeeper.GetLocksPastTimeDenom(suite.Ctx, "stake", now)
	suite.Require().Len(locks, 1)
}

func (suite *KeeperTestSuite) TestLocksLongerThanDurationDenom() {
	suite.SetupTest()

	// initial check
	duration := time.Second
	locks := suite.App.LockupKeeper.GetLocksLongerThanDurationDenom(suite.Ctx, "stake", duration)
	suite.Require().Len(locks, 0)

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// final check
	locks = suite.App.LockupKeeper.GetLocksLongerThanDurationDenom(suite.Ctx, "stake", duration)
	suite.Require().Len(locks, 1)
}

func (suite *KeeperTestSuite) TestCreateLock() {
	suite.SetupTest()

	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}

	// test locking without balance
	_, err := suite.App.LockupKeeper.CreateLock(suite.Ctx, addr1, coins, time.Second)
	suite.Require().Error(err)

	suite.FundAcc(addr1, coins)

	lock, err := suite.App.LockupKeeper.CreateLock(suite.Ctx, addr1, coins, time.Second)
	suite.Require().NoError(err)

	// check new lock
	suite.Require().Equal(coins, lock.Coins)
	suite.Require().Equal(time.Second, lock.Duration)
	suite.Require().Equal(time.Time{}, lock.EndTime)
	suite.Require().Equal(uint64(1), lock.ID)

	lockID := suite.App.LockupKeeper.GetLastLockID(suite.Ctx)
	suite.Require().Equal(uint64(1), lockID)

	// check accumulation store
	accum := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		LockQueryType: types.ByDuration,
		Denom:         "stake",
		Duration:      time.Second,
	})
	suite.Require().Equal(accum.String(), "10")

	// create new lock
	coins = sdk.Coins{sdk.NewInt64Coin("stake", 20)}
	suite.FundAcc(addr1, coins)

	lock, err = suite.App.LockupKeeper.CreateLock(suite.Ctx, addr1, coins, time.Second)
	suite.Require().NoError(err)

	lockID = suite.App.LockupKeeper.GetLastLockID(suite.Ctx)
	suite.Require().Equal(uint64(2), lockID)

	// check accumulation store
	accum = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		LockQueryType: types.ByDuration,
		Denom:         "stake",
		Duration:      time.Second,
	})
	suite.Require().Equal(accum.String(), "30")

	// check balance
	balance := suite.App.BankKeeper.GetBalance(suite.Ctx, addr1, "stake")
	suite.Require().Equal(sdk.ZeroInt(), balance.Amount)

	acc := suite.App.AccountKeeper.GetModuleAccount(suite.Ctx, types.ModuleName)
	balance = suite.App.BankKeeper.GetBalance(suite.Ctx, acc.GetAddress(), "stake")
	suite.Require().Equal(sdk.NewInt(30), balance.Amount)
}

func (suite *KeeperTestSuite) TestAddTokensToLock() {
	initialLockCoin := sdk.NewInt64Coin("stake", 10)
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	addr2 := sdk.AccAddress([]byte("addr2---------------"))

	testCases := []struct {
		name                         string
		tokenToAdd                   sdk.Coin
		duration                     time.Duration
		lockingAddress               sdk.AccAddress
		expectAddTokensToLockSuccess bool
		expectedError                bool
	}{
		{
			name:                         "normal add tokens to lock case",
			tokenToAdd:                   initialLockCoin,
			duration:                     time.Second,
			lockingAddress:               addr1,
			expectAddTokensToLockSuccess: true,
		},
		{
			name:           "not the owner of the lock",
			tokenToAdd:     initialLockCoin,
			duration:       time.Second,
			lockingAddress: addr2,
			expectedError:  true,
		},
		{
			name:           "lock with matching duration not existing",
			tokenToAdd:     initialLockCoin,
			duration:       time.Second * 2,
			lockingAddress: addr1,
			expectedError:  true,
		},
		{
			name:           "lock invalid tokens",
			tokenToAdd:     sdk.NewCoin("unknown", sdk.NewInt(10)),
			duration:       time.Second,
			lockingAddress: addr1,
			expectedError:  true,
		},
		{
			name:           "token to add exceeds balance",
			tokenToAdd:     sdk.NewCoin("stake", sdk.NewInt(20)),
			duration:       time.Second,
			lockingAddress: addr1,
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		// lock with balance
		suite.FundAcc(addr1, sdk.Coins{initialLockCoin})
		originalLock, err := suite.App.LockupKeeper.CreateLock(suite.Ctx, addr1, sdk.Coins{initialLockCoin}, time.Second)
		suite.Require().NoError(err)

		suite.FundAcc(addr1, sdk.Coins{initialLockCoin})
		balanceBeforeLock := suite.App.BankKeeper.GetAllBalances(suite.Ctx, tc.lockingAddress)

		lockID, err := suite.App.LockupKeeper.AddToExistingLock(suite.Ctx, tc.lockingAddress, tc.tokenToAdd, tc.duration)
		if tc.expectedError {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}

		if tc.expectAddTokensToLockSuccess {
			suite.Require().NoError(err)

			// get updated lock
			lock, err := suite.App.LockupKeeper.GetLockByID(suite.Ctx, lockID)
			suite.Require().NoError(err)

			// check that tokens have been added successfully to the lock
			suite.Require().True(sdk.Coins{initialLockCoin.Add(tc.tokenToAdd)}.IsEqual(lock.Coins))

			// check balance has decreased
			balanceAfterLock := suite.App.BankKeeper.GetAllBalances(suite.Ctx, tc.lockingAddress)
			suite.Require().True(balanceBeforeLock.IsEqual(balanceAfterLock.Add(tc.tokenToAdd)))

			// check accumulation store
			accum := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
				LockQueryType: types.ByDuration,
				Denom:         "stake",
				Duration:      time.Second,
			})
			suite.Require().Equal(initialLockCoin.Amount.Add(tc.tokenToAdd.Amount), accum)
		} else {
			suite.Require().Equal(uint64(0), lockID)

			lock, err := suite.App.LockupKeeper.GetLockByID(suite.Ctx, originalLock.ID)
			suite.Require().NoError(err)

			// check that locked coins haven't changed
			suite.Require().True(lock.Coins.IsEqual(sdk.Coins{initialLockCoin}))

			// check accumulation store didn't change
			accum := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
				LockQueryType: types.ByDuration,
				Denom:         "stake",
				Duration:      time.Second,
			})
			suite.Require().Equal(initialLockCoin.Amount, accum)
		}
	}
}

func (suite *KeeperTestSuite) TestLock() {
	suite.SetupTest()

	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}

	lock := types.PeriodLock{
		ID:       1,
		Owner:    addr1.String(),
		Duration: time.Second,
		EndTime:  time.Time{},
		Coins:    coins,
	}

	// test locking without balance
	err := suite.App.LockupKeeper.Lock(suite.Ctx, lock, coins)
	suite.Require().Error(err)

	// check accumulation store
	accum := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		LockQueryType: types.ByDuration,
		Denom:         "stake",
		Duration:      time.Second,
	})
	suite.Require().Equal(accum.String(), "0")

	suite.FundAcc(addr1, coins)
	err = suite.App.LockupKeeper.Lock(suite.Ctx, lock, coins)
	suite.Require().NoError(err)

	// check accumulation store
	accum = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		LockQueryType: types.ByDuration,
		Denom:         "stake",
		Duration:      time.Second,
	})
	suite.Require().Equal(accum.String(), "10")

	balance := suite.App.BankKeeper.GetBalance(suite.Ctx, addr1, "stake")
	suite.Require().Equal(sdk.ZeroInt(), balance.Amount)

	acc := suite.App.AccountKeeper.GetModuleAccount(suite.Ctx, types.ModuleName)
	balance = suite.App.BankKeeper.GetBalance(suite.Ctx, acc.GetAddress(), "stake")
	suite.Require().Equal(sdk.NewInt(10), balance.Amount)
}

func (suite *KeeperTestSuite) AddTokensToLockForSynth() {
	suite.SetupTest()

	// lock coins
	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr1, coins, time.Second)

	// lock coins on other durations
	coins = sdk.Coins{sdk.NewInt64Coin("stake", 20)}
	suite.LockTokens(addr1, coins, time.Second*2)
	coins = sdk.Coins{sdk.NewInt64Coin("stake", 30)}
	suite.LockTokens(addr1, coins, time.Second*3)

	synthlocks := []types.SyntheticLock{}
	// make three synthetic locks on each locks
	for i := uint64(1); i <= 3; i++ {
		// testing not unlocking synthlock, with same duration with underlying
		synthlock := types.SyntheticLock{
			UnderlyingLockId: i,
			SynthDenom:       fmt.Sprintf("synth1/%d", i),
			Duration:         time.Second * time.Duration(i),
		}
		err := suite.App.LockupKeeper.CreateSyntheticLockup(suite.Ctx, i, synthlock.SynthDenom, synthlock.Duration, false)
		suite.Require().NoError(err)
		synthlocks = append(synthlocks, synthlock)

		// testing not unlocking synthlock, different duration with underlying
		synthlock.SynthDenom = fmt.Sprintf("synth2/%d", i)
		synthlock.Duration = time.Second * time.Duration(i) / 2
		err = suite.App.LockupKeeper.CreateSyntheticLockup(suite.Ctx, i, synthlock.SynthDenom, synthlock.Duration, false)
		suite.Require().NoError(err)
		synthlocks = append(synthlocks, synthlock)

		// testing unlocking synthlock, different duration with underlying
		synthlock.SynthDenom = fmt.Sprintf("synth3/%d", i)
		err = suite.App.LockupKeeper.CreateSyntheticLockup(suite.Ctx, i, synthlock.SynthDenom, synthlock.Duration, true)
		suite.Require().NoError(err)
		synthlocks = append(synthlocks, synthlock)
	}

	// check synthlocks are all set
	checkSynthlocks := func(amounts []uint64) {
		// by GetAllSyntheticLockups
		for i, synthlock := range suite.App.LockupKeeper.GetAllSyntheticLockups(suite.Ctx) {
			suite.Require().Equal(synthlock, synthlocks[i])
		}
		// by GetAllSyntheticLockupsByLockup
		for i := uint64(1); i <= 3; i++ {
			for j, synthlockByLockup := range suite.App.LockupKeeper.GetAllSyntheticLockupsByLockup(suite.Ctx, i) {
				suite.Require().Equal(synthlockByLockup, synthlocks[(int(i)-1)*3+j])
			}
		}
		// by GetAllSyntheticLockupsByAddr
		for i, synthlock := range suite.App.LockupKeeper.GetAllSyntheticLockupsByAddr(suite.Ctx, addr1) {
			suite.Require().Equal(synthlock, synthlocks[i])
		}
		// by GetPeriodLocksAccumulation
		for i := 1; i <= 3; i++ {
			for j := 1; j <= 3; j++ {
				// get accumulation with always-qualifying condition
				acc := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
					Denom:    fmt.Sprintf("synth%d/%d", j, i),
					Duration: time.Second / 10,
				})
				// amount retrieved should be equal with underlying lock's locked amount
				suite.Require().Equal(acc.Int64(), amounts[i])

				// get accumulation with non-qualifying condition
				acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
					Denom:    fmt.Sprintf("synth%d/%d", j, i),
					Duration: time.Second * 100,
				})
				suite.Require().Equal(acc.Int64(), 0)
			}
		}
	}

	checkSynthlocks([]uint64{10, 20, 30})

	// call AddTokensToLock
	for i := uint64(1); i <= 3; i++ {
		coins := sdk.NewInt64Coin("stake", int64(i)*10)
		suite.FundAcc(addr1, sdk.Coins{coins})
		_, err := suite.App.LockupKeeper.AddTokensToLockByID(suite.Ctx, i, addr1, coins)
		suite.Require().NoError(err)
	}

	// check if all invariants holds after calling AddTokensToLock
	checkSynthlocks([]uint64{20, 40, 60})
}

func (suite *KeeperTestSuite) TestEndblockerWithdrawAllMaturedLockups() {
	suite.SetupTest()

	addr1 := sdk.AccAddress([]byte("addr1---------------"))
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	totalCoins := coins.Add(coins...).Add(coins...)

	// lock coins for 5 second, 1 seconds, and 3 seconds in that order
	times := []time.Duration{time.Second * 5, time.Second, time.Second * 3}
	sortedTimes := []time.Duration{time.Second, time.Second * 3, time.Second * 5}
	sortedTimesIndex := []uint64{2, 3, 1}
	unbondBlockTimes := make([]time.Time, len(times))

	// setup locks for 5 second, 1 second, and 3 seconds, and begin unbonding them.
	setupInitLocks := func() {
		for i := 0; i < len(times); i++ {
			unbondBlockTimes[i] = suite.Ctx.BlockTime().Add(sortedTimes[i])
		}

		for i := 0; i < len(times); i++ {
			suite.LockTokens(addr1, coins, times[i])
		}

		// consistency check locks
		locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
		suite.Require().NoError(err)
		suite.Require().Len(locks, 3)
		for i := 0; i < len(times); i++ {
			suite.Require().Equal(locks[i].EndTime, time.Time{})
			suite.Require().Equal(locks[i].IsUnlocking(), false)
		}

		// begin unlock
		locks, err = suite.App.LockupKeeper.BeginUnlockAllNotUnlockings(suite.Ctx, addr1)
		unlockedCoins := suite.App.LockupKeeper.GetCoinsFromLocks(locks)
		suite.Require().NoError(err)
		suite.Require().Len(locks, len(times))
		suite.Require().Equal(unlockedCoins, totalCoins)
		for i := 0; i < len(times); i++ {
			suite.Require().Equal(sortedTimesIndex[i], locks[i].ID)
		}

		// check locks, these should now be sorted by unbonding completion time
		locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
		suite.Require().NoError(err)
		suite.Require().Len(locks, 3)
		for i := 0; i < 3; i++ {
			suite.Require().NotEqual(locks[i].EndTime, time.Time{})
			suite.Require().Equal(locks[i].EndTime, unbondBlockTimes[i])
			suite.Require().Equal(locks[i].IsUnlocking(), true)
		}
	}
	setupInitLocks()

	// try withdrawing before mature
	suite.App.LockupKeeper.WithdrawAllMaturedLocks(suite.Ctx)
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 3)

	// withdraw at 1 sec, 3 sec, and 5 sec intervals, check automatically withdrawn
	for i := 0; i < len(times); i++ {
		suite.App.LockupKeeper.WithdrawAllMaturedLocks(suite.Ctx.WithBlockTime(unbondBlockTimes[i]))
		locks, err = suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
		suite.Require().NoError(err)
		suite.Require().Len(locks, len(times)-i-1)
	}
	suite.Require().Equal(suite.App.BankKeeper.GetAccountsBalances(suite.Ctx)[1].Address, addr1.String())
	suite.Require().Equal(suite.App.BankKeeper.GetAccountsBalances(suite.Ctx)[1].Coins, totalCoins)

	suite.SetupTest()
	setupInitLocks()
	// now withdraw all locks and ensure all got withdrawn
	suite.App.LockupKeeper.WithdrawAllMaturedLocks(suite.Ctx.WithBlockTime(unbondBlockTimes[len(times)-1]))
	suite.Require().Len(locks, 0)
}

func (suite *KeeperTestSuite) TestLockAccumulationStore() {
	suite.SetupTest()

	// initial check
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr := sdk.AccAddress([]byte("addr1---------------"))

	// 1 * time.Second: 10 + 20
	// 2 * time.Second: 20 + 30
	// 3 * time.Second: 30
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr, coins, time.Second)
	coins = sdk.Coins{sdk.NewInt64Coin("stake", 20)}
	suite.LockTokens(addr, coins, time.Second)
	suite.LockTokens(addr, coins, time.Second*2)
	coins = sdk.Coins{sdk.NewInt64Coin("stake", 30)}
	suite.LockTokens(addr, coins, time.Second*2)
	suite.LockTokens(addr, coins, time.Second*3)

	// check accumulations
	acc := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: 0,
	})
	suite.Require().Equal(int64(110), acc.Int64())
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second * 1,
	})
	suite.Require().Equal(int64(110), acc.Int64())
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second * 2,
	})
	suite.Require().Equal(int64(80), acc.Int64())
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second * 3,
	})
	suite.Require().Equal(int64(30), acc.Int64())
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second * 4,
	})
	suite.Require().Equal(int64(0), acc.Int64())
}

func (suite *KeeperTestSuite) TestSlashTokensFromLockByID() {
	suite.SetupTest()

	// initial check
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr := sdk.AccAddress([]byte("addr1---------------"))

	// 1 * time.Second: 10
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr, coins, time.Second)

	// check accumulations
	acc := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second,
	})
	suite.Require().Equal(int64(10), acc.Int64())

	suite.App.LockupKeeper.SlashTokensFromLockByID(suite.Ctx, 1, sdk.Coins{sdk.NewInt64Coin("stake", 1)})
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second,
	})
	suite.Require().Equal(int64(9), acc.Int64())

	lock, err := suite.App.LockupKeeper.GetLockByID(suite.Ctx, 1)
	suite.Require().NoError(err)
	suite.Require().Equal(lock.Coins.String(), "9stake")

	_, err = suite.App.LockupKeeper.SlashTokensFromLockByID(suite.Ctx, 1, sdk.Coins{sdk.NewInt64Coin("stake", 11)})
	suite.Require().Error(err)

	_, err = suite.App.LockupKeeper.SlashTokensFromLockByID(suite.Ctx, 1, sdk.Coins{sdk.NewInt64Coin("stake1", 1)})
	suite.Require().Error(err)
}

func (suite *KeeperTestSuite) TestEditLockup() {
	suite.SetupTest()

	// initial check
	locks, err := suite.App.LockupKeeper.GetPeriodLocks(suite.Ctx)
	suite.Require().NoError(err)
	suite.Require().Len(locks, 0)

	// lock coins
	addr := sdk.AccAddress([]byte("addr1---------------"))

	// 1 * time.Second: 10
	coins := sdk.Coins{sdk.NewInt64Coin("stake", 10)}
	suite.LockTokens(addr, coins, time.Second)

	// check accumulations
	acc := suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second,
	})
	suite.Require().Equal(int64(10), acc.Int64())

	lock, _ := suite.App.LockupKeeper.GetLockByID(suite.Ctx, 1)

	// duration decrease should fail
	err = suite.App.LockupKeeper.ExtendLockup(suite.Ctx, lock.ID, addr, time.Second/2)
	suite.Require().Error(err)
	// extending lock with same duration should fail
	err = suite.App.LockupKeeper.ExtendLockup(suite.Ctx, lock.ID, addr, time.Second)
	suite.Require().Error(err)

	// duration increase should success
	err = suite.App.LockupKeeper.ExtendLockup(suite.Ctx, lock.ID, addr, time.Second*2)
	suite.Require().NoError(err)

	// check queries
	lock, _ = suite.App.LockupKeeper.GetLockByID(suite.Ctx, lock.ID)
	suite.Require().Equal(lock.Duration, time.Second*2)
	suite.Require().Equal(uint64(1), lock.ID)
	suite.Require().Equal(coins, lock.Coins)

	locks = suite.App.LockupKeeper.GetLocksLongerThanDurationDenom(suite.Ctx, "stake", time.Second)
	suite.Require().Equal(len(locks), 1)

	locks = suite.App.LockupKeeper.GetLocksLongerThanDurationDenom(suite.Ctx, "stake", time.Second*2)
	suite.Require().Equal(len(locks), 1)

	// check accumulations
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second,
	})
	suite.Require().Equal(int64(10), acc.Int64())
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second * 2,
	})
	suite.Require().Equal(int64(10), acc.Int64())
	acc = suite.App.LockupKeeper.GetPeriodLocksAccumulation(suite.Ctx, types.QueryCondition{
		Denom:    "stake",
		Duration: time.Second * 3,
	})
	suite.Require().Equal(int64(0), acc.Int64())
}
