package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type tokenBucket struct {
	capacity     int
	tokens       int
	refilRate    int
	blockedCount int //no of request
	lockUntil    time.Time
	lastRefill   time.Time
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

func middleware(tb *tokenBucket, fn func()) {
	
	if time.Now().Before(tb.lockUntil){
		fmt.Printf("🚧 User Locked Until %s\n", tb.lockUntil.Format("15:04:05"))aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa
	}
	
	if !tb.AllowRequest() {
		tb.blockedCount++

		if tb.blockedCount >= 5 {
			tb.lockUntil = time.Now().Add(time.Minute)
			fmt.Printf(
				"🔒 USER LOCKED for 1 minute (until %s)\n",
				tb.lockUntil.Format("15:04:05"),
			)
			return
		}
		fmt.Printf(
			"SECURITY ALERT: Blocked=%d Time=%s ❌\n",
			tb.blockedCount,
			time.Now().Format("15:04:05"),
		) 
		return
	}
	fn()
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
	for i := 1; i <= 10; i++ {

		// For delaying between the request
		time.Sleep(time.Millisecond * 200)

		fmt.Println("\n-------------------------")
		fmt.Println("[DEBUG] token bucket state updated")
		middleware(tb, func() {
			fmt.Printf("Request %d Allowed ✅\n", i)
			apiFetch(i)
		})
	}
}
