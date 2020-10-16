package pvtdata

import (
	"encoding/json"
	"errors"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

const (
	implicitCollectionPrefix = "_implicit_org_"
)

// Technical pvtdata errors
// TODO: make errors constant
var (
	ErrWrongTransientFieldName  = errors.New("Field is not present in transient data map")
	ErrEmptyTransientFieldValue = errors.New("Transient field has empty value")
)

// GetTransientDataValue is a function that obtains the value of a specific field of the private data
func GetTransientDataValue(ctx contractapi.TransactionContextInterface, fieldName string, v interface{}) (err error) {
	TransientMap, err := ctx.GetStub().GetTransient()
	if err != nil {
		return
	}

	value, ok := TransientMap[fieldName]
	if !ok {
		err = ErrWrongTransientFieldName
		return
	}

	if len(value) == 0 {
		err = ErrEmptyTransientFieldValue
	}

	err = json.Unmarshal(value, v)
	return
}

// func PutImplicitPrivateData(ctx contractapi.TransactionContextInterface) {}

// func GetImplicitPrivateData(ctx contractapi.TransactionContextInterface) {}
