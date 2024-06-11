package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func ReadEntireFile() {
	// This function will read the entire file and print it to the screen.
	// First, let's open the file. We'll need the "os" package to do this.
	// I'll store the location of the file in the thisFile variable
	thisFile := "./filesandstrings/myfile.txt"

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
	thisFile := "./filesandstrings/largefile.txt"

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

func ReadLineByLine() {
	// This function will read a file line by line. It will also make
	// use of the io.Reader interface that os.Open() will return.
	fileReader, err := os.Open("./filesandstrings/poetry.txt")
	if err != nil {
		panic(err)
	}
	defer fileReader.Close()

	// Now we'll set up a bufio scanner which has a lot more efficient
	// means of parsing a file. It accepts io.Reader as input!
	fr := bufio.NewScanner(fileReader)

	// Now we'll loop over each line, which fr.Scan() is built to handle!
	// It will loop until EOF. This should be fine for most scenarios, because
	// you can then split the resulting string using the strings library if
	// necessary.
	var counter int = 0
	for fr.Scan() {
		fmt.Println("Line", counter, "-", fr.Text())
		counter++
	}
}

func WriteToFile(someData string) {
	// Writing to a file, by default, accepts a byte array.
	// If you want to write a string however, you can use the
	// WriteString() function.

	// First, let's Create a file
	createThisFile := "./filesandstrings/NewlyCreatedFile.txt"

	file, err := os.Create(createThisFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
	}
	defer file.Close()

	// You can just use `bytesWritten, err := f.WriteString(someData)`
	// here, but it's better to use a buffered writer provided by bufio
	w := bufio.NewWriter(file)

	bytesWritten, err := w.WriteString(someData)
	if err != nil {
		fmt.Println("Could not write to file:", err)
	}

	fmt.Printf("Wrote %d bytes.\n", bytesWritten)

	// now flush the buffer to the file
	w.Flush()

	// note that we can execute a similar file write buffer example
	// as reading if you need to write to larger files
}

func EnterSomething() {
	// This will read from user input and work with the data. We'll just
	// convert the string to uppercase and lowercase. For this, I'll use
	// the fmt library to scan a line, meaning it will listen on STDIN
	// until a newline is received. In this case, this happens when the
	// user presses the "Enter" key.

	fmt.Printf("Enter a string for me to play with: ")

	// The below part will work if you want to delimit on a specific
	// rune, in this case \n -- but the scanner might be better
	// suited for reading line-by-line

	// reader := bufio.NewReader(os.Stdin)
	// fmt.Printf("Enter text: ")
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Error reading input:", err)
	// }
	// fmt.Println(input[:len(input)-1])

	scanner := bufio.NewScanner(os.Stdin)
	var inputText string
	if scanner.Scan() {
		inputText = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input: ", err)
	}
	// Now let's manipulate this string!

	fmt.Printf("Uppercase: %s!!!\n", strings.ToUpper(inputText))
	fmt.Printf("Lowercase: ...%s...\n", strings.ToLower(inputText))

}

func SplitStrings(myString string) []string {
	// This function will split a string into words
	// you can also split them by lines
	return strings.Split(myString, " ")
}

func main() {
	ReadEntireFile()
	ReadInChunks(1024)
	ReadLineByLine()
	WriteToFile("Hello World!\n")

	EnterSomething()
	words := SplitStrings("The quick brown fox jumped over the lazy dogs")

	fmt.Printf("Print the first 3 words: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%s ", words[i])
	}
	fmt.Println()
}
