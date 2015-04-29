Logging SQL Driver
==================

This driver is intended to wrap around existing sql drivers to provide logging functionality. This is great for timing
queries, inspecting arguments, and quickly peeking at the actual results without needing to log them on your own

Usage
=====
`sql.Open("logging:mysql", "DSN")`