package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// First, let's list the contents of a directory in Go. This can be
	// accomplished with the `os` library. Let's read the content of the
	// ./example/ directory and print it to the screen.

	thisDir := "./files/examples/"

	files, err := os.ReadDir(thisDir)
	if err != nil {
		panic("Could not read from directory")
	}

	// Now print them out
	for _, f := range files {
		fmt.Printf("FILE: %s\n", f)
	}

	// We can also get a SHA256 sum of each file that we loop through. Go
	// is nice enough to give us a slice of files within the whole directory,
	// so we can use the io.Copy() function to pass the contents of the file
	// into the sha256 object.

	// First, we have the contents of the directory in the "files" variable.
	for _, f := range files {

		// skip over any potential directories.
		if f.IsDir() {
			continue
		}

		// Construct the fill path of the file using the 'path/filepath" module
		abs_path_file := filepath.Join(thisDir, f.Name())

		// Now open the file
		fp, err := os.Open(abs_path_file)
		if err != nil {
			panic("Could not read file")
		}

		// Create the hash object from the crypto/sha256 module
		hash := sha256.New()

		// use io.Copy to copy the contents of the file (being a ReadObject) to the
		// sha256 object, being a WriteObject. The io.Copy() function will return
		// the amount of bytes read, 0 if EOF (not an error), and an error if
		// something fatal happens.
		if _, err := io.Copy(hash, fp); err != nil {
			fp.Close()
			fmt.Printf("Could not hash file: %s\n", fp.Name())
			// I won't panic here I guess
			continue
		}
		fp.Close()

		// Now let's convert the hash to a readable string, typically this means
		// convert the data to hex.
		hashInBytes := hash.Sum(nil)                  // use nil because we're not adding any more
		hashString := hex.EncodeToString(hashInBytes) // hash.EncodeToString is useful!

		fmt.Printf("File: %s, SHA256: %s\n", fp.Name(), hashString)
	}

	// Another _dare I say_ more elegant way of doing this is to use the
	// filepath.WalkDir() function, which itself takes a function to determine
	// what to do with the files it finds. I'll show that here.
	dirErr := filepath.WalkDir(thisDir, func(path string, file fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !file.IsDir() {
			fmt.Printf("File: %s\n", file.Name())
		}
		return nil
	})
	if dirErr != nil {
		fmt.Printf("Could not read directory")
	}
}
