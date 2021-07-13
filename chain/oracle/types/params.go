package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var _ paramtypes.ParamSet = &Params{}

var LargestDecPrice sdk.Dec = sdk.MustNewDecFromStr("10000000")

// Parameter keys

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams() Params {
	return Params{}
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{}
}

// DefaultParams returns a default set of parameters.
func DefaultParams() Params {
	return Params{}
}

// Validate performs basic validation on auction parameters.
func (p Params) Validate() error {
	return nil
}
