package state

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// Technical state errors
// TODO: make errors constant
var (
	ErrStateNotFound = errors.New("State was not found")
)

// PutState takes care of marshalling the state value before storing it in the World State with the specified key
func PutState(ctx contractapi.TransactionContextInterface, key string, v interface{}) (err error) {
	value, err := json.Marshal(v)
	if err != nil {
		return
	}

	err = ctx.GetStub().PutState(key, value)
	return
}

// GetState is a function to get data from the World State that will return an error if the state does not exist and that will unmarshal to a specified empty interface if found
func GetState(ctx contractapi.TransactionContextInterface, key string, v interface{}) (err error) {
	bytes, err := ctx.GetStub().GetState(key)
	if err != nil {
		return
	}
	if bytes == nil {
		err = ErrStateNotFound
		return
	}

	err = json.Unmarshal(bytes, v)
	return
}
