package panictest

import (
	"math/big"
	"testing"

	test "github.com/multiversx/mx-chain-vm-v1_3-go/testcommon"
	"github.com/stretchr/testify/require"
)

const increment = "increment"

func TestExecution_PanicInGoWithSilentWasmer_SIGSEGV(t *testing.T) {
	code := test.GetTestSCCode("counter", "../../../")
	host, blockchain := test.DefaultTestArwenForCallSigSegv(t, code, big.NewInt(1), true)

	blockchain.GetStorageDataCalled = func(_ []byte, _ []byte) ([]byte, uint32, error) {
		var i *int
		i = nil

		// dereference a nil pointer
		*i = *i + 1
		return nil, 0, nil
	}

	input := test.CreateTestContractCallInputBuilder().
		WithGasProvided(1000000).
		WithFunction(increment).
		Build()

	// Ensure that no more panic
	defer func() {
		r := recover()
		require.Nil(t, r)
	}()

	expectedError := "runtime error: invalid memory address or nil pointer dereference"

	_, err := host.RunSmartContractCall(input)
	require.Equal(t, expectedError, err.Error())
}
