package main

import "time"

type RateLimiter interface {
	Wait()
}

type ChannelRate struct {
	bucket chan struct{}
	ticker *time.Ticker
	done   chan struct{}
}

func NewChannelRate(rate float64, limit int) *ChannelRate {
	ret := &ChannelRate{
		bucket: make(chan struct{}, limit),
		ticker: time.NewTicker(time.Duration(1 / rate * 1000000000)),
		done:   make(chan struct{}),
	}

	for i := 0; i < limit; i++ {
		ret.bucket <- struct{}{}
	}
	go func() {
		for {
			select {
			case <-ret.done:
				return
			case <-ret.ticker.C:
				select {
				case ret.bucket <- struct{}{}:
				default:
				}

			}
		}
	}()
	return ret
}

func (r *ChannelRate) Wait() {
	<-r.bucket
}

func (r *ChannelRate) Close() {
	close(r.done)
	r.ticker.Stop()
}

func main() {

}
