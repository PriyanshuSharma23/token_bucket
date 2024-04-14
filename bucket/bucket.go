// Package bucket implements a token bucket
package bucket

import (
	"time"
)

type Bucket struct {
	r int       // fill rate
	b int       // current size
	c int       // capacity
	t time.Time // last operation time
}

func (b *Bucket) sync() {
	ts := time.Now()
	seconds := int(ts.Sub(b.t).Seconds())

	if seconds <= 0 {
		return
	}

	b.t = b.t.Add(time.Duration(seconds) * time.Second)
	b.b = min(b.b+b.r*seconds, b.c)
}

// Returns the maximum capcaity if the bucket
func (b *Bucket) Cap() int {
	return b.c
}

func (b *Bucket) Size() int {
	b.sync()
	return b.b
}

func NewBucket(r, c int) *Bucket {
	if r < 0 {
		panic("expected r to be greater than 0")
	}

	if c < 0 {
		panic("expected c to be greater than 0")
	}

	b := &Bucket{
		r: r,
		b: c,
		c: c,
		t: time.Now(),
	}

	return b
}

func (b *Bucket) Check(n int) bool {
	b.sync()

	if b.b < n {
		return false
	}

	b.b -= n
	return true
}
