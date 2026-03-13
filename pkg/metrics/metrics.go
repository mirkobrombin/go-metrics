package metrics

import "sync"

// Counter is the shared counter contract.
type Counter interface {
	Inc()
	Add(int64)
	Value() int64
}

// SimpleCounter is an in-memory counter useful for tests and local metrics.
type SimpleCounter struct {
	mu sync.Mutex
	v  int64
}

// NewCounter creates a new in-memory counter.
func NewCounter() *SimpleCounter {
	return &SimpleCounter{}
}

// Inc increments the counter by one.
func (c *SimpleCounter) Inc() {
	c.Add(1)
}

// Add increments the counter by the provided value.
func (c *SimpleCounter) Add(n int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v += n
}

// Value returns the current counter value.
func (c *SimpleCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}
