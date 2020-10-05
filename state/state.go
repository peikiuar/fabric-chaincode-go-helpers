package state

import (
	"encoding/json"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// PutState takes care of marshalling the state value before storing it in the World State with the specified key
func PutState(ctx contractapi.TransactionContextInterface, key string, v interface{}) (err error) {
	value, err := json.Marshal(&v)
	if err != nil {
		return
	}

	err = ctx.GetStub().PutState(key, value)
	return
}
