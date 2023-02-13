package core

import (
	"errors"
	"fmt"

	ctypes "github.com/DeltaChain/DeltaChain/rpc/core/types"
	rpctypes "github.com/DeltaChain/DeltaChain/rpc/jsonrpc/types"
	"github.com/DeltaChain/DeltaChain/types"
)

// BroadcastEvidence broadcasts evidence of the misbehavior.
// More: https://docs.DeltaChain.com/v0.34/rpc/#/Info/broadcast_evidence
func BroadcastEvidence(ctx *rpctypes.Context, ev types.Evidence) (*ctypes.ResultBroadcastEvidence, error) {
	if ev == nil {
		return nil, errors.New("no evidence was provided")
	}

	if err := ev.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("evidence.ValidateBasic failed: %w", err)
	}

	if err := env.EvidencePool.AddEvidence(ev); err != nil {
		return nil, fmt.Errorf("failed to add evidence: %w", err)
	}
	return &ctypes.ResultBroadcastEvidence{Hash: ev.Hash()}, nil
}
