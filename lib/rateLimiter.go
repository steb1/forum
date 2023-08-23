package lib

import (
	"net/http"
	"sync"
	"time"
)

type RateLimiter struct {
	WindowDuration time.Duration
	Limit          int
	Requests       map[string][]time.Time
	Mutex          sync.Mutex
}

func NewRateLimiter(window time.Duration, limit int) *RateLimiter {
	return &RateLimiter{
		WindowDuration: window,
		Limit:          limit,
		Requests:       make(map[string][]time.Time),
	}
}

func (rl *RateLimiter) Wrap(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		remoteIP := r.RemoteAddr

		rl.Mutex.Lock()
		defer rl.Mutex.Unlock()

		// Clear expired request records
		now := time.Now()
		validRequests := rl.Requests[remoteIP][:0]
		for _, t := range rl.Requests[remoteIP] {
			if now.Sub(t) <= rl.WindowDuration {
				validRequests = append(validRequests, t)
			}
		}
		rl.Requests[remoteIP] = validRequests

		if len(rl.Requests[remoteIP]) >= rl.Limit {
			http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
			return
		}

		rl.Requests[remoteIP] = append(rl.Requests[remoteIP], time.Now())

		next(w, r)
	}
}

