package bucket

import (
	"sync"
	"testing"
	"time"
)

func TestBucket(t *testing.T) {
	t.Run("Burst", func(t *testing.T) {
		trueCounts := 0
		b := NewBucket(0, 10)

		for i := 0; i < 10; i++ {
			if b.Check(1) {
				trueCounts++
			}
		}

		if trueCounts != 10 {
			t.Logf("expected %d trueCounts got %d", 10, trueCounts)
			t.Fail()
		}
	})

	t.Run("Refill", func(t *testing.T) {
		b := NewBucket(2, 10)

		b.Check(4)              // b.b = 6
		time.Sleep(time.Second) // b.b += 2

		if b.Size() != 8 {
			t.Logf("expected %d size of bucket got %d", 8, b.Size())
			t.Fail()
		}
	})

	t.Run("Thread safe", func(t *testing.T) {
		b := NewBucket(2, 10)

		var wg sync.WaitGroup

		var currSize int = int(^uint(0) >> 1) // INT_MAX

		for i := 0; i < 2; i++ {
			wg.Add(1)
			currSize = min(func() int {
				b.Check(2)
				return b.Size()
			}(), currSize)
		}

		wg.Done()

		if currSize != 6 {
			t.Logf("expected the size to be %d but it is %d", 6, currSize)
			t.Fail()
		}
	})

	t.Run("ExhaustCapacity", func(t *testing.T) {
		b := NewBucket(2, 5)

		// Exhaust the bucket's capacity
		for i := 0; i < 5; i++ {
			if !b.Check(1) {
				t.Error("expected Check to return true when bucket has enough tokens")
			}
		}

		// Ensure bucket is empty after exceeding capacity
		if b.Size() != 0 {
			t.Errorf("expected bucket size to be 0 after exceeding capacity, got %d", b.Size())
		}
	})

	t.Run("NegativeRate", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected NewBucket to panic with negative fill rate")
			}
		}()

		_ = NewBucket(-2, 10)
	})

	t.Run("NegativeCapacity", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("expected NewBucket to panic with negative capacity")
			}
		}()

		_ = NewBucket(2, -10)
	})

	t.Run("Non integer duration in milliseconds", func(t *testing.T) {
		b := NewBucket(2, 5)

		b.Check(2)
		time.Sleep(500 * time.Millisecond)
		b.Check(3)
		time.Sleep(500 * time.Millisecond)

		if b.Size() != 2 {
			t.Errorf("expected size to be %d but it is %d", 2, b.Size())
		}
	})

	t.Run("Non integer duration in milliseconds 2", func(t *testing.T) {
		b := NewBucket(2, 5)

		b.Check(1)
		time.Sleep(200 * time.Millisecond)
		b.Check(2)
		time.Sleep(500 * time.Millisecond)
		b.Check(2)
		time.Sleep(200 * time.Millisecond)

		if b.Size() != 0 {
			t.Errorf("expected size to be %d but it is %d", 0, b.Size())
		}
	})
}
