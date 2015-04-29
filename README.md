Logging SQL Driver
==================

This driver is intended to wrap around existing sql drivers to provide logging functionality. This is great for timing
queries, inspecting arguments, and quickly peeking at the actual results without needing to log them on your own

Usage
=====
```
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/teltechsystems/logging-sql-driver"
)

func main() {
	conn, err := sql.Open("logging:mysql", "DSN")
	fmt.Printf("conn: %s", conn)
	fmt.Printf("err: %s", err)
	defer conn.Close()

	// ....
}
```