package main

import (
	"fmt"
	"math/rand"
)

type Work struct {
	value int
}

type Result struct {
	value int
}

func main() {

	workCh := make(chan Work)
	resultCh:=make(chan Result)
	done:=make(chan bool)

	workQueue := make([]Work, 100)

	for i := range workQueue {
		workQueue[i].value = rand.Intn(101)
	}

	for i := 0; i < 10; i++ {
		go func() {
			for work:= range workCh {
				res:=Result{value:work.value*2}
				resultCh<-res
			}
		}()
	}

	results:=make([]Result,0)

	go func() {
		for i := 0; i < len(workQueue); i++ {
			results = append(results, <-resultCh)
		}
		done <- true
	}()

	for _, work := range workQueue {
		workCh <- work
	}
	close(workCh)
	<-done
	fmt.Println(results)

}
