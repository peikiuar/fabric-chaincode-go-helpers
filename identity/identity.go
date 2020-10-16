package identity

import (
	"github.com/hyperledger/fabric-chaincode-go/pkg/cid"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func AssertClientMSPID(ctx contractapi.TransactionContextInterface, cid cid.ClientIdentity, mspID string) bool {
	return true
}

func AssertClientOU(ctx contractapi.TransactionContextInterface, cid cid.ClientIdentity, ou string) bool {
	return true
}
