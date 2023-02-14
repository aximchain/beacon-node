package scripts

import (
	sdk "github.com/aximchain/axc-cosmos-sdk/types"
	"github.com/aximchain/axc-cosmos-sdk/x/auth"

	cmntypes "github.com/aximchain/flash-node/common/types"
)

func isFlagEnabled(ctx sdk.Context, am auth.AccountKeeper, addr sdk.AccAddress, targetFlag uint64) bool {
	acc := am.GetAccount(ctx, addr)
	if acc == nil {
		return false
	}
	account, ok := acc.(cmntypes.NamedAccount)
	if !ok {
		return false
	}
	if account.GetFlags()&targetFlag == 0 {
		return false
	}
	return true
}
