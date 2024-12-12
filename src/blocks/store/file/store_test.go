package fileblockstore

// Implement test cases for the FileBlockStore struct.

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/assert"

	"github.com/kkrt-labs/kakarot-controller/pkg/ethereum/rpc"
	blockinputs "github.com/kkrt-labs/kakarot-controller/src/blocks/inputs"
)

func TestFileBlockStore(t *testing.T) {
	baseDir := t.TempDir()
	store := New(baseDir)

	// Test storing and loading preflight data
	preflightData := &blockinputs.PreflightData{
		ChainConfig: &params.ChainConfig{
			ChainID: big.NewInt(1),
		},
		Block: &rpc.Block{
			Header: rpc.Header{
				Number: (*hexutil.Big)(hexutil.MustDecodeBig("0xa")),
			},
		},
	}
	err := store.StorePreflightData(context.Background(), preflightData)
	assert.NoError(t, err)

	_, err = store.LoadPreflightData(context.Background(), 1, 10)
	assert.NoError(t, err)

	// Test loading non-existent preflight data
	_, err = store.LoadPreflightData(context.Background(), 1, 20)
	assert.Error(t, err)

	// Test storing and loading provable inputs
	provableInputs := &blockinputs.ProvableInputs{
		ChainConfig: &params.ChainConfig{
			ChainID: big.NewInt(2),
		},
		Block: &rpc.Block{
			Header: rpc.Header{
				Number: (*hexutil.Big)(hexutil.MustDecodeBig("0xf")),
			},
		},
	}
	err = store.StoreProvableInputs(context.Background(), provableInputs)
	assert.NoError(t, err)

	_, err = store.LoadProvableInputs(context.Background(), 2, 15)
	assert.NoError(t, err)

	// Test loading non-existent provable inputs
	_, err = store.LoadProvableInputs(context.Background(), 2, 25)
	assert.Error(t, err)
}
