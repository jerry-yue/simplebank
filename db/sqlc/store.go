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

func NewStore(db *sql.DB) *Store {
	fmt.Println(">> NewStore()")
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	fmt.Println(">> execTx()")
	defer func() {
		err := recover() //内置函数，可以捕捉到函数异常
		if err != nil {
			//这里是打印错误，还可以进行报警处理，例如微信，邮箱通知
			fmt.Println("err错误信息：", err)
		}
	}()
	tx, err := store.db.BeginTx(ctx, nil)

	fmt.Println(tx, err)
	if err != nil {
		fmt.Println("<< execTx() Error 01")
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			fmt.Println("<< execTx() Error 02")
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		fmt.Println("<< execTx() Error 03")
		return err
	}
	fmt.Println("<< execTx()")
	return tx.Commit()
}

type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult
	fmt.Println(">>TransferTx()")
	err := store.execTx(ctx, func(q *Queries) error {
		fmt.Println("121212")
		var err error
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			fmt.Println("<< TransferTx() with error!")
			return err
		}

		result.FromEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToEntry, err = q.CreateEntry(ctx, CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		return nil
	})
	fmt.Println("<< TransferTx()")
	return result, err
}