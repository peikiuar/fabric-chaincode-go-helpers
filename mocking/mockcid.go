package mocking

import "crypto/x509"

// MockClientIdentity Implements the ClientIdentity interface for unit testing chaincode.
type MockClientIdentity struct{}

func (m *MockClientIdentity) GetID() (string, error) {
	return "", nil
}

func (m *MockClientIdentity) GetMSPID() (string, error) {
	return "", nil
}

func (m *MockClientIdentity) GetAttributeValue(string) (string, bool, error) {
	return "", false, nil
}

func (m *MockClientIdentity) AssertAttributeValue(string, string) error {
	return nil
}

func (m *MockClientIdentity) GetX509Certificate() (*x509.Certificate, error) {
	return nil, nil
}
