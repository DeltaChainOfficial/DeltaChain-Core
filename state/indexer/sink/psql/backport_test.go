package psql

import (
	"github.com/DeltaChain/DeltaChain/state/indexer"
	"github.com/DeltaChain/DeltaChain/state/txindex"
)

var (
	_ indexer.BlockIndexer = BackportBlockIndexer{}
	_ txindex.TxIndexer    = BackportTxIndexer{}
)
