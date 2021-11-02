// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"github.com/axelarnetwork/axelar-core/x/reward/exported"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"sync"
)

// Ensure, that RewardPoolMock does implement exported.RewardPool.
// If this is not the case, regenerate this file with moq.
var _ exported.RewardPool = &RewardPoolMock{}

// RewardPoolMock is a mock implementation of exported.RewardPool.
//
// 	func TestSomethingThatUsesRewardPool(t *testing.T) {
//
// 		// make and configure a mocked exported.RewardPool
// 		mockedRewardPool := &RewardPoolMock{
// 			AddRewardFunc: func(valAddress sdk.ValAddress, coin sdk.Coin)  {
// 				panic("mock out the AddReward method")
// 			},
// 			ClearRewardsFunc: func(valAddress sdk.ValAddress)  {
// 				panic("mock out the ClearRewards method")
// 			},
// 			ReleaseRewardsFunc: func(valAddress sdk.ValAddress) error {
// 				panic("mock out the ReleaseRewards method")
// 			},
// 		}
//
// 		// use mockedRewardPool in code that requires exported.RewardPool
// 		// and then make assertions.
//
// 	}
type RewardPoolMock struct {
	// AddRewardFunc mocks the AddReward method.
	AddRewardFunc func(valAddress sdk.ValAddress, coin sdk.Coin)

	// ClearRewardsFunc mocks the ClearRewards method.
	ClearRewardsFunc func(valAddress sdk.ValAddress)

	// ReleaseRewardsFunc mocks the ReleaseRewards method.
	ReleaseRewardsFunc func(valAddress sdk.ValAddress) error

	// calls tracks calls to the methods.
	calls struct {
		// AddReward holds details about calls to the AddReward method.
		AddReward []struct {
			// ValAddress is the valAddress argument value.
			ValAddress sdk.ValAddress
			// Coin is the coin argument value.
			Coin sdk.Coin
		}
		// ClearRewards holds details about calls to the ClearRewards method.
		ClearRewards []struct {
			// ValAddress is the valAddress argument value.
			ValAddress sdk.ValAddress
		}
		// ReleaseRewards holds details about calls to the ReleaseRewards method.
		ReleaseRewards []struct {
			// ValAddress is the valAddress argument value.
			ValAddress sdk.ValAddress
		}
	}
	lockAddReward      sync.RWMutex
	lockClearRewards   sync.RWMutex
	lockReleaseRewards sync.RWMutex
}

// AddReward calls AddRewardFunc.
func (mock *RewardPoolMock) AddReward(valAddress sdk.ValAddress, coin sdk.Coin) {
	if mock.AddRewardFunc == nil {
		panic("RewardPoolMock.AddRewardFunc: method is nil but RewardPool.AddReward was just called")
	}
	callInfo := struct {
		ValAddress sdk.ValAddress
		Coin       sdk.Coin
	}{
		ValAddress: valAddress,
		Coin:       coin,
	}
	mock.lockAddReward.Lock()
	mock.calls.AddReward = append(mock.calls.AddReward, callInfo)
	mock.lockAddReward.Unlock()
	mock.AddRewardFunc(valAddress, coin)
}

// AddRewardCalls gets all the calls that were made to AddReward.
// Check the length with:
//     len(mockedRewardPool.AddRewardCalls())
func (mock *RewardPoolMock) AddRewardCalls() []struct {
	ValAddress sdk.ValAddress
	Coin       sdk.Coin
} {
	var calls []struct {
		ValAddress sdk.ValAddress
		Coin       sdk.Coin
	}
	mock.lockAddReward.RLock()
	calls = mock.calls.AddReward
	mock.lockAddReward.RUnlock()
	return calls
}

// ClearRewards calls ClearRewardsFunc.
func (mock *RewardPoolMock) ClearRewards(valAddress sdk.ValAddress) {
	if mock.ClearRewardsFunc == nil {
		panic("RewardPoolMock.ClearRewardsFunc: method is nil but RewardPool.ClearRewards was just called")
	}
	callInfo := struct {
		ValAddress sdk.ValAddress
	}{
		ValAddress: valAddress,
	}
	mock.lockClearRewards.Lock()
	mock.calls.ClearRewards = append(mock.calls.ClearRewards, callInfo)
	mock.lockClearRewards.Unlock()
	mock.ClearRewardsFunc(valAddress)
}

// ClearRewardsCalls gets all the calls that were made to ClearRewards.
// Check the length with:
//     len(mockedRewardPool.ClearRewardsCalls())
func (mock *RewardPoolMock) ClearRewardsCalls() []struct {
	ValAddress sdk.ValAddress
} {
	var calls []struct {
		ValAddress sdk.ValAddress
	}
	mock.lockClearRewards.RLock()
	calls = mock.calls.ClearRewards
	mock.lockClearRewards.RUnlock()
	return calls
}

// ReleaseRewards calls ReleaseRewardsFunc.
func (mock *RewardPoolMock) ReleaseRewards(valAddress sdk.ValAddress) error {
	if mock.ReleaseRewardsFunc == nil {
		panic("RewardPoolMock.ReleaseRewardsFunc: method is nil but RewardPool.ReleaseRewards was just called")
	}
	callInfo := struct {
		ValAddress sdk.ValAddress
	}{
		ValAddress: valAddress,
	}
	mock.lockReleaseRewards.Lock()
	mock.calls.ReleaseRewards = append(mock.calls.ReleaseRewards, callInfo)
	mock.lockReleaseRewards.Unlock()
	return mock.ReleaseRewardsFunc(valAddress)
}

// ReleaseRewardsCalls gets all the calls that were made to ReleaseRewards.
// Check the length with:
//     len(mockedRewardPool.ReleaseRewardsCalls())
func (mock *RewardPoolMock) ReleaseRewardsCalls() []struct {
	ValAddress sdk.ValAddress
} {
	var calls []struct {
		ValAddress sdk.ValAddress
	}
	mock.lockReleaseRewards.RLock()
	calls = mock.calls.ReleaseRewards
	mock.lockReleaseRewards.RUnlock()
	return calls
}