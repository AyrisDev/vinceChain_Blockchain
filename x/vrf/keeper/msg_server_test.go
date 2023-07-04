package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/AyrisDev/vinceChain_Blockchain/testutil/keeper"
	"github.com/AyrisDev/vinceChain_Blockchain/x/vrf/keeper"
	"github.com/AyrisDev/vinceChain_Blockchain/x/vrf/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.VRFKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
