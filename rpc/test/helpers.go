package rpctest

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	abci "github.com/DeltaChain/DeltaChain/abci/types"
	"github.com/DeltaChain/DeltaChain/libs/log"

	cfg "github.com/DeltaChain/DeltaChain/config"
	tmnet "github.com/DeltaChain/DeltaChain/libs/net"
	nm "github.com/DeltaChain/DeltaChain/node"
	"github.com/DeltaChain/DeltaChain/p2p"
	"github.com/DeltaChain/DeltaChain/privval"
	"github.com/DeltaChain/DeltaChain/proxy"
	ctypes "github.com/DeltaChain/DeltaChain/rpc/core/types"
	core_grpc "github.com/DeltaChain/DeltaChain/rpc/grpc"
	rpcclient "github.com/DeltaChain/DeltaChain/rpc/jsonrpc/client"
)

// Options helps with specifying some parameters for our RPC testing for greater
// control.
type Options struct {
	suppressStdout bool
	recreateConfig bool
}

var globalConfig *cfg.Config
var defaultOptions = Options{
	suppressStdout: false,
	recreateConfig: false,
}

func waitForRPC() {
	laddr := GetConfig().RPC.ListenAddress
	client, err := rpcclient.New(laddr)
	if err != nil {
		panic(err)
	}
	result := new(ctypes.ResultStatus)
	for {
		_, err := client.Call(context.Background(), "status", map[string]interface{}{}, result)
		if err == nil {
			return
		}

		fmt.Println("error", err)
		time.Sleep(time.Millisecond)
	}
}

func waitForGRPC() {
	client := GetGRPCClient()
	for {
		_, err := client.Ping(context.Background(), &core_grpc.RequestPing{})
		if err == nil {
			return
		}
	}
}

// f**ing long, but unique for each test
func makePathname() string {
	// get path
	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	// fmt.Println(p)
	sep := string(filepath.Separator)
	return strings.ReplaceAll(p, sep, "_")
}

func randPort() int {
	port, err := tmnet.GetFreePort()
	if err != nil {
		panic(err)
	}
	return port
}

func makeAddrs() (string, string, string) {
	return fmt.Sprintf("tcp://127.0.0.1:%d", randPort()),
		fmt.Sprintf("tcp://127.0.0.1:%d", randPort()),
		fmt.Sprintf("tcp://127.0.0.1:%d", randPort())
}

func createConfig() *cfg.Config {
	pathname := makePathname()
	c := cfg.ResetTestRoot(pathname)

	// and we use random ports to run in parallel
	tm, rpc, grpc := makeAddrs()
	c.P2P.ListenAddress = tm
	c.RPC.ListenAddress = rpc
	c.RPC.CORSAllowedOrigins = []string{"https://DeltaChain.com/"}
	c.RPC.GRPCListenAddress = grpc
	return c
}

// GetConfig returns a config for the test cases as a singleton
func GetConfig(forceCreate ...bool) *cfg.Config {
	if globalConfig == nil || (len(forceCreate) > 0 && forceCreate[0]) {
		globalConfig = createConfig()
	}
	return globalConfig
}

func GetGRPCClient() core_grpc.BroadcastAPIClient {
	grpcAddr := globalConfig.RPC.GRPCListenAddress
	return core_grpc.StartGRPCClient(grpcAddr)
}

// StartDeltaChain starts a test DeltaChain server in a go routine and returns when it is initialized
func StartDeltaChain(app abci.Application, opts ...func(*Options)) *nm.Node {
	nodeOpts := defaultOptions
	for _, opt := range opts {
		opt(&nodeOpts)
	}
	node := NewDeltaChain(app, &nodeOpts)
	err := node.Start()
	if err != nil {
		panic(err)
	}

	// wait for rpc
	waitForRPC()
	waitForGRPC()

	if !nodeOpts.suppressStdout {
		fmt.Println("DeltaChain running!")
	}

	return node
}

// StopDeltaChain stops a test DeltaChain server, waits until it's stopped and
// cleans up test/config files.
func StopDeltaChain(node *nm.Node) {
	if err := node.Stop(); err != nil {
		node.Logger.Error("Error when trying to stop node", "err", err)
	}
	node.Wait()
	os.RemoveAll(node.Config().RootDir)
}

// NewDeltaChain creates a new DeltaChain server and sleeps forever
func NewDeltaChain(app abci.Application, opts *Options) *nm.Node {
	// Create & start node
	config := GetConfig(opts.recreateConfig)
	var logger log.Logger
	if opts.suppressStdout {
		logger = log.NewNopLogger()
	} else {
		logger = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
		logger = log.NewFilter(logger, log.AllowError())
	}
	pvKeyFile := config.PrivValidatorKeyFile()
	pvKeyStateFile := config.PrivValidatorStateFile()
	pv := privval.LoadOrGenFilePV(pvKeyFile, pvKeyStateFile)
	papp := proxy.NewLocalClientCreator(app)
	nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
	if err != nil {
		panic(err)
	}
	node, err := nm.NewNode(config, pv, nodeKey, papp,
		nm.DefaultGenesisDocProviderFunc(config),
		nm.DefaultDBProvider,
		nm.DefaultMetricsProvider(config.Instrumentation),
		logger)
	if err != nil {
		panic(err)
	}
	return node
}

// SuppressStdout is an option that tries to make sure the RPC test DeltaChain
// node doesn't log anything to stdout.
func SuppressStdout(o *Options) {
	o.suppressStdout = true
}

// RecreateConfig instructs the RPC test to recreate the configuration each
// time, instead of treating it as a global singleton.
func RecreateConfig(o *Options) {
	o.recreateConfig = true
}
