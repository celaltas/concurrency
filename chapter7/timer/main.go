package main

import (
	"fmt"
	"time"
)



type TimerMock struct {
	C chan<- time.Time
} 


func NewTimerMock(d time.Duration) *TimerMock {
	t:= &TimerMock{
		C: make(chan time.Time,1),
	}
	go func ()  {
		time.Sleep(d)
		t.C <- time.Now()
	}()
	return t
}


func main() {
	timer:=time.NewTimer(100*time.Millisecond)
	timeout:=make(chan struct{})
	go func() {
		<- timer.C
		close(timeout)
		fmt.Println("Timeout")
	}()
	x:=0
	done:=false
	for !done {
		select {
			case <-timeout:
				done=true
			default:
		}
		time.Sleep(time.Millisecond)
		x++
	}
	fmt.Println("x=",x)
}