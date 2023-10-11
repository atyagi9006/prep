package main

import (
	"fmt"
	"math"
	"time"
)

//rate limiting algo
//fixed rate limiting
//sliding window
//leaky bucket
//token bucket algo
//20rpm

type tokenBucket struct {
	tokens         float64
	mxTokens       float64
	RefillRate     float64
	lastRefilltime time.Time
}

func NewtokenBucket(maxToken, RefillRate float64) *tokenBucket {
	return &tokenBucket{
		tokens:         maxToken,
		mxTokens:       maxToken,
		RefillRate:     RefillRate,
		lastRefilltime: time.Now(),
	}
}

func (tb *tokenBucket) refill() {
	now := time.Now()
	duration := now.Sub(tb.lastRefilltime)
	tokenstoadd := tb.RefillRate * float64(duration.Seconds())
	tb.tokens = math.Min(tb.tokens+tokenstoadd, tb.mxTokens)
	tb.lastRefilltime = now
}

func (tb *tokenBucket) request(tokens float64) bool {
	tb.refill()
	if tokens <= tb.tokens {
		tb.tokens -= tokens
		return true
	}
	return false
}
func main() {
	tb := NewtokenBucket(10, 1)
	for i := 0; i < 20; i++ {
		fmt.Printf("required: %d, granted : %v \n", i+1, tb.request(1))
		time.Sleep(50* time.Millisecond)
	}
}
