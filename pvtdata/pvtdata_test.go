package pvtdata

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/peikiuar/fabric-chaincode-go-helpers/mocking"
)

func TestGetTransientDataValueBytes(t *testing.T) {
	t.Run("get existing string field", func(t *testing.T) {
		transientFieldName := "field"
		transientValue := []byte("value")
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = transientValue
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValueBytes", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		got, err := GetTransientDataValueBytes(mockTransactionContext, transientFieldName)
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
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValueBytes", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		_, err := GetTransientDataValueBytes(mockTransactionContext, "otherFieldName")
		mockStub.MockTransactionEnd("1")

		assertError(t, err, ErrWrongTransientFieldName)
	})

	t.Run("empty transient value", func(t *testing.T) {
		transientFieldName := "field"
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = []byte("")
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValueBytes", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		_, err := GetTransientDataValueBytes(mockTransactionContext, transientFieldName)
		mockStub.MockTransactionEnd("1")

		assertError(t, err, ErrEmptyTransientFieldValue)
	})
}

func TestGetTransientDataValue(t *testing.T) {
	t.Run("get existing JSON string field", func(t *testing.T) {
		type mockDataValue struct {
			Key string `json:"key"`
		}
		transientFieldName := "field"
		transientFieldValue := mockDataValue{"value"}
		transientMapValue, _ := json.Marshal(transientFieldValue)
		mockTransient := make(map[string][]byte)
		mockTransient[transientFieldName] = transientMapValue
		mockStub := mocking.NewMockChaincodeStub("TestGetTransientDataValueBytes", nil, mockTransient)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart("1")
		var got mockDataValue
		err := GetTransientDataValue(mockTransactionContext, transientFieldName, &got)
		mockStub.MockTransactionEnd("1")

		assertError(t, err, nil)
		if got != transientFieldValue {
			t.Errorf("got %q want %q", got, transientFieldValue)
		}
	})
}

func TestPutImplicitPrivateData(t *testing.T) {
	mockStub := mocking.NewMockChaincodeStub("TestPutImplicitPrivateData", nil, nil)
	mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)
	collectionMSP := "Org1"
	type privateData struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
		Field3 bool   `json:"field3"`
	}
	pvtData := privateData{"value1", 2, true}
	txID := "1"

	mockStub.MockTransactionStart(txID)
	err := PutImplicitPrivateData(mockTransactionContext, collectionMSP, mockStub.TxID, pvtData)
	mockStub.MockTransactionEnd(txID)

	mockStub.MockTransactionStart("2")
	var got privateData
	gotBytes, _ := mockStub.GetPrivateData(ImplicitCollectionPrefix+collectionMSP, txID)
	_ = json.Unmarshal(gotBytes, &got)
	mockStub.MockTransactionEnd("2")

	assertError(t, err, nil)
	if !reflect.DeepEqual(got, pvtData) {
		t.Errorf("got %v want %v", got, pvtData)
	}
}

func TestGetImplicitPrivateData(t *testing.T) {
	mockStub := mocking.NewMockChaincodeStub("TestGetImplicitPrivateData", nil, nil)
	mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)
	collectionMSP := "Org1"
	type privateData struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
		Field3 bool   `json:"field3"`
	}
	pvtData := privateData{"value1", 2, true}
	pvtDataBytes, _ := json.Marshal(pvtData)
	txID := "1"

	mockStub.MockTransactionStart(txID)
	_ = mockStub.PutPrivateData(ImplicitCollectionPrefix+collectionMSP, mockStub.TxID, pvtDataBytes)
	mockStub.MockTransactionEnd(txID)

	mockStub.MockTransactionStart("2")
	var got privateData
	err := GetImplicitPrivateData(mockTransactionContext, collectionMSP, txID, &got)
	mockStub.MockTransactionEnd("2")

	assertError(t, err, nil)
	if !reflect.DeepEqual(got, pvtData) {
		t.Errorf("got %v want %v", got, pvtData)
	}
}

func TestGetPrivateData(t *testing.T) {
	t.Run("Get existing private data", func(t *testing.T) {
		type privateData struct {
			Field1 string `json:"field1"`
			Field2 int    `json:"field2"`
			Field3 bool   `json:"field3"`
		}
		pvtData := privateData{"value1", 2, true}
		pvtDataBytes, err := json.Marshal(pvtData)
		if err != nil {
			t.Errorf("got unexpected error %q on Marshal of pvtData", err)
		}
		collection := "mockCollectionName"
		key := "mockPrivateDataKey"
		putTxID := "1"
		getTxID := "2"

		mockStub := mocking.NewMockChaincodeStub("TestGetPrivateData", nil, nil)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)
		mockStub.MockTransactionStart(putTxID)
		_ = mockStub.PutPrivateData(collection, key, pvtDataBytes)
		mockStub.MockTransactionEnd(putTxID)

		mockStub.MockTransactionStart(getTxID)
		got, err := GetPrivateData(mockTransactionContext, collection, key)
		mockStub.MockTransactionEnd(getTxID)

		assertError(t, err, nil)
		if !reflect.DeepEqual(got, pvtDataBytes) {
			t.Errorf("got %v want %v", got, pvtDataBytes)
		}
	})

	t.Run("private data not existing in collection or key", func(t *testing.T) {
		collection := "mockCollectionName"
		key := "mockPrivateDataKey"
		getTxID := "1"

		mockStub := mocking.NewMockChaincodeStub("TestGetPrivateData", nil, nil)
		mockTransactionContext := mocking.NewMockTransactionContext(mockStub, nil)

		mockStub.MockTransactionStart(getTxID)
		_, err := GetPrivateData(mockTransactionContext, collection, key)
		mockStub.MockTransactionEnd(getTxID)

		assertError(t, err, ErrPrivateDataNotFound)
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
