package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const minSercretKeySize = 32

//JWTMaker is a JSON Web Token maker
type JWTMaker struct {
	secretKey string
}
// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(secretKey string) (Maker, error) {
	if len(secretKey) < minSercretKeySize {
		return nil, fmt.Errorf("invalid key size: must be at least %d characters", minSercretKeySize)
	}

	return &JWTMaker{secretKey}, nil
}

func (maker *JWTMaker) CreateToken(username string, duration time.Duration) (string, error){
	payload, err := NewPayload(username, duration)
	if err != nil {
		return "", err
	}
	//jwt encode the payload with SigningMethod
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyToken(token string) (*Payload, error){
	keyFunc := func(token *jwt.Token) (interface{}, error){
		//token.Method is a interface, convert it to jwt.SigningMethodHS256 as we use in CreateToken
		_, ok :=token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, ErrInvalidToken
		}
		return []byte(maker.secretKey), nil
	}
	//decode base on base64.URLEncoding.DecodeString method
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}