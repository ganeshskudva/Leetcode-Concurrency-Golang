package main

import (
	"fmt"
	"sync"
)

// done struct{} is used as a signal for the channels.
// Empty structs take up no space in memory, so they're ideal for signaling.
var done struct{}

// Streams struct holds two channels and a WaitGroup.
// The channels coordinate the execution order of the goroutines.
type Streams struct {
	firstStream  chan struct{}   // Channel for signaling that first() is complete.
	secondStream chan struct{}   // Channel for signaling that second() is complete.
	wg           *sync.WaitGroup // Pointer to a WaitGroup for synchronization.
}

// first method prints "first" and signals via firstStream when done.
func (s *Streams) first() {
	defer s.wg.Done() // Ensure WaitGroup counter is decremented when this method returns.

	fmt.Println("first")  // Print "first".
	s.firstStream <- done // Send signal to firstStream channel indicating that first() is complete.
}

// second method waits for first() to complete, prints "second", and signals via secondStream.
func (s *Streams) second() {
	defer s.wg.Done() // Ensure WaitGroup counter is decremented when this method returns.

	<-s.firstStream        // Block execution until signal is received from firstStream (i.e., first() is done).
	fmt.Println("second")  // Print "second".
	s.secondStream <- done // Send signal to secondStream channel indicating that second() is complete.
}

// third method waits for second() to complete, then prints "third".
func (s *Streams) third() {
	defer s.wg.Done() // Ensure WaitGroup counter is decremented when this method returns.

	<-s.secondStream     // Block execution until signal is received from secondStream (i.e., second() is done).
	fmt.Println("third") // Print "third".
}

func main() {
	// Initialize Streams struct with buffered channels and a WaitGroup.
	st := &Streams{
		firstStream:  make(chan struct{}, 1), // Buffered channel allows non-blocking send.
		secondStream: make(chan struct{}, 1), // Buffered channel allows non-blocking send.
		wg:           &sync.WaitGroup{},      // Initialize WaitGroup for synchronizing goroutines.
	}

	st.wg.Add(3) // Add three tasks to the WaitGroup (one for each method).

	// Launch the methods as goroutines. These will run concurrently but will be controlled by the channels.
	go st.third()  // third() waits for the signal from secondStream.
	go st.second() // second() waits for the signal from firstStream.
	go st.first()  // first() runs immediately and signals firstStream.

	st.wg.Wait() // Wait for all three goroutines to finish before exiting main.
}
