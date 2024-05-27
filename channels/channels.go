package main

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func SleepAndWakeUp(seconds int, ch chan string) {
	// pulling the channel first, this will block
	// this function until something populates it
	shout := <-ch
	retval := "Sleeping..."
	time.Sleep(time.Duration(seconds) * time.Second)
	retval += shout
	ch <- retval
}

// This function pretty much only opens the file and
// hashes it. It is completely channel un-aware!
func HashFile(fp string) (string, error) {
	thisFile, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	// always defer file closure!
	defer thisFile.Close()

	hash := sha512.New()
	if _, err := io.Copy(hash, thisFile); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

// Note the directions of the channels stated in the parameter list, as well as
// the pointer to the waitgroup that we define in the main function! We know that
// we can only ingest items from the files channel, and output only to the results
// channel.
func worker(files <-chan string, results chan<- map[string]string, wg *sync.WaitGroup) {
	// This basically says once we're done with everything, decrement the
	// waitGroup value regardless of anything.
	defer wg.Done()

	for file := range files {
		hash, err := HashFile(file)
		if err != nil {
			fmt.Printf("Failed to hash file %s: %v\n", file, err)
			continue
		}
		// put the file name and the hash value into the map
		results <- map[string]string{file: hash}
	}

	// and since we're not returning anything and the defer directive handles the
	// waitGroup decrement, we can just exit here.
}

func main() {
	// All channels must be created with the make() builtin function.
	ch := make(chan string, 2)
	ch <- "Hello World!"
	fmt.Println(<-ch)

	// Note however that I created a _buffered_ channel for this simple
	// statement. The reason is because if I chose an unbuffered channel,
	// Go would deadlock once it populated the channel, since it will
	// block until another process pulls from that channel -- and no
	// parallel processes were created.

	// So let's create an unbuffered channel and work with go subroutines

	// First, create the unbuffered channel. This will take the input
	// and then populate the channel with the result again!
	ch2 := make(chan string)

	// Kick off the subroutine. This will block until we populate the channel.
	go SleepAndWakeUp(5, ch2)

	// Obligatory "here's what i'm about to do" text
	fmt.Printf("First, I'll sleep for 3 seconds and then fill the channel.\n")
	fmt.Printf("At this point, the go routine will be running but blocking,\n")
	fmt.Printf("Waiting for me to input data.\n")

	// Arbitrary sleep timer
	time.Sleep(3 * time.Second)

	// Populate the channel. Note the use of <- and -> to send and receive.
	// Generall you can refer to the value of a channel with <-ch, which will
	// use the value of that channel AND pop the result off of it.
	ch2 <- "Hello World! Again!"

	// result := <-ch2
	// fmt.Printf("Result: %s\n", result)

	// Now something particularly useful when dealing with channels is working on
	// a large amount of files. If you don't do this with channels and routines,
	// you will operate sequentially and this can potentially take a lot longer.
	// To demonstrate this, let's take the directory filled with large files and
	// sequentially SHA512 each file and time it.

	myFileDir := "./channels/randomfiles"

	t1Now := time.Now()
	fileList, err := os.ReadDir(myFileDir)
	if err != nil {
		panic("Could not read dir")
	}

	// Now we'll open each file
	for _, fp := range fileList {
		if fp.IsDir() {
			fmt.Printf("Directory found, skipping...\n")
			continue
		}

		fullFilePath := filepath.Join(myFileDir, fp.Name())

		oFile, err := os.Open(fullFilePath)
		if err != nil {
			fmt.Printf("Could not open file: %s\n", oFile.Name())
			continue
		}

		hash := sha512.New()
		if _, err := io.Copy(hash, oFile); err != nil {
			fmt.Printf("Could not hash file: %s\n", oFile.Name())
			oFile.Close()
			continue
		}
		oFile.Close()

		// now compute the print-friendly hash
		hashString := hex.EncodeToString(hash.Sum(nil))
		fmt.Printf("File: %s, SHA512: %s\n", oFile.Name(), hashString)
	}

	fmt.Printf("Time to compute hashes sequentially: %s\n", time.Since(t1Now))

	// Doesn't take THAT long...but what if we spin up 5 different functions
	// that all pull from the same channel to work all in parallel?

	t2Now := time.Now()

	// This will also allow me to play with anonymous functions! Since you have
	// to preface whatever routine you want to do with "go" to perform concurrency,
	// you can wrap it in a lambda function which will just execute on the spot.

	// A prime example is using the filepath.WalkDir() function, which itself takes
	// a filepath.WalkDirFunc() method as an input parameter! THAT function itself
	// takes 3 arguments: (path string, d dirEntry, err error). It's within THAT
	// function that you make the determination of what you will do with the value of
	// each entry it finds while walking through a directory. This winds up looking
	// very similar to what you'd see in a javascript application. At least that's
	// what I see.

	// Anyway, let's set up two channels. The first will be a channel of ten items
	// (arbitrary number) that will hold a string. This will be the channel where I
	// will put the location of files that should be hashed.
	filesChannel := make(chan string, 10)

	// The second channel will be the results channel, also 10 items long. The workers
	// will put the results of hashing into this channel, which will be read by an
	// anonymous function displayed later. Also this will be a map just to be silly.
	resultsChannel := make(chan map[string]string, 10)

	// Also, I will be utilizing the waitGroup library! This is crucial to ensure that
	// work is done before terminating the program!
	var wg sync.WaitGroup

	// At this point i'll be using two functions as labelled above. a worker function
	// which will be "concurrency aware," and the actual hashing function which will
	// get called by the worker function. So now let's spawn 5 worker processes
	// concurrently which will wait for input because they will block independently
	// on each file input to the filesChannel.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go worker(filesChannel, resultsChannel, &wg)
	}

	// Now I'll use TWO anonymous functions! The first has the over-arching instruction
	// to populate the filesChannel, and the second will be passed to the filepath.WalkDir
	// function, since that function itself requires a function to operate. Also note the
	// use of fs.DirEntry, since that object satifies the interface for the DirEntry object
	// in the filepath module.
	go func() {
		err := filepath.WalkDir(myFileDir, func(path string, info fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if !info.IsDir() {
				// as long as this isn't a directory, this is
				// a file that we want to hash. So we'll add
				// it to the filesChannel.
				filesChannel <- path
			}
			return nil
		})
		// Now we're in the filepath.WalkDir function's return, which would populate the first "err"
		// variable. We'll handle any errors as we always do.
		if err != nil {
			fmt.Printf("Failed to walk directory: %v\n", err)
		}

		// So we've populated everything we need at this point, so let's close the channel.
		close(filesChannel)

		// and don't forget, this is an anonymous function so an additional () at the end will invoke
		// the function as if it were a regular old function that we're calling.
	}()

	// finally, we'll block at this channel until all the waitgroups are done! This is a concurrent
	// function because it will block on the side, allowing us to continue processing. Again we'll
	// use an anonymous function.
	go func() {
		wg.Wait()
		close(resultsChannel)
	}()

	// Now we can process the results as soon as we receive them in real time!
	for result := range resultsChannel {
		for k, v := range result {
			fmt.Printf("File: %s, Hash: %s\n", k, v)
		}
	}

	fmt.Printf("Second hash took: %s\n", time.Since(t2Now))

	// On my PC here, the first half took 24 seconds to complete. The second hash running concurrently took
	// only 941ms! Huge difference when working on multiple files!

	// Now let's work on Mutexes. I'll call the function here but all the code
	// and commentary are in mutex.go.
	GuestbookInAction()
}
