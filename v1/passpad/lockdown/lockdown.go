package lockdown

import (
	"math"
	"time"
)

type state struct {
	last              time.Time
	exponentialFactor float64
}

var m = make(map[string]*state)

func init() {
	go func() {
		for {
			for name, s := range m {
				if !IsLocked(name) {
					if s.exponentialFactor > 0 {
						s.exponentialFactor = s.exponentialFactor - 1
					}
				}
			}
			time.Sleep(1 * time.Hour)
		}
	}()
}

func getState(username string) *state {
	s, ok := m[username]
	if !ok {
		s = &state{exponentialFactor: 0}
		m[username] = s
	}
	return s
}

func Fail(username string) {
	s := getState(username)
	s.last = time.Now()
	s.exponentialFactor = s.exponentialFactor + 1
}

func IsLocked(username string) bool {
	s := getState(username)
	t := s.last
	factor := math.Pow(2, s.exponentialFactor)
	t = t.Add(time.Second * time.Duration(factor))
	if t.Before(time.Now()) {
		return false
	} else {
		return true
	}
}
