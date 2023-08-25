package main

import (
	"sync"
	"time"
)



type Limiter struct {
	mu sync.Mutex
	rate int
	bucketSize int
	nTokens int
	lastToken time.Time
}


func NewLimiter(rate,limit int) *Limiter {
	return &Limiter{
		rate: rate,
		bucketSize: limit,
		nTokens: limit,
		lastToken: time.Now(),
	}

}


func (r *Limiter) Wait() {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.nTokens>0{
		r.nTokens--
		return 
	}

	tElapsed:=time.Since(r.lastToken)
	period:=time.Second/time.Duration(r.rate)
	nTokens:=tElapsed.Nanoseconds()/period.Nanoseconds()
	r.nTokens = int(nTokens)
	if r.nTokens>r.bucketSize{
		r.nTokens=r.bucketSize
	}
	r.lastToken = r.lastToken.Add(time.Duration(nTokens))

	if r.nTokens>0{
		r.nTokens--
		return 
	}
	next := r.lastToken.Add(period)
	wait := next.Sub(time.Now())

	if wait >= 0 {
		time.Sleep(wait)
	}
	r.lastToken = next

}


func main() {
	
}