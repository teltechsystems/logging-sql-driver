package loggingsql

import (
	"database/sql"
	"database/sql/driver"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var logger = log.New(os.Stdout, "logging-sql-driver: ", log.Ldate|log.Ltime)

type LoggingDriver struct {
	Driver string
}

func (d LoggingDriver) Open(dsn string) (driver.Conn, error) {
	mysqlDriver := &mysql.MySQLDriver{}
	conn, err := mysqlDriver.Open(dsn)
	if err != nil {
		logger.Printf("> failed to open connection: %s", err)

		return nil, err
	}

	return &loggingConn{wrappedConn: conn}, nil
}

func init() {
	sql.Register("logging:mysql", &LoggingDriver{"mysql"})
}
