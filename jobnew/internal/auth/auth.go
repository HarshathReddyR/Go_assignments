package auth

import (
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
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
func (a *Auth) GenerateToken(claims jwt.RegisteredClaims) (string, error) {
	tkn := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	token, err := tkn.SignedString(a.privateKey)
	if err != nil {
		return "", fmt.Errorf("signing token %w", err)
	}
	return token, nil
}
func (a *Auth) ValidateToken(token string) (jwt.RegisteredClaims, error) {
	var c jwt.RegisteredClaims
	tkn, err := jwt.ParseWithClaims(token, &c, func(token *jwt.Token) (interface{}, error) {
		return a.publicKey, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, fmt.Errorf("parsing token %w", err)
	}
	// Check if the parsed token is valid.
	if !tkn.Valid {
		return jwt.RegisteredClaims{}, errors.New("invalid token")
	}
	return c, nil
}
