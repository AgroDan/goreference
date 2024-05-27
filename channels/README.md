# Concurrency

Go is particularly unique in its implementation of concurrency. Sure this can be done in Python, but I've found that compared to Go, it is incredibly intense to set up and get working. Go has a built-in directive specifically designed for concurrency, the `go` directive. Any function called with `go` in front of it will split off to run in parallel with the remainder of the application. However, just like any other language, there are a few "gotchas" to concurrent processes. First of all, when the parent process exits, it kills all child processes regardless if they are done computing. To help with that, the common libraries include a `WaitGroup` object, which will block the main process until all waitgroups decrement to zero.

Additionally, to communicate _between_ concurrent processes, Go offers a datatype specifically designed for inter-process communication: Channels!

Channels are a unique (in my opinion anyway) datatype in go that sort of works like a queue. The big difference here is that go channels will block until another side receives the data stored within the channel. This is the first step to understanding parallel processing within Go. The basic idea is that one process will insert data into the channel and then wait for another alternate process to pull the data from the channel before Go will continue with the original process.

And not to forget `mutex` as well! Sometimes processes don't need to communicate with each other, but they will all modify and populate the same data. To prevent a race condition unique to parallel processing, Mutual Exclusion is necessary to ensure that data is written to or read from one at a time, preventing oddball behavior like two processes writing data to the same memory location at the same time, corrupting the output and potentially doing bad things.

## Buffered vs Unbuffered Channels

In the preceeding paragraph I only referred to an unbuffered channel. You can create a _buffered_ channel by specifying its size. But note that Go will block until it receives what it needs to continue processing from the channel. Meaning that a program like this:

```go
func main() {
    ch := make(chan int, 2)
    ch <- 1
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
```

Will error out because it deadlocked itself waiting for the channel to populate so it could print the output a second time.

## Concurrency and WaitGroups

Because channels are typically used with concurrency, waitgroups are necessary to ensure that the program does not terminate before the workers are done doing whatever it is they are doing. If the main program terminates, all the child processes it spawns will also terminate whether they're done or not. Using waitgroups allows the caller to block until all the subroutines are finished before moving on.

More information in the code itself!