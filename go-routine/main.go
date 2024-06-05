package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	counter = 0
	lock    sync.Mutex
)

func main() {
	/*
		for i := 0; i < 20; i++ {
			go incr()
		}
		time.Sleep(time.Millisecond * 10)
	*/
	/*
		go func() {
			lock.Lock()
		}()
		time.Sleep(time.Millisecond)
		lock.Lock()
	*/

	c := make(chan int, 100)
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	/*
		for {
			c <- rand.Int()
			fmt.Println(len(c))
			time.Sleep(time.Millisecond * 50)
		}
	*/
	for {
		select {
		case c <- rand.Int():
		case <-time.After(time.Millisecond * 100):
			fmt.Println("timed out")
		default:
			fmt.Println("Dropped")
		}
		fmt.Println(len(c))
		time.Sleep(time.Millisecond * 50)
	}

}

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d received data: %d\n", w.id, data)
		time.Sleep(time.Millisecond * 500)
	}
}

func incr() {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}
