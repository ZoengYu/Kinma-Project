package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransfer(t *testing.T) {
	store := NewStore(testDB)

	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)
	product1 := createRandomProduct(t, account2)
	fundraise1 := createRandomFundraise(t, product1)
	//concurrent transfer transactions
	n := 5
	amount := int64(10)
	//get the ProgressAmount before account sponsor
	beforeTxAmount := fundraise1.ProgressAmount

	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			//Transfer will include `createTransfer` Record and `addMoney` to product's fundraise
			result, err := store.TransferTx(context.Background(), TransferParams{
				FromAccountID: account1.ID,
				ToProductID:   fundraise1.ProductID,
				Amount:        amount,
			})

			errs <- err
			results <- result
		}()
	}

	for i := 0; i < n; i++ {
		err := <-errs
		require.NoError(t, err)

		result := <-results
		require.NotEmpty(t, result)

		//check transfer detail is correct
		transfer := result.Transfer

		require.NotEmpty(t, transfer)
		require.Equal(t, account1.ID, transfer.FromAccountID)
		require.Equal(t, fundraise1.ProductID, transfer.ToProductID)
		require.Equal(t, amount, transfer.Amount)
		require.NotZero(t, transfer.ID)
		require.NotZero(t, transfer.CreatedAt)

		transferSuccess, err := store.GetTransfer(context.Background(), GetTransferParams{
			FromAccountID: account1.ID,
			ToProductID:   fundraise1.ProductID,
		})
		require.NoError(t, err)

		//check account ID & fundraise ID is correct
		require.Equal(t, account1.ID, result.FromAccountID.ID)
		require.Equal(t, fundraise1.ID, result.Fundraise.ID)

		// check amount progress is equivalent as expected
		require.Equal(t, beforeTxAmount+int64(i+1)*amount, result.Fundraise.ProgressAmount)
		fmt.Print(result.Fundraise.ProgressAmount)

		// check get fundraise is working well during the transaction
		_, err = store.GetProductFundraise(context.Background(), transfer.ToProductID)
		require.NoError(t, err)

		//check transfer is success
		require.True(t, transferSuccess.Success)
	}

	//check final update of fundraise amount
	updatedFundraise, err := store.GetProductFundraise(context.Background(), fundraise1.ProductID)
	require.NoError(t, err)
	require.Equal(t, updatedFundraise.ProgressAmount, fundraise1.ProgressAmount+amount*int64(n))
}
