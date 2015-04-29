package loggingsql

import (
	"database/sql/driver"
)

type loggingTx struct {
	wrappedTx driver.Tx
}

func (tx *loggingTx) Commit() error {
	logger.Println("> committing transaction")

	if err := tx.wrappedTx.Commit(); err != nil {
		logger.Printf("> failed to commit transaction: %s", err)

		return err
	}

	logger.Println("> committed transaction")

	return nil
}

func (tx *loggingTx) Rollback() error {
	logger.Println("> rolling back transaction")

	if err := tx.wrappedTx.Rollback(); err != nil {
		logger.Printf("> failed to rollback transaction: %s", err)

		return err
	}

	logger.Println("> rollback transaction successful")

	return nil
}
