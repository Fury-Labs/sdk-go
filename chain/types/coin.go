package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// KAI defines the default coin denomination used in Ethermint in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	KaijuCoin string = "kai"

	// BaseDenomUnit defines the base denomination unit for Photons.
	// 1 photon = 1x10^{BaseDenomUnit} kai
	BaseDenomUnit = 18
)

// NewKaijuCoin is a utility function that returns an "kai" coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewKaijuCoin(amount sdk.Int) sdk.Coin {
	return sdk.NewCoin(KaijuCoin, amount)
}

// NewKaijuDecCoin is a utility function that returns an "kai" decimal coin with the given sdk.Int amount.
// The function will panic if the provided amount is negative.
func NewKaijuDecCoin(amount sdk.Int) sdk.DecCoin {
	return sdk.NewDecCoin(KaijuCoin, amount)
}

// NewKaijuCoinInt64 is a utility function that returns an "kai" coin with the given int64 amount.
// The function will panic if the provided amount is negative.
func NewKaijuCoinInt64(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(KaijuCoin, amount)
}
