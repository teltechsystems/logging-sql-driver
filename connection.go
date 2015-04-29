package loggingsql

import (
	"database/sql/driver"
)

type loggingConn struct {
	wrappedConn driver.Conn
}

func (c *loggingConn) Begin() (driver.Tx, error) {
	logger.Println("> beginning transaction")

	tx, err := c.wrappedConn.Begin()
	if err != nil {
		logger.Printf("> failed to begin transaction: %s", err)

		return nil, err
	}

	return &loggingTx{wrappedTx: tx}, nil
}

func (c *loggingConn) Close() error {
	logger.Println("> closing connection")

	if err := c.wrappedConn.Close(); err != nil {
		logger.Printf("> failed to close connection: %s", err)

		return err
	}

	logger.Println("> closed connection successfully")

	return nil
}

func (c *loggingConn) Prepare(query string) (driver.Stmt, error) {
	logger.Printf("preparing query: %s", query)

	stmt, err := c.wrappedConn.Prepare(query)
	if err != nil {
		logger.Printf("> failed to prepare query: %s", err)

		return nil, err
	}

	logger.Println("> query prepared successfully")

	return &loggingStmt{wrappedStmt: stmt}, nil
}
