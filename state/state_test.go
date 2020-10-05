package state

import (
	"testing"

	"github.com/braduf/fabric-chaincode-go-helpers/mocking"
)

func TestPutState(t *testing.T) {
	mockStub := mocking.NewMockChaincodeStub("TestPutState", nil)
	mockStub.MockTransactionStart("tx#1")
	mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

	err := PutState(mockTransactionContext, "test1", struct{ key string }{"value"})

	assertError(t, err, nil)
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatalf("expected to get error %q", want)
	}
}
