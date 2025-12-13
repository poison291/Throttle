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
func apiFetch(number int) {
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
	tb := newTokenBucket(4, 1)
	for i := 1; i <= 5; i++ {

		// For delaying between the request
		// time.Sleep(time.Second)
		
		fmt.Println("\n-------------------------")
		fmt.Println("Tokens available before request:", tb.tokens)

		result := tb.AllowRequest()
		if result {
			fmt.Printf("Request %d Allowed ✅ \n", i)
			apiFetch(i)
		} else {
			fmt.Printf("Request %d Blocked ❌ \n", i)
			fmt.Print("The Request Block due to insufficent token. ")
		}
	}

}
