package types

import (
	"fmt"

	sdk "github.com/aximchain/axc-cosmos-sdk/types"

	cmmtypes "github.com/aximchain/flash-node/common/types"
)

func ConvertAXCAmountToBCAmountBigInt(contractDecimals int8, axcAmount sdk.Int) (sdk.Int, sdk.Error) {
	if contractDecimals == cmmtypes.TokenDecimals {
		return axcAmount, nil
	}

	var bcAmount sdk.Int
	if contractDecimals >= cmmtypes.TokenDecimals {
		decimals := sdk.NewIntWithDecimal(1, int(contractDecimals-cmmtypes.TokenDecimals))
		if !axcAmount.Mod(decimals).IsZero() {
			return sdk.Int{}, ErrInvalidAmount(fmt.Sprintf("can't convert bep2(decimals: 8) axcAmount to ERC20(decimals: %d) axcAmount", contractDecimals))
		}
		bcAmount = axcAmount.Div(decimals)
	} else {
		decimals := sdk.NewIntWithDecimal(1, int(cmmtypes.TokenDecimals-contractDecimals))
		bcAmount = axcAmount.Mul(decimals)
	}
	return bcAmount, nil
}

func ConvertAXCAmountToBCAmount(contractDecimals int8, axcAmount sdk.Int) (int64, sdk.Error) {
	res, err := ConvertAXCAmountToBCAmountBigInt(contractDecimals, axcAmount)
	if err != nil {
		return 0, err
	}
	// since we only convert axc amount in transfer out package to bc amount,
	// so it should not overflow
	return res.Int64(), nil
}

func ConvertBCAmountToAXCAmount(contractDecimals int8, bcAmount int64) (sdk.Int, sdk.Error) {
	if contractDecimals == cmmtypes.TokenDecimals {
		return sdk.NewInt(bcAmount), nil
	}

	var axcAmount sdk.Int
	if contractDecimals >= cmmtypes.TokenDecimals {
		decimals := sdk.NewIntWithDecimal(1, int(contractDecimals-cmmtypes.TokenDecimals))
		axcAmount = sdk.NewInt(bcAmount).Mul(decimals)
	} else {
		decimals := sdk.NewIntWithDecimal(1, int(cmmtypes.TokenDecimals-contractDecimals))
		if !sdk.NewInt(bcAmount).Mod(decimals).IsZero() {
			return sdk.Int{}, ErrInvalidAmount(fmt.Sprintf("can't convert bep2(decimals: 8) amount to ERC20(decimals: %d) amount", contractDecimals))
		}
		axcAmount = sdk.NewInt(bcAmount).Div(decimals)
	}
	return axcAmount, nil
}
