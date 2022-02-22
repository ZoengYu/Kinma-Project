package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

// Payload contains the payload data of token
type Payload struct {
	ID 				uuid.UUID `json:"id"`
	Username 	string 		`json:"username"`
	IssuedAt 	time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		ID				: tokenID,
		Username	: username,
		IssuedAt	: time.Now(),
		ExpiredAt	: time.Now().Add(duration),
	}
	return payload, err 
}

// Valid checks if the token payload is valid or not
// Valid() func is required for jwt.NewWithClaims()
func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt){
		return ErrExpiredToken
	}
	return nil
}