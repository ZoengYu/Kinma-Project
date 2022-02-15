package db

import (
	"database/sql"
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

// execTx executes a function within a database transaction
// nil here represent type &sql.TxOptions{}, nil we set isolation level as default
// which is read-committed
// func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
// 	tx, err := store.db.BeginTx(ctx, nil)
// 	if err != nil {
// 		return err
// 	}

// 	query := New(tx)
// 	err = fn(query)
// 	if err != nil {
// 		if rollbackErr := tx.Rollback(); rollbackErr != nil {
// 			return fmt.Errorf("tx err: %v, rb err: %v", err, rollbackErr)
// 		}
// 		return err
// 	}
// 	return tx.Commit()
// }