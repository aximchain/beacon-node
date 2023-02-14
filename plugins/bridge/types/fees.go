package types

import (
	sdk "github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/types/fees"
)

const (
	BindRelayFeeName        = "crossBindRelayFee"
	UnbindRelayFeeName      = "crossUnbindRelayFee"
	TransferOutRelayFeeName = "crossTransferOutRelayFee"
)

func GetFee(feeName string) (sdk.Fee, sdk.Error) {
	calculator := fees.GetCalculator(feeName)
	if calculator == nil {
		return sdk.Fee{}, ErrFeeNotFound("missing calculator for fee type:" + feeName)
	}
	return calculator(nil), nil
}
