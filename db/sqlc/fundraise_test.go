package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/kinmaBackend/util"
	"github.com/stretchr/testify/require"
)

func createRandomFundraise(t *testing.T) Fundraise{
	accountArg := CreateAccountParams{
		Owner: 		util.RandomOwner(),
		Currency: util.RandomCurrency(),
	}

	account, _ := testQueries.CreateAccount(context.Background(),accountArg)

	productArg := CreateProductParams{
		AccountID	 : account.ID,
		Title      : util.RandomString(5),
		Content    : util.RandomString(20) + "_Content",
		ProductTag : util.RandomTag(),
	}

	product, _ := testQueries.CreateProduct(context.Background(), productArg)

	fundraiseArg := CreateFundraiseParams{
		ProductID				: product.ID,
		TargetAmount		: 10000,
		ProgressAmount	: 5000,
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
	createRandomFundraise(t)
}

func TestGetFundraise(t *testing.T){
	fundraise := createRandomFundraise(t)
	fundraise2, err := testQueries.GetFundraise(context.Background(), fundraise.ProductID)
	
	require.NoError(t, err)
	require.NotEmpty(t, fundraise2)
	require.Equal(t, fundraise.ProductID, fundraise2.ProductID)
	require.Equal(t, fundraise.TargetAmount, fundraise2.TargetAmount)
	require.Equal(t, fundraise.ProgressAmount, fundraise2.ProgressAmount)
	require.WithinDuration(t, fundraise.StartDate, fundraise2.StartDate, time.Second)
}

//if progress amount exceed target amount, Success should return true
func TestExitFundraise(t *testing.T){
	fundraise := createRandomFundraise(t)
	endTime := sql.NullTime {
		Time: time.Now().UTC(),
		Valid: true,
	}
	//fundraise Success
	if (fundraise.TargetAmount < fundraise.ProgressAmount){
		fundraise.Success = true
	}
	arg := ExitFundraiseParams{
		ProductID 	: fundraise.ProductID,
		Success			: fundraise.Success,
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