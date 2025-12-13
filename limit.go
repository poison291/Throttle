package main

import (
	"fmt"
	"time"
)

type tokenBucket struct {
	capacity   int
	tokens     int
	refilRate  int
	lastRefill time.Time
}

func newTokenBucket(capacity int, refillrate int) *tokenBucket {
	return &tokenBucket{
		capacity:   capacity,
		tokens:     capacity,
		refilRate:  refillrate,
		lastRefill: time.Now(),
	}
}

func (tb *tokenBucket) refil() {
	now := time.Now()
	elapsed := int(now.Sub(tb.lastRefill).Seconds())
	tb.tokens += elapsed * tb.refilRate
	
	if tb.tokens > tb.capacity{
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

func (tb *tokenBucket) AllowRequest() bool{
	tb.refil()
	if tb.tokens > 0 {
		tb.tokens = tb.tokens - 1
		fmt.Println("The request Passed")
		return true
	}
	fmt.Println("The Request Block due to insufficent token")
	return false
}

func main() {
	tb := newTokenBucket(3, 1)
	for i := 0; i < 4; i++{
	result := 	tb.AllowRequest()
		fmt.Println(result)
	}
	
}
