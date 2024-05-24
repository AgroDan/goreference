# Working with Files and Strings

In this section, I have some functions for basic working with opening, creating, and writing to files. Additionally, some string functions as well. If you need to split on strings, it's really not that difficult as long as you import the `strings` library.

The main function here shows several functions:

```go
func main() {
	ReadEntireFile()
	ReadInChunks(1024)
	WriteToFile("Hello World!\n")

	EnterSomething()
	words := SplitStrings("The quick brown fox jumped over the lazy dogs")

	fmt.Printf("Print the first 3 words: ")
	for i := 0; i < 3; i++ {
		fmt.Printf("%s ", words[i])
	}
	fmt.Println()
```

## ReadEntireFile()

This function loads the entire file you choose directly into memory. This is a fairly useful method if you plan on working with small, manageable amounts of data to be ingested.

## ReadInChunks()

Similar to the above, this is a bit more "memory-friendly." This will read a file in chunks of N bytes (N is provided by argument, in the above case 1024). This will let you do some in-line editing on extremely large files without blowing up your RAM.

## WriteToFile()

This function demonstrates -- you guessed it -- how to write to a file.

## EnterSomething()

This will show you how to get user input from the command line. This is a little wonky and I feel like it can be done better, but this is the Go way I guess.

## SplitStrings()

This is how you split strings on a delimiter you want. This uses the `strings` library.