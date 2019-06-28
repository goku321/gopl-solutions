/* Implement a function CountingWriter with io.Writer as
argument and returns a new Writer */

package main

import (
	"fmt"
	"io"
)

type NewWriter struct {
	w            io.Writer
	bytesWritten *int64
}
