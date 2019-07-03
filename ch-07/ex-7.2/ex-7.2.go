/* Implement a function CountingWriter with io.Writer as
argument and returns a new Writer */

package main

import (
	"fmt"
	"io"
	"os"
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

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	nw := NewWriter{w, 0}
	return &nw, &nw.bytesWritten
}

func main() {
	writer, count := CountingWriter(os.Stdout)
	fmt.Fprint(writer, "Psychotic\n")
	fmt.Println(*count)
}
