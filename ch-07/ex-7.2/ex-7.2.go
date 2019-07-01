/* Implement a function CountingWriter with io.Writer as
argument and returns a new Writer */

package main

import (
	"fmt"
	"io"
)

type NewWriter struct {
	w            io.Writer
	bytesWritten int64
}

func (n *NewWriter) Write(p []byte) (int, error) {
	count, err := n.w.Write(p)
	n.bytesWritten += int64(count)
	return count, err
}
