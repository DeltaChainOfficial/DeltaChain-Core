package core

import (
	ctypes "github.com/DeltaChain/DeltaChain/rpc/core/types"
	rpctypes "github.com/DeltaChain/DeltaChain/rpc/jsonrpc/types"
)

// UnsafeFlushMempool removes all transactions from the mempool.
func UnsafeFlushMempool(ctx *rpctypes.Context) (*ctypes.ResultUnsafeFlushMempool, error) {
	env.Mempool.Flush()
	return &ctypes.ResultUnsafeFlushMempool{}, nil
}
