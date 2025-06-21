package concepts

import (
	"fmt"
	"sync"
	"time"
)

type ProducerConsumer struct {
	buffer        []int
	mu            *sync.Mutex
	cond          *sync.Cond
	maxBufferSize int
}

func NewProducerConsumer(maxSize int) *ProducerConsumer {
	mutex := &sync.Mutex{}
	return &ProducerConsumer{
		buffer:        make([]int, 0),
		mu:            mutex,
		cond:          sync.NewCond(mutex),
		maxBufferSize: maxSize,
	}
}

func (p *ProducerConsumer) Produce() {
	fmt.Println("Producer started")
	defer fmt.Println("Producer stopped")
	for i := 0; i < 10; i++ {
		p.mu.Lock()

		for len(p.buffer) == p.maxBufferSize {
			fmt.Println("Producer waiting as buffer full")
			p.cond.Wait()
		}

		fmt.Println("Producer pushed value ", i)
		p.buffer = append(p.buffer, i)

		p.cond.Signal()
		p.mu.Unlock()

		time.Sleep(time.Millisecond * 500)
	}
}

func (p *ProducerConsumer) Consume() {
	fmt.Println("Consumer started")
	defer fmt.Println("Consumer stopped")
	for i := 0; i < 10; i++ {
		p.mu.Lock()
		for len(p.buffer) == 0 {
			fmt.Println("Consumer waiting as buffer empty")
			p.cond.Wait()
		}

		val := p.buffer[0]
		fmt.Println("Consumed value : ", val)

		p.buffer = p.buffer[1:]

		p.cond.Signal()
		p.mu.Unlock()

		time.Sleep(time.Second * 1)
	}
}
