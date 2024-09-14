// FooBar struct contains:
// - n: the number of times "foo" and "bar" should alternate.
// - ch: a channel to synchronize the printing of "foo" and "bar".
type FooBar struct {
	n  int      // The number of FooBar alternations.
	ch chan int // Channel for synchronization between Foo and Bar.
}

// NewFooBar initializes the FooBar struct with the given n and a new unbuffered channel.
func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:  n,              // Number of alternations.
		ch: make(chan int), // Unbuffered channel for synchronization.
	}
}

// Foo prints "foo" n times and signals the Bar function to print "bar".
func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo() // Print "foo".
		fb.ch <- 0 // Send a signal to the Bar function to proceed.
		<-fb.ch    // Wait for Bar to finish printing "bar".
	}
}

// Bar prints "bar" n times after Foo prints "foo".
func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		<-fb.ch // Wait for Foo to finish printing "foo".
		// printBar() outputs "bar". Do not change or remove this line.
		printBar() // Print "bar".
		fb.ch <- 0 // Signal Foo to proceed with the next iteration.
	}
}