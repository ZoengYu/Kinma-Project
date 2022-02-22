package token

import (
	"crypto/ed25519"
	"fmt"
	"time"

	"github.com/o1egl/paseto"
	"golang.org/x/crypto/chacha20poly1305"
)

//AsymmetricPasetoMaker is a asymmetric PASETO token maker which suite for public
type AsymmetricPasetoMaker struct {
	paseto *paseto.V2
	publickey ed25519.PublicKey
	privatekey ed25519.PrivateKey
}

func NewAsymmetricPasetoMaker() (Maker, error){
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil{
		return nil, fmt.Errorf("generate publicKey and privateKey failed")
	}

	maker := &AsymmetricPasetoMaker{
		paseto				: paseto.NewV2(),
		publickey			: publicKey,
		privatekey 		: privateKey,
	}

	return maker, nil
}

//CreateToken creates a new token for specific username and duration
func (maker *AsymmetricPasetoMaker) CreateToken(username string, duration time.Duration) (string, error){
		payload, err := NewPayload(username, duration)
		if err != nil {
			return "", err
		}

		return maker.paseto.Sign(maker.privatekey, payload, nil)
	}

// VerifyToken checks if the token is valid or not
func (maker *AsymmetricPasetoMaker) VerifyToken(token string) (*Payload, error){
		payload := &Payload{}

		err := maker.paseto.Verify(token, maker.publickey, payload, nil)
		if err != nil {
			return nil, ErrInvalidToken
		}

		err = payload.Valid()
		if err != nil{
			return nil, err
		}

		return payload, nil
	}

	//SymmetricPasetoMaker is a Symmetric PASETO token maker which suite for local
type SymmetricPasetoMaker struct {
	paseto *paseto.V2
	symmetricKey []byte
}

// NewPasetoMaker creates a new PasetoMaker
func NewSymmetricPasetoMaker(symmetricKey string) (Maker, error){
	if len(symmetricKey) != chacha20poly1305.KeySize {
		return nil, fmt.Errorf("invalid key size: must be exactly %d characters", chacha20poly1305.KeySize)
	}

	maker := &SymmetricPasetoMaker{
		paseto			: paseto.NewV2(),
		symmetricKey: []byte(symmetricKey),
	}

	return maker, nil
}

//CreateToken creates a new token for specific username and duration
func (maker *SymmetricPasetoMaker) CreateToken(username string, duration time.Duration) (string, error){
		payload, err := NewPayload(username, duration)
		if err != nil {
			return "", err
		}

		return maker.paseto.Encrypt(maker.symmetricKey, payload, nil)
	}

// VerifyToken checks if the token is valid or not
func (maker *SymmetricPasetoMaker) VerifyToken(token string) (*Payload, error){
		payload := &Payload{}

		err := maker.paseto.Decrypt(token, maker.symmetricKey, payload, nil)
		if err != nil {
			return nil, ErrInvalidToken
		}

		err = payload.Valid()
		if err != nil{
			return nil, err
		}

		return payload, nil
	}