package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T) {
	integer := RandomInt(0, 100)
	integer2 := RandomInt(0, 100)
	require.NotEqual(t, integer, integer2)
}

func TestRandomOwner(t *testing.T) {
	owner1 := RandomOwner()
	owner2 := RandomOwner()
	require.NotEqual(t, owner1, owner2)
}

func TestRandomTag(t *testing.T) {
	tag1 := RandomTag()
	tag2 := RandomTag()
	tag3 := RandomTag()
	require.NotEqual(t, tag1, tag2, tag3)
}

func TestRandomMoney(t *testing.T) {
	money1 := RandomMoney()
	money2 := RandomMoney()
	require.NotEqual(t, money1, money2)
}

func TestRandomCurrency(t *testing.T) {
	currencies := []string{"TWD", "USD", "CNY"}
	currency1 := RandomCurrency()
	require.Contains(t, currencies, currency1)
}

func TestRandomEmail(t *testing.T) {
	emailType := "@email.com"
	email1 := RandomEmail()
	email2 := RandomEmail()
	require.Contains(t, email1, emailType)
	require.Contains(t, email2, emailType)
}
