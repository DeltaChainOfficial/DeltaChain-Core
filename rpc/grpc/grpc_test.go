package coregrpc_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/DeltaChain/DeltaChain/abci/example/kvstore"
	core_grpc "github.com/DeltaChain/DeltaChain/rpc/grpc"
	rpctest "github.com/DeltaChain/DeltaChain/rpc/test"
)

func TestMain(m *testing.M) {
	// start a DeltaChain node in the background to test against
	app := kvstore.NewApplication()
	node := rpctest.StartDeltaChain(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopDeltaChain(node)
	os.Exit(code)
}

func TestBroadcastTx(t *testing.T) {
	res, err := rpctest.GetGRPCClient().BroadcastTx(
		context.Background(),
		&core_grpc.RequestBroadcastTx{Tx: []byte("this is a tx")},
	)
	require.NoError(t, err)
	require.EqualValues(t, 0, res.CheckTx.Code)
	require.EqualValues(t, 0, res.DeliverTx.Code)
}
