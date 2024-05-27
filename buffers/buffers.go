package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	// Let's set up a generic all-purpose buffer. We can use the
	// bytes.Buffer object for that.

	// I will use the new() builtin function here to allocate a
	// zero bytes object. This is similar to the make() function
	// but is used for value types.
	buf := new(bytes.Buffer)

	// now I'll write data to the buffer by sending a non-initialized
	// byte slice with an initialized value of "Hello"
	buf.Write([]byte("Hello"))

	// Or we can just write a string to the buffer too. Use the
	// WriteString() method for that.
	buf.WriteString(" World!")

	// If we need the length, we have to call the Len() method for it.
	// The len() builtin won't work.
	fmt.Println("Length of buffer:", buf.Len())

	// and we can just dump the contents to the screen as a string:
	fmt.Println(buf.String())

	// Now...let's dive into why this is useful. bytes.Buffer implements
	// the Reader and Writer functions, which most connection methods
	// do as well. In fact most methods that take in and put out data
	// could potentially implement the io.Reader and io.Writer functions.

	// init a new buffer for this purpose
	newbuf := new(bytes.Buffer)
	if err := WriteTo(newbuf, []byte("Hello World again!")); err != nil {
		fmt.Printf("Error: %+v\n", err)
	}
	fmt.Println(newbuf.String())
}

// This function could be different in that io.Writer could be a connection
// handle. But if that connection handle implements io.Writer as well, then
// this will work just the same!
func WriteTo(w io.Writer, b []byte) error {
	_, err := w.Write(b)
	return err
}
