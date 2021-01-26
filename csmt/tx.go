package csmt

import (
	"github.com/olympus-protocol/ogen/pkg/chainhash"
)

// ShardTransaction is a single transaction on a shard chain.
type ShardTransaction struct {
	TransactionData []byte
}

// Copy returns a copy of the shard transaction.
func (st *ShardTransaction) Copy() ShardTransaction {
	tx := ShardTransaction{
		TransactionData: make([]byte, len(st.TransactionData)),
	}

	copy(tx.TransactionData, st.TransactionData)

	return tx
}

// CombineHashes combines two branches of the tree.
func CombineHashes(left *chainhash.Hash, right *chainhash.Hash) chainhash.Hash {
	return chainhash.HashH(append(left[:], right[:]...))
}

// EmptyTrees are empty trees for each level.
var EmptyTrees [256]chainhash.Hash

// EmptyTree is the hash of an empty tree.
var EmptyTree = chainhash.Hash{}

func init() {
	EmptyTrees[0] = chainhash.Hash{}
	for i := range EmptyTrees[1:] {
		EmptyTrees[i+1] = CombineHashes(&EmptyTrees[i], &EmptyTrees[i])
	}

	EmptyTree = EmptyTrees[255]
}

// UpdateWitness allows an executor to securely update the tree root so that only a single key is changed.
type UpdateWitness struct {
	Key             chainhash.Hash
	OldValue        chainhash.Hash
	NewValue        chainhash.Hash
	WitnessBitfield chainhash.Hash
	LastLevel       uint8
	Witnesses       []chainhash.Hash
}

// Copy returns a copy of the update witness.
func (uw *UpdateWitness) Copy() UpdateWitness {
	newUw := *uw

	newUw.Witnesses = make([]chainhash.Hash, len(uw.Witnesses))
	for i := range newUw.Witnesses {
		copy(newUw.Witnesses[i][:], uw.Witnesses[i][:])
	}

	return newUw
}

// VerificationWitness allows an executor to verify a specific node in the tree.
type VerificationWitness struct {
	Key             chainhash.Hash
	Value           chainhash.Hash
	WitnessBitfield chainhash.Hash
	Witnesses       []chainhash.Hash
	LastLevel       uint8
}

// Copy returns a copy of the update witness.
func (vw *VerificationWitness) Copy() VerificationWitness {
	newVw := *vw

	newVw.Witnesses = make([]chainhash.Hash, len(vw.Witnesses))
	for i := range newVw.Witnesses {
		copy(newVw.Witnesses[i][:], vw.Witnesses[i][:])
	}

	return newVw
}

// TransactionPackage is a way to update the state root without having the entire state.
type TransactionPackage struct {
	StartRoot     chainhash.Hash
	EndRoot       chainhash.Hash
	Updates       []UpdateWitness
	Verifications []VerificationWitness
	Transactions  []ShardTransaction
}

// Copy copies the transaction package.
func (tp *TransactionPackage) Copy() TransactionPackage {
	newTp := *tp

	newTp.Updates = make([]UpdateWitness, len(tp.Updates))
	newTp.Verifications = make([]VerificationWitness, len(tp.Verifications))
	newTp.Transactions = make([]ShardTransaction, len(tp.Transactions))

	for i := range newTp.Updates {
		newTp.Updates[i] = tp.Updates[i].Copy()
	}

	for i := range newTp.Verifications {
		newTp.Verifications[i] = tp.Verifications[i].Copy()
	}

	for i := range newTp.Transactions {
		newTp.Transactions[i] = tp.Transactions[i].Copy()
	}

	return newTp
}
