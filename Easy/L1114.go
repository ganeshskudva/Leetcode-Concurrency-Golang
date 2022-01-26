var done struct{}

type Streams struct {
	firstStream  chan struct{}
	secondStream chan struct{}
	wg           *sync.WaitGroup
}

func (s Streams) first() {
	defer s.wg.Done()

	fmt.Println("first")
	s.firstStream <- done
}

func (s Streams) second() {
	defer s.wg.Done()

	<-s.firstStream
	fmt.Println("second")
	s.secondStream <- done
}

func (s Streams) third() {
	defer s.wg.Done()

	<-s.secondStream
	fmt.Println("third")
}

func main() {
	st := Streams{
		firstStream:  make(chan struct{}),
		secondStream: make(chan struct{}),
		wg:           &sync.WaitGroup{},
	}

	st.wg.Add(3)
	go st.third()
	go st.second()
	go st.first()

	st.wg.Wait()
}
