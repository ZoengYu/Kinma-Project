package db

import (
	"context"
	"testing"

	"github.com/kinmaBackend/util"
	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, FromAccount Account, ToFundraise Fundraise) Transfer {
	arg := CreateTransferParams{
		FromAccountID	: FromAccount.ID,
		ToProductID		: ToFundraise.ProductID,
		Amount				: util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, arg.FromAccountID, transfer.FromAccountID)
	require.Equal(t, arg.ToProductID, transfer.ToProductID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)
	require.Equal(t, transfer.Success, false)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	product1 := createRandomProduct(t, account2)

	targetFundraise := createRandomFundraise(t, product1)

	createRandomTransfer(t, account1, targetFundraise)
}

func TestGetTransfer(t *testing.T) {
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	product1 := createRandomProduct(t, account2)

	fundraise1 := createRandomFundraise(t, product1)
	createdTransfer := createRandomTransfer(t, account1, fundraise1)

	getTransfer, err := testQueries.GetTransfer(context.Background(), GetTransferParams{
		FromAccountID	: account1.ID,
		ToProductID		: fundraise1.ProductID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, getTransfer)
	require.Equal(t, getTransfer.FromAccountID, createdTransfer.FromAccountID)
	require.Equal(t, getTransfer.ToProductID, createdTransfer.ToProductID)

	require.Equal(t, getTransfer.ID, createdTransfer.ID)
	require.Equal(t, getTransfer.Amount, createdTransfer.Amount)
	require.NotZero(t, getTransfer.CreatedAt)
}

func TestListTransfers(t *testing.T){
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	product1 := createRandomProduct(t, account2)
	fundraise1 := createRandomFundraise(t, product1)

	for i := 0; i < 10; i++ {
		TransferArg := CreateTransferParams{
			FromAccountID	: product1.AccountID,
			ToProductID		: fundraise1.ProductID,
			Amount			 	: util.RandomMoney(),
		}
		testQueries.CreateTransfer(context.Background(), TransferArg)
	}

	listArg := ListTransfersParams{
		FromAccountID	: account1.ID,
		ToProductID		: fundraise1.ProductID,
		Limit					:	5,
		Offset				: 5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), listArg)
	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers{
		require.NotEmpty(t, transfer)
	}
}