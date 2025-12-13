package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	if tb.tokens > tb.capacity {
		tb.tokens = tb.capacity
	}
	tb.lastRefill = now
}

func (tb *tokenBucket) AllowRequest() bool {
	tb.refil()
	if tb.tokens > 0 {
		tb.tokens = tb.tokens - 1
		return true
	}
	return false
}
func apiFetch (number int) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/todos/%d", number)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while fetching Api:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error while reading body:", err)
		return
	}
	fmt.Println(string(body))
}

func main() {
	tb := newTokenBucket(3, 1)
	for i := 0; i <= 3; i++ {
		result := tb.AllowRequest()

		// For delaying between the request
		// time.Sleep(time.Second)
		
		
		fmt.Println("\n-------------------------")
		fmt.Println("Request number:", i+1)
		
		if result {
			fmt.Print("The request Passed. ")
			apiFetch(i+1 )
		} else {
			fmt.Print("The Request Block due to insufficent token. ")
		}
	}

}
