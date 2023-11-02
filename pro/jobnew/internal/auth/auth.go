package auth

import (
	"crypto/rsa"
	"errors"
)

type ctxKey int

const Key ctxKey = 1

type Auth struct {
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewAuth(privateKey *rsa.PrivateKey, publicKey *rsa.PublicKey) (*Auth, error) {
	if privateKey == nil || publicKey == nil {
		return nil, errors.New("Private key/Public key cannot be nill")
	}
	return &Auth{
		privateKey: privateKey,
		publicKey:  publicKey,
	}, nil
}
