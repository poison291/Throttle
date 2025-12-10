package main

import (
	"fmt"
	"time"
)

type tokenBucket struct {
	capacity int
	tokens int
	refilRate int
	lastRefill time.Time
}

func newTokenBucket(capacity int, refillrate int) *tokenBucket{
	return &tokenBucket{
		capacity: capacity,
		tokens: capacity,
		refilRate: refillrate,
		lastRefill: time.Now(),
	}
}


func main(){
	fmt.Print("Limit Bucket")
}