package couchdb

import (
	"bytes"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// QueryCouchDB lets you execute rich CouchDB queries
func QueryCouchDB(ctx contractapi.TransactionContextInterface, query string) (queryResult *bytes.Buffer, err error) {
	resultsIterator, err := ctx.GetStub().GetQueryResult(query)
	if err != nil {
		return
	}
	defer resultsIterator.Close()

	queryResult, err = constructQueryResponseFromIterator(resultsIterator)
	return
}

// QueryPrivatedata lets you execute rich private data queries
func QueryPrivatedata(ctx contractapi.TransactionContextInterface, collection string, query string) (queryResult *bytes.Buffer, err error) {
	resultsIterator, err := ctx.GetStub().GetPrivateDataQueryResult(collection, query)
	if err != nil {
		return
	}
	defer resultsIterator.Close()

	queryResult, err = constructQueryResponseFromIterator(resultsIterator)
	return
}

func constructQueryResponseFromIterator(it shim.StateQueryIteratorInterface) (*bytes.Buffer, error) {
	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for it.HasNext() {
		queryResponse, err := it.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten {
			buffer.WriteString(",")
		}
		bArrayMemberAlreadyWritten = true

		buffer.WriteString(string(queryResponse.Value))
	}
	buffer.WriteString("]")

	return &buffer, nil
}
