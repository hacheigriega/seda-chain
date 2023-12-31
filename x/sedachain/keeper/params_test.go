package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	testkeeper "github.com/sedaprotocol/seda-chain/testutil/keeper"
	"github.com/sedaprotocol/seda-chain/x/sedachain/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SedachainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
