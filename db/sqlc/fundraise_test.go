package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/kinmaBackend/util"
	"github.com/stretchr/testify/require"
)

func createRandomFundraise(t *testing.T, product Product) Fundraise{
	amount := util.RandomMoney()
	fundraiseArg := CreateFundraiseParams{
		ProductID				: product.ID,
		TargetAmount		: amount,
		ProgressAmount	: amount/2,
	}
	
	fundraise, err := testQueries.CreateFundraise(context.Background(), fundraiseArg)
	require.NoError(t, err)
	require.NotEmpty(t, fundraise)
	require.Equal(t, fundraise.TargetAmount, fundraiseArg.TargetAmount)
	require.Equal(t, fundraise.ProgressAmount, fundraiseArg.ProgressAmount)
	require.NotZero(t, fundraise.StartDate)
	require.Equal(t, fundraise.Success, false)

	require.Equal(t, fundraise.ProductID, product.ID)
	require.NotZero(t, fundraise.ID)

	return fundraise
}

func TestCreateFundraise(t *testing.T){
	account1 := createRandomAccount(t)
	product1 := createRandomProduct(t, account1)
	
	createRandomFundraise(t, product1)
}

func TestGetFundraise(t *testing.T){
	account1 := createRandomAccount(t)
	product1 := createRandomProduct(t, account1)
	fundraise1 := createRandomFundraise(t, product1)
	fundraise2, err := testQueries.GetFundraise(context.Background(), fundraise1.ProductID)
	
	require.NoError(t, err)
	require.NotEmpty(t, fundraise2)
	require.Equal(t, fundraise1.ProductID, fundraise2.ProductID)
	require.Equal(t, fundraise1.TargetAmount, fundraise2.TargetAmount)
	require.Equal(t, fundraise1.ProgressAmount, fundraise2.ProgressAmount)
	require.WithinDuration(t, fundraise1.StartDate, fundraise2.StartDate, time.Second)
}

//if progress amount exceed target amount, Success should return true
func TestExitFundraise(t *testing.T){
	account1 := createRandomAccount(t)
	product1 := createRandomProduct(t, account1)
	fundraise1 := createRandomFundraise(t, product1)

	endTime := sql.NullTime {
		Time: time.Now().UTC(),
		Valid: true,
	}
	//fundraise Success
	if (fundraise1.TargetAmount < fundraise1.ProgressAmount){
		fundraise1.Success = true
	}
	arg := ExitFundraiseParams{
		ProductID 	: fundraise1.ProductID,
		Success			: fundraise1.Success,
		EndDate			: endTime,
	}

	endFundraise, err := testQueries.ExitFundraise(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, endFundraise)
	
	fundraise2, err := testQueries.GetFundraise(context.Background(), endFundraise.ProductID)
	require.NoError(t, err)
	require.NotEmpty(t, fundraise2)
	require.Equal(t, endFundraise.EndDate, fundraise2.EndDate)
	require.Equal(t, endFundraise.ID, fundraise2.ID)
	require.Equal(t, endFundraise.Success, fundraise2.Success)
}