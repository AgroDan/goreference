package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// This will explain how Mutexes work in Go. I will have a
// Guestbook struct which will accept a signature to append
// to a giant string. Note that this is _extremely_ contrived
// and I would never develop it this way, but this is a good
// way of explaining Mutexes.

type Guestbook struct {
	mut         sync.Mutex
	book        string
	lastSigned  time.Time
	totalSigned int
}

// just a constructor function
func NewGuestbook() Guestbook {
	return Guestbook{}
}

// This function will sign the guestbook with the specified signature
// which it will happily append to its value. To ensure that it is not
// overwritten by a concurrent process, I'll use mutexes.
func (g *Guestbook) Sign(signature string) {
	// lock the mutex
	g.mut.Lock()

	// now defer unlocking until after the function exits!
	defer g.mut.Unlock()

	// proceed with signing
	g.book += fmt.Sprintf("\n%s", signature)
	g.lastSigned = time.Now()

	// Don't forget to increment the counter!
	g.totalSigned++
}

// This function prints the guestbook. Since it could potentially be
// written to at the time it was printing, we'll lock it as well.
func (g *Guestbook) Print() {
	// lock it up
	g.mut.Lock()

	// defer unlocking
	defer g.mut.Unlock()

	fmt.Printf("Guestbook contents:\n%s", g.book)
	fmt.Printf("\nLast Signed: %s\n", g.lastSigned.String())
	fmt.Printf("Total signed: %d\n", g.totalSigned)
}

func GuestbookInAction() {
	// This is a contrived function that just shows how mutual exclusives
	// prevent data from being corrupted in its shared state.

	// First, let's set up the waitgroup
	var wg sync.WaitGroup

	// declare the guestbook object
	gb := NewGuestbook()

	// How many workers should I make?
	totalWorkers := 100

	// Add the waitgroups
	wg.Add(totalWorkers)

	// Now let's fire it up.
	for i := 0; i < totalWorkers; i++ {
		// anonymous function ho!
		go func(num int) {
			defer wg.Done()
			// sleep for a random number of milliseconds
			randNum := rand.Int() % 1000
			time.Sleep(time.Duration(randNum * int(time.Millisecond)))
			gb.Sign(fmt.Sprintf("Hello from worker #%d", num))
		}(i)
	}

	// block until all workers are done
	wg.Wait()

	// Print the guestbook!
	gb.Print()
}
