package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrmenting the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := NewCounter()
		numInc := 3
		for i := 0; i < numInc; i++ {
			counter.Inc()
		}

		assertCounter(t, counter, numInc)
	})
	t.Run("it runs safely concurrently", func(t *testing.T) {
		counter := Counter{}
		numInc := 1000

		var wg sync.WaitGroup
		wg.Add(numInc)

		for i := 0; i < numInc; i++ {
			go func(w *sync.WaitGroup) {
				counter.Inc()
				w.Done()
			}(&wg)
		}
		wg.Wait()
	})
}
func assertCounter(t *testing.T, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
