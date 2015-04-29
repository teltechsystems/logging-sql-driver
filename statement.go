package loggingsql

import (
	"database/sql/driver"
)

type loggingStmt struct {
	wrappedStmt driver.Stmt
}

func (s *loggingStmt) Close() error {
	logger.Println("> closing statement")

	if err := s.wrappedStmt.Close(); err != nil {
		logger.Printf("> failed to close statement: %s", err)

		return err
	}

	logger.Println("> closed statement successfully")

	return nil
}

func (s *loggingStmt) Exec(args []driver.Value) (driver.Result, error) {
	logger.Printf("> running exec with args: %#v", args)

	result, err := s.wrappedStmt.Exec(args)
	if err != nil {
		logger.Printf("> failed to exec: %s", err)

		return nil, err
	}

	logger.Printf("> exec'ed successfully: %#v", result)

	return result, nil
}

func (s *loggingStmt) NumInput() int {
	logger.Println("> getting number of inputs")

	numInput := s.wrappedStmt.NumInput()

	logger.Printf("> total number of inputs: %d", numInput)

	return numInput
}

func (s *loggingStmt) Query(args []driver.Value) (driver.Rows, error) {
	logger.Printf("> running query with args: %#v", args)

	rows, err := s.wrappedStmt.Query(args)
	if err != nil {
		logger.Printf("> failed to query: %s", err)

		return nil, err
	}

	logger.Printf("> queried successfully: %#v", rows)

	return rows, nil
}
