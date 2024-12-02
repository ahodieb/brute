package histogram

import (
	"encoding/json"
	"fmt"
)

// Histogram  provides frequency counting of values
// it wraps a map[T]int and provides append and get methods
// it eliminates the boilerplate of declaring maps, checking and incrementing
//
// Not concurrency safe, check and increment is not atomic
type Histogram[T comparable] struct {
	m map[T]int
}

// New creates a new Histogram
func New[T comparable]() *Histogram[T] {
	return &Histogram[T]{m: make(map[T]int)}
}

// FromItems creates a new Histogram of frequency of the items passed in
func FromItems[T comparable](items ...T) *Histogram[T] {
	return FromSlice(items)
}

// FromSlice creates a new Histogram of frequency of the items in the slice
func FromSlice[T comparable](s []T) *Histogram[T] {
	h := New[T]()
	h.Add(s...)
	return h
}

// Add adds a new value to histogram
func (h *Histogram[T]) Add(values ...T) {
	for _, v := range values {
		h.m[v]++
	}
}

// Count returns the count of specified value, returns 0 if not seen before
func (h *Histogram[T]) Count(v T) int {
	return h.m[v]
}

// String returns string representation of Histogram
func (h *Histogram[T]) String() string {
	j, err := json.Marshal(h.m)
	if err != nil {
		return fmt.Sprintf("error marshalling: %v, %v", h.m, err)
	}

	return string(j)
}
