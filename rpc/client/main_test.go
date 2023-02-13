package client_test

import (
	"os"
	"testing"

	"github.com/DeltaChain/DeltaChain/abci/example/kvstore"
	nm "github.com/DeltaChain/DeltaChain/node"
	rpctest "github.com/DeltaChain/DeltaChain/rpc/test"
)

var node *nm.Node

func TestMain(m *testing.M) {
	// start a DeltaChain node (and kvstore) in the background to test against
	dir, err := os.MkdirTemp("/tmp", "rpc-client-test")
	if err != nil {
		panic(err)
	}

	app := kvstore.NewPersistentKVStoreApplication(dir)
	node = rpctest.StartDeltaChain(app)

	code := m.Run()

	// and shut down proper at the end
	rpctest.StopDeltaChain(node)
	_ = os.RemoveAll(dir)
	os.Exit(code)
}
