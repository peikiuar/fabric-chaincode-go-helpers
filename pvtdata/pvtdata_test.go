package pvtdata

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/braduf/fabric-chaincode-go-helpers/mocking"
)

func TestGetTransientDataValue(t *testing.T) {
	t.Run("get existing string field", func(t *testing.T) {
		transientFieldName := "field"
		transientValue := "value"
		transientMapValue, _ := json.Marshal(transientValue)
		mockTransient := map[string][]byte{transientFieldName: transientMapValue}
		fmt.Println(mockTransient)
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValue", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		var got string
		err := GetTransientDataValue(mockTransactionContext, transientFieldName, &got)
		mockStub.MockTransactionEnd("1")

		assertError(t, err, nil)
		if got != transientValue {
			t.Errorf("got %q want %q", got, transientValue)
		}
	})
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
