package main

import (
	"fmt"
	"io"
	"os"
)

func ReadEntireFile() {
	// This function will read the entire file and print it to the screen.
	// First, let's open the file. We'll need the "os" package to do this.
	// I'll store the location of the file in the thisFile variable
	thisFile := "./myfile.txt"

	// Now I'll open the file path for reading:
	file, err := os.Open(thisFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// Use the `defer` directive to ensure the file is closed when this
	// application wraps up, successful or not.
	defer file.Close()

	// Now we can do a couple of things. If the file is relatively small,
	// we can just slurp it all up into memory and call it a day. Let's
	// do that here.

	// Read file content
	content, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Now we can do whatever to the file. Like print it.
	fmt.Println(string(content))
}

func ReadInChunks(bytesReadSize int) {
	// This does the same thing as the above, but will instead read
	// a relatively large file in chunks of `bytes` length at a time
	// so it doesn't overload memory
	thisFile := "./largefile.txt"

	// Open for reading
	file, err := os.Open(thisFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	// defer file closure
	defer file.Close()

	// Now let's create a buffer. I think this can be done with a
	// go library but we can just use memory to store this data.
	// Let's create the buffer of size "bytesReadSize"
	buf := make([]byte, bytesReadSize)

	// Now let's read the file in chunks of buffSize
	for {
		// Loop indefinitely
		bytesRead, err := file.Read(buf)
		if err != nil {
			fmt.Println("File read returned an error:", err)
		}

		// If bytesRead ever returns a zero, we hit EOF
		if bytesRead == 0 {
			break
		}

		// Otherwise, work with the contents we read. Remember
		// that we ONLY care about the bytes that were read. If
		// we don't specify the amount to read to, it will read
		// the previous iteration's data!
		fmt.Print(string(buf[:bytesRead]))

		// also note that the bytes it returns will be in a byte array.
		// We can turn that into a string.
	}

}

func main() {
	ReadEntireFile()
	ReadInChunks()

}
