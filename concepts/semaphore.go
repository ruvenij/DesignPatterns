package concepts

type Semaphore struct {
	BufferedSemaphore chan struct{}
}

func NewSemaphore(bufferSize int) *Semaphore {
	return &Semaphore{BufferedSemaphore: make(chan struct{}, bufferSize)}
}

func (s *Semaphore) Acquire() {
	s.BufferedSemaphore <- struct{}{}
}

func (s *Semaphore) Release() {
	<-s.BufferedSemaphore
}
