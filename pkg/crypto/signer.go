package crypto

import (
	"fmt"
)

// SignData signs the data using the private key directly obtained from keys.go.
// This function assumes that the private key is already loaded or generated by keys.go.
func SignData(keyFile string, data []byte) ([]byte, error) {
	privKey, _, _, err := GetOrCreatePrivateKey(keyFile)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve or create private key: %w", err)
	}
	if privKey == nil {
		return nil, fmt.Errorf("private key is nil after retrieval or creation")
	}
	return privKey.Sign(data)
}

// VerifySignature is updated to retrieve the private key using GetOrCreatePrivateKey
// and then derive the public key from it for verification.
func VerifySignature(keyFile string, data []byte, signature []byte) (bool, error) {
	// Retrieve or create the private key
	privKey, _, _, err := GetOrCreatePrivateKey(keyFile)
	if err != nil {
		return false, fmt.Errorf("failed to retrieve or create private key: %w", err)
	}
	if privKey == nil {
		return false, fmt.Errorf("private key is nil after retrieval or creation")
	}

	// Derive the public key from the private key
	pubKey := privKey.GetPublic()
	if pubKey == nil {
		return false, fmt.Errorf("public key is nil after derivation")
	}

	// Use the derived public key to verify the signature
	return pubKey.Verify(data, signature)
}
