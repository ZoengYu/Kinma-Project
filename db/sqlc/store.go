package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	*Queries
	db *sql.DB
}

// create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
// nil here represent type &sql.TxOptions{}, nil we set isolation level as default
// which is read-committed
func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := New(tx)
	err = fn(query)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rollbackErr)
		}
		return err
	}
	return tx.Commit()
}

// Currency is a optional user input which will be verified during the transfer
type TransferParams struct {
	FromAccountID int64  `json:"from_account_id"`
	ToProductID   int64  `json:"to_product_id"`
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
}

type TransferTxResult struct {
	Transfer      Transfer  `json:"transfer"`
	FromAccountID Account   `json:"from_account"`
	Fundraise     Fundraise `json:"to_fundraise"`
}

//export TransferTx to implement transfer from account to the fundraise.
func (store *Store) TransferTx(ctx context.Context, arg TransferParams) (TransferTxResult, error) {
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error {
		var err error
		//check if account exist
		result.FromAccountID, err = q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		//check if fundraise exist
		getFundraise, err := q.GetProductFundraise(ctx, arg.ToProductID)
		if err != nil {
			return err
		}
		//create transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToProductID:   arg.ToProductID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		//add Amount to fundraise project
		result.Fundraise, err = q.AddFundraiseProgressAmount(ctx, AddFundraiseProgressAmountParams{
			Amount: arg.Amount,
			ID:     arg.ToProductID,
		})
		if err != nil {
			return err
		}
		//update transfer if is successed
		if result.Fundraise.ProgressAmount == getFundraise.ProgressAmount+result.Transfer.Amount {
			result.Transfer, err = q.UpdateTransferStatus(ctx, UpdateTransferStatusParams{
				result.Transfer.ID, true})
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}
