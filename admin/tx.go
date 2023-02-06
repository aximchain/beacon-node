package admin

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank"

	"github.com/aximchain/beacon-node/common/runtime"
	"github.com/aximchain/beacon-node/plugins/account"
	"github.com/aximchain/beacon-node/plugins/bridge"
	"github.com/aximchain/beacon-node/plugins/dex/order"
	list "github.com/aximchain/beacon-node/plugins/dex/types"
	"github.com/aximchain/beacon-node/plugins/tokens/burn"
	"github.com/aximchain/beacon-node/plugins/tokens/freeze"
	"github.com/aximchain/beacon-node/plugins/tokens/issue"
	"github.com/aximchain/beacon-node/plugins/tokens/ownership"
	"github.com/aximchain/beacon-node/plugins/tokens/seturi"
	"github.com/aximchain/beacon-node/plugins/tokens/swap"
	"github.com/aximchain/beacon-node/plugins/tokens/timelock"
)

var transferOnlyModeBlackList = []string{
	burn.BurnMsg{}.Type(),
	freeze.FreezeMsg{}.Type(),
	freeze.UnfreezeMsg{}.Type(),
	issue.IssueMsg{}.Type(),
	issue.MintMsg{}.Type(),
	order.NewOrderMsg{}.Type(),
	order.CancelOrderMsg{}.Type(),
	timelock.TimeLockMsg{}.Type(),
	timelock.TimeUnlockMsg{}.Type(),
	timelock.TimeRelockMsg{}.Type(),
	issue.IssueMiniMsg{}.Type(),
	issue.IssueTinyMsg{}.Type(),
	seturi.SetURIMsg{}.Type(),
	list.ListMsg{}.Type(),
	list.ListMiniMsg{}.Type(),
	ownership.TransferOwnershipMsg{}.Type(),
	swap.HTLTMsg{}.Type(),
	swap.DepositHTLTMsg{}.Type(),
	swap.ClaimHTLTMsg{}.Type(),
	swap.RefundHTLTMsg{}.Type(),
	account.SetAccountFlagsMsg{}.Type(),
	bridge.BindMsg{}.Type(),
	bridge.UnbindMsg{}.Type(),
	bridge.TransferOutMsg{}.Type(),
}

var TxBlackList = map[runtime.Mode][]string{
	runtime.TransferOnlyMode: transferOnlyModeBlackList,
	runtime.RecoverOnlyMode:  append(transferOnlyModeBlackList, bank.MsgSend{}.Type()),
}

func TxNotAllowedError() sdk.Error {
	return sdk.ErrInternal(fmt.Sprintf("The tx is not allowed, RunningMode: %v", runtime.GetRunningMode()))
}

func IsTxAllowed(tx sdk.Tx) bool {
	mode := runtime.GetRunningMode()
	if mode == runtime.NormalMode {
		return true
	}

	for _, msg := range tx.GetMsgs() {
		if !isMsgAllowed(msg, mode) {
			return false
		}
	}
	return true
}

func isMsgAllowed(msg sdk.Msg, mode runtime.Mode) bool {
	for _, msgType := range TxBlackList[mode] {
		if msgType == msg.Type() {
			return false
		}
	}

	return true
}
