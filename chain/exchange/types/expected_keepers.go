package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/ethereum/go-ethereum/common"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	insurancetypes "github.com/Fury-Labs/sdk-go/chain/insurance/types"
	oracletypes "github.com/Fury-Labs/sdk-go/chain/oracle/types"
	wasmxtypes "github.com/Fury-Labs/sdk-go/chain/wasmx/types"
)

// BankKeeper defines the expected bank keeper methods.
type BankKeeper interface {
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	GetSupply(ctx sdk.Context) sdk.Coin
}

// OracleKeeper defines the expected oracle keeper methods.
type OracleKeeper interface {
	GetPrice(ctx sdk.Context, oracletype oracletypes.OracleType, base string, quote string) *sdk.Dec
	GetCumulativePrice(ctx sdk.Context, oracleType oracletypes.OracleType, base string, quote string) *sdk.Dec
	GetHistoricalPriceRecords(ctx sdk.Context, oracleType oracletypes.OracleType, symbol string, from int64) (entry *oracletypes.PriceRecords, omitted bool)
	GetMixedHistoricalPriceRecords(ctx sdk.Context, baseOracleType, quoteOracleType oracletypes.OracleType, baseSymbol, quoteSymbol string, from int64) (mixed *oracletypes.PriceRecords, ok bool)
	GetStandardDeviationForPriceRecords(priceRecords []*oracletypes.PriceRecord) *sdk.Dec
	GetProviderInfo(ctx sdk.Context, provider string) *oracletypes.ProviderInfo
	GetProviderPrice(ctx sdk.Context, provider, symbol string) *sdk.Dec
	GetProviderPriceState(ctx sdk.Context, provider, symbol string) *oracletypes.ProviderPriceState
}

// InsuranceKeeper defines the expected insurance keeper methods.
type InsuranceKeeper interface {
	// HasInsuranceFund returns true if InsuranceFund for the given marketID exists.
	HasInsuranceFund(ctx sdk.Context, marketID common.Hash) bool
	// GetInsuranceFund returns the insurance fund corresponding to the given marketID.
	GetInsuranceFund(ctx sdk.Context, marketID common.Hash) *insurancetypes.InsuranceFund
	// DepositIntoInsuranceFund increments the insurance fund balance by amount.
	DepositIntoInsuranceFund(ctx sdk.Context, marketID common.Hash, amount sdk.Int) error
	// WithdrawFromInsuranceFund decrements the insurance fund balance by amount and sends
	WithdrawFromInsuranceFund(ctx sdk.Context, marketID common.Hash, amount sdk.Int) error
	// UpdateInsuranceFundOracleParams updates the insurance fund's oracle parameters
	UpdateInsuranceFundOracleParams(ctx sdk.Context, marketID common.Hash, oracleParams *OracleParams) error
}

type GovKeeper interface {
	IterateActiveProposalsQueue(ctx sdk.Context, endTime time.Time, cb func(proposal govtypes.Proposal) (stop bool))
	GetVotingParams(ctx sdk.Context) govtypes.VotingParams
}

type DistributionKeeper interface {
	GetFeePool(ctx sdk.Context) (feePool types.FeePool)
	DistributeFromFeePool(ctx sdk.Context, amount sdk.Coins, receiveAddr sdk.AccAddress) error
	FundCommunityPool(ctx sdk.Context, amount sdk.Coins, sender sdk.AccAddress) error
}

type StakingKeeper interface {
	GetDelegatorDelegations(ctx sdk.Context, delegator sdk.AccAddress, maxRetrieve uint16) (delegations []stakingtypes.Delegation)
	Validator(sdk.Context, sdk.ValAddress) stakingtypes.ValidatorI
}

type WasmViewKeeper interface {
	wasmtypes.ViewKeeper
}

type WasmContractOpsKeeper interface {
	wasmtypes.ContractOpsKeeper
}

type WasmxExecutionKeeper interface {
	KaijuExec(ctx sdk.Context, contractAddress sdk.AccAddress, funds sdk.Coins, msg *wasmxtypes.KaijuExecMsg) ([]byte, error)
}
