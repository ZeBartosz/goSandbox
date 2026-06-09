package exercises

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Lab rule: fix the implementation functions below, then run:
//
//   go test ./exercises
//   go test -race ./exercises
//
// Start with one test at a time, e.g.:
//
//   go test ./exercises -run TestWaitGroupWaitsForWorkers

func brokenWaitGroupWaitsForWorkers() int {
	// BUG: this starts work but returns before the goroutine is finished.
	// Goal: use sync.WaitGroup so the function always returns 42.
	var wg sync.WaitGroup
	result := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		result = 42
	}()
	wg.Wait()
	return result
}

func TestWaitGroupWaitsForWorkers(t *testing.T) {
	for range 100 {
		if got := brokenWaitGroupWaitsForWorkers(); got != 42 {
			t.Fatalf("got %d, want 42; goroutine was not coordinated", got)
		}
	}
}

func brokenChannelFanIn() <-chan int {
	// BUG: sends two values but never closes the output channel.
	// Goal: use a WaitGroup and close(out) after all sending goroutines finish.
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		out <- 1
	}()

	go func() {
		defer wg.Done()
		out <- 2
	}()

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func TestChannelFanInCloses(t *testing.T) {
	out := brokenChannelFanIn()
	seen := map[int]bool{}
	deadline := time.After(500 * time.Millisecond)

	for {
		select {
		case v, ok := <-out:
			if !ok {
				if !seen[1] || !seen[2] {
					t.Fatalf("channel closed, but saw values %v; want 1 and 2", seen)
				}
				return
			}
			seen[v] = true
		case <-deadline:
			t.Fatalf("timed out waiting for channel to close; likely goroutine/channel coordination bug")
		}
	}
}

func brokenMutexCounter(workers int, increments int) int {
	// BUG: count++ is shared memory accessed by many goroutines without a lock.
	// Goal: use sync.Mutex and sync.WaitGroup so the final count is exact.
	count := 0
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for range increments {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return count
}

func TestMutexProtectsSharedCounter(t *testing.T) {
	got := brokenMutexCounter(20, 10_000)
	want := 20 * 10_000
	if got != want {
		t.Fatalf("got %d, want %d; shared counter updates were lost", got, want)
	}
}

func brokenAtomicCounter(workers int, increments int) int64 {
	// BUG: Load + Store is not an atomic increment; goroutines can overwrite each other.
	// Goal: use atomic.Int64.Add, not Load followed by Store.
	var count atomic.Int64
	var wg sync.WaitGroup

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for range increments {
				count.Add(1)
			}
		}()
	}

	wg.Wait()
	return count.Load()
}

func TestAtomicCounterUsesAtomicAdd(t *testing.T) {
	got := brokenAtomicCounter(20, 10_000)
	want := int64(20 * 10_000)
	if got != want {
		t.Fatalf("got %d, want %d; Load+Store was not a safe increment", got, want)
	}
}
