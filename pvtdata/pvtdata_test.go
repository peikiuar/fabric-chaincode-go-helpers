package pvtdata

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/braduf/fabric-chaincode-go-helpers/mocking"
)

func TestGetTransientDataValue(t *testing.T) {
	t.Run("get existing string field", func(t *testing.T) {
		transientFieldName := "field"
		transientValue := []byte("value")
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = transientValue
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValue", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		got, err := GetTransientDataValue(mockTransactionContext, transientFieldName)
		mockStub.MockTransactionEnd("1")

		assertError(t, err, nil)
		if !reflect.DeepEqual(got, transientValue) {
			t.Errorf("got %q want %q", got, transientValue)
		}
	})

	t.Run("wrong transient field name", func(t *testing.T) {
		transientFieldName := "field"
		transientValue := "value"
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = []byte(transientValue)
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValue", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		_, err := GetTransientDataValue(mockTransactionContext, "otherFieldName")
		mockStub.MockTransactionEnd("1")

		assertError(t, err, ErrWrongTransientFieldName)
	})

	t.Run("empty transient value", func(t *testing.T) {
		transientFieldName := "field"
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = []byte("")
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValue", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		_, err := GetTransientDataValue(mockTransactionContext, transientFieldName)
		mockStub.MockTransactionEnd("1")

		assertError(t, err, ErrEmptyTransientFieldValue)
	})
}

func TestGetTransientDataValueUnmarshaled(t *testing.T) {
	t.Run("get existing JSON string field", func(t *testing.T) {
		type mockDataValue struct {
			Key string `json:"key"`
		}
		transientFieldName := "field"
		transientFieldValue := mockDataValue{"value"}
		transientMapValue, _ := json.Marshal(transientFieldValue)
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = transientMapValue
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValue", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		var got mockDataValue
		err := GetTransientDataValueUnmarshaled(mockTransactionContext, transientFieldName, &got)
		mockStub.MockTransactionEnd("1")

		assertError(t, err, nil)
		if got != transientFieldValue {
			t.Errorf("got %q want %q", got, transientFieldValue)
		}
	})
}

func TestPutImplicitPrivateData(t *testing.T) {

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
