package main

import (
	"fmt"
	"io"
)

type NewWriter struct {
	w            io.Writer
	bytesWritten *int64
}
