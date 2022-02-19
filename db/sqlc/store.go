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

//create a new Store
func NewStore(db *sql.DB) *Store {
	return &Store{
		db			:	db,
		Queries : New(db),
	}
}

//execTx executes a function within a database transaction
//nil here represent type &sql.TxOptions{}, nil we set isolation level as default
//which is read-committed
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

type TransferParams struct {
	FromeAccountID int64 `json:"from_account_id"`
	ToFundraiseID  int64 `json:"to_fundraise_id"`
	Amount				 int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer 				Transfer 	`json:"transfer"`
	FromeAccount		Account  	`json:"from_account"`
	Fundraise  			Fundraise `json:"to_fundraise"`
}

//export TransferTx to implement transfer from account to the fundraise.
func (store *Store) TransferTx(ctx context.Context, arg TransferParams) (TransferTxResult, error){
	var result TransferTxResult
	err := store.execTx(ctx, func(q *Queries) error{
		var err error
		//create transfer record
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID	: arg.FromeAccountID,
			ToFundraiseID	: arg.ToFundraiseID,
			Amount				: arg.Amount,
		})
		if err != nil {
			return err
		}
		//add Amount to fundraise project
		result.Fundraise, err = q.AddFundraiseProgressAmount(ctx, AddFundraiseProgressAmountParams{
			Amount	: arg.Amount,
			ID			: arg.ToFundraiseID,
		})
		if err != nil{
			return err
		}

		result.FromeAccount, err = q.GetAccount(ctx, arg.FromeAccountID)
		if err != nil{
			return err
		}

		return nil
	})
	return result, err
}