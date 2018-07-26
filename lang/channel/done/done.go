package main

import (
	"fmt"
	"sync"
)

type worker struct {
	in   chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{
		in: make(chan int),
		done: func() {
			wg.Done()
		},
	}
	go doWork(id, w)
	return w
}
func chanDemo() {
	var wg sync.WaitGroup
	wg.Add(20) //20个任务

	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'a' + i
	}
	for i := 0; i < 10; i++ {
		workers[i].in <- 'A' + i
	}

	//wait for all of them
	wg.Wait()
}

func doWork(id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n",
			id, n)
		go func() {
			w.done()
		}()
	}
}

func main() {
	chanDemo()
}
