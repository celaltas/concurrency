package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)



func philosopher(index int, firstFork,secondFork *sync.Mutex){
	for {
		fmt.Printf("Philosopher %d is thinking\n",index)
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(1000)))
		firstFork.Lock()
		secondFork.Lock()
		fmt.Printf("Philosopher %d is eating\n",index)
		time.Sleep(time.Millisecond*time.Duration(rand.Intn(1000)))
		firstFork.Unlock()
		secondFork.Unlock()

	}
}


func main() {
	forks:=[5]sync.Mutex{}
	go philosopher(0,&forks[4],&forks[0])
	go philosopher(1,&forks[0],&forks[1])
	go philosopher(2,&forks[1],&forks[2])
	go philosopher(3,&forks[2],&forks[3])
	go philosopher(4,&forks[3],&forks[4])

	select{}
}