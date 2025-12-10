package main

import (
	"fmt"
	"time"
)

type rateLimiter struct {
	time map[string]  []time.Time
	limit int
	winDuration time.Duration	
}

func main(){
	fmt.Print()
}