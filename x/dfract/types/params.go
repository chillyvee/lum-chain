package types

import (
	"fmt"
	"strings"

	sdktypes "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys.
var (
	KeyDepositDenom     = []byte("DepositDenom")
	KeyMinDepositAmount = []byte("MinDepositAmount")
)

// ParamKeyTable for dfract module.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams return the default dfract module params
func DefaultParams() Params {
	return Params{
		DepositDenom:     "ibc/05554A9BFDD28894D7F18F4C707AA0930D778751A437A9FE1F4684A3E1199728", // USDC ibc denom from Osmosis to Lum Network mainnet
		MinDepositAmount: 1000000,
	}
}

func (p *Params) Validate() error {
	if err := validateDepositDenom(p.DepositDenom); err != nil {
		return err
	}

	if err := validateMinDepositAmount(p.MinDepositAmount); err != nil {
		return err
	}
	return nil
}

func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyDepositDenom, &p.DepositDenom, validateDepositDenom),
		paramtypes.NewParamSetPair(KeyMinDepositAmount, &p.MinDepositAmount, validateMinDepositAmount),
	}
}

func validateDepositDenom(i interface{}) error {
	v, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if strings.TrimSpace(v) == "" {
		return ErrInvalidDepositDenom
	}
	if err := sdktypes.ValidateDenom(v); err != nil {
		return err
	}
	return nil
}

func validateMinDepositAmount(i interface{}) error {
	v, ok := i.(uint32)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}
	if v <= 0 {
		return ErrInvalidMinDepositAmount
	}
	return nil
}
