package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomInt(t *testing.T){
	integer := RandomInt(0,100)
	integer2 := RandomInt(0,100)
	require.NotEqual(t, integer, integer2)
}

func TestRandomOwner(t *testing.T){
	owner1 := RandomOwner()
	owner2 := RandomOwner()
	require.NotEqual(t, owner1, owner2)
}

//It's possible the currency1-3 is equal since the current option only have 3
func TestRandomCurrency(t *testing.T){
	currency1 := RandomCurrency()
	currency2 := RandomCurrency()
	currency3 := RandomCurrency()
	require.NotEqual(t, currency1, currency2, currency3)
}

func TestRandomTag(t *testing.T){
	tag1 := RandomTag()
	tag2 := RandomTag()
	tag3 := RandomTag()
	require.NotEqual(t, tag1, tag2, tag3)
}

func TestRandomMoney(t *testing.T){
	money1 := RandomInt(0,1000)
	money2 := RandomInt(0,1000)
	require.NotEqual(t, money1, money2)
}