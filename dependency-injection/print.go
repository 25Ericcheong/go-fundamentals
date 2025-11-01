package dependency_injection

import (
	"bytes"
	"fmt"
)

func Greet(writer *bytes.Buffer, name string) {
	_, _ = fmt.Fprintf(writer, "Hello, %s", name)
}

// Wanting to mock fmt.Printf
// Implementation looks like
// func Printf(format string, a ...interface{}) (n int, err error) {
//	return Fprintf(os.Stdout, format, a...)
//}

// Look at implementation of Fprintf so we can mock that instead
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
//	p := newPrinter()
//	p.doPrintf(format, a)
//	n, err = w.Write(p.buf)
//	p.free()
//	return
//}

// Can see that io.Writer has the following interface
// type Writer interface {
//	Write(p []byte) (n int, err error)
//}
