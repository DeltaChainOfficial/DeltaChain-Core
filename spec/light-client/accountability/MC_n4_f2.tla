----------------------------- MODULE MC_n4_f2 -------------------------------
CONSTANT Proposer \* the proposer function from 0..NRounds to 1..N

\* the variables declared in DeltaChainAcc3
VARIABLES
  round, step, decision, lockedValue, lockedRound, validValue, validRound,
  msgsPropose, msgsPrevote, msgsPrecommit, evidence, action

INSTANCE DeltaChainAccDebug_004_draft WITH
  Corr <- {"c1", "c2"},
  Faulty <- {"f3", "f4"},
  N <- 4,
  T <- 1,
  ValidValues <- { "v0", "v1" },
  InvalidValues <- {"v2"},
  MaxRound <- 2

\* run Apalache with --cinit=ConstInit
ConstInit == \* the proposer is arbitrary -- works for safety
  Proposer \in [Rounds -> AllProcs]

=============================================================================    
