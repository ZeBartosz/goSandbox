package exercises

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// Mixed debugging lab.
//
// Your job: fix the implementation functions, not the tests.
// Suggested flow for each failing test:
//   1. Predict: compile, run, block, panic, race, leak, or wrong value?
//   2. Explain why using the glossary terms.
//   3. Fix the smallest thing.
//   4. Run one test, then the race detector:
//
//      go test ./exercises -run TestMixed
//      go test -race ./exercises -run TestMixed

func mixedUnbufferedSendBeforeReceive() int {
	ch := make(chan int)

	// BUG: unbuffered send happens before any receiver can run.
	go func() {
		ch <- 42
	}()

	return <-ch
}

func TestMixedUnbufferedSendBeforeReceive(t *testing.T) {
	done := make(chan int, 1)
	go func() {
		done <- mixedUnbufferedSendBeforeReceive()
	}()

	select {
	case got := <-done:
		if got != 42 {
			t.Fatalf("got %d, want 42", got)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("function blocked; likely unbuffered send before receiver")
	}
}

func mixedBufferedChannelFull() (int, int) {
	ch := make(chan int, 1)

	go func() {
		ch <- 10
		// BUG: buffer capacity is 1, so this send blocks before receives below.
		ch <- 20
	}()

	return <-ch, <-ch
}

func TestMixedBufferedChannelFull(t *testing.T) {
	done := make(chan [2]int, 1)
	go func() {
		a, b := mixedBufferedChannelFull()
		done <- [2]int{a, b}
	}()

	select {
	case got := <-done:
		if got != [2]int{10, 20} {
			t.Fatalf("got %v, want [10 20]", got)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("function blocked; likely send to full buffered channel")
	}
}

func mixedFanInNeverCloses() <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)

		out <- 1
		out <- 2
		// BUG: receiver ranges until close, but this channel is never closed.
	}()

	return out
}

func TestMixedFanInNeverCloses(t *testing.T) {
	out := mixedFanInNeverCloses()
	seen := map[int]bool{}
	deadline := time.After(200 * time.Millisecond)

	for {
		select {
		case v, ok := <-out:
			if !ok {
				if !seen[1] || !seen[2] {
					t.Fatalf("closed too early; saw %v", seen)
				}
				return
			}
			seen[v] = true
		case <-deadline:
			t.Fatalf("timed out waiting for close; saw %v", seen)
		}
	}
}

func mixedSendAfterClose() {
	ch := make(chan int, 1)

	go func() {
		defer close(ch)

		ch <- 1
		// BUG: sending to a closed channel panics.
		ch <- 2
	}()
}

func TestMixedSendAfterClose(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("function panicked: %v", r)
		}
	}()

	mixedSendAfterClose()
}

func mixedWaitGroupMismatch() int {
	var wg sync.WaitGroup
	count := 0

	// BUG: counter says two tasks, but only one goroutine calls Done.
	wg.Add(1)
	go func() {
		defer wg.Done()
		count++
	}()

	wg.Wait()
	return count
}

func TestMixedWaitGroupMismatch(t *testing.T) {
	done := make(chan int, 1)
	go func() {
		done <- mixedWaitGroupMismatch()
	}()

	select {
	case got := <-done:
		if got != 1 {
			t.Fatalf("got %d, want 1", got)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatalf("function blocked; likely WaitGroup counter mismatch")
	}
}

func mixedRaceCounter(workers, increments int) int {
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for range increments {
				// BUG: count++ is shared state with no mutex.
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return count
}

func TestMixedRaceCounter(t *testing.T) {
	got := mixedRaceCounter(20, 10_000)
	want := 20 * 10_000
	if got != want {
		t.Fatalf("got %d, want %d; shared updates were lost", got, want)
	}
}

func mixedWrongAtomicCounter(workers, increments int) int64 {
	var count atomic.Int64
	var wg sync.WaitGroup

	wg.Add(workers)
	for range workers {
		go func() {
			defer wg.Done()
			for range increments {
				// BUG: Load and Store are each atomic, but this pair is not an atomic increment.
				count.Add(1)
			}
		}()
	}

	wg.Wait()
	return count.Load()
}

func TestMixedWrongAtomicCounter(t *testing.T) {
	got := mixedWrongAtomicCounter(20, 10_000)
	want := int64(20 * 10_000)
	if got != want {
		t.Fatalf("got %d, want %d; Load+Store lost updates", got, want)
	}
}
