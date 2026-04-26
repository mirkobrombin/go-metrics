package metrics

import (
	"sort"
	"sync"
	"time"
)

type Counter interface {
	Inc()
	Add(int64)
	Value() int64
}

type Gauge interface {
	Set(float64)
	Inc()
	Dec()
	Add(float64)
	Sub(float64)
	Value() float64
}

type Histogram interface {
	Observe(float64)
}

type Timer interface {
	Start() *Timing
	ObserveDuration()
}

type Timing struct {
	start time.Time
}

func (t *Timing) Stop() time.Duration {
	return time.Since(t.start)
}

type SimpleCounter struct {
	mu sync.Mutex
	v  int64
}

func NewCounter() *SimpleCounter { return &SimpleCounter{} }
func (c *SimpleCounter) Inc()    { c.Add(1) }
func (c *SimpleCounter) Add(n int64) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v += n
}
func (c *SimpleCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v
}

type SimpleGauge struct {
	mu sync.Mutex
	v  float64
}

func NewGauge() *SimpleGauge { return &SimpleGauge{} }
func (g *SimpleGauge) Set(v float64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.v = v
}
func (g *SimpleGauge) Inc()    { g.Add(1) }
func (g *SimpleGauge) Dec()    { g.Sub(1) }
func (g *SimpleGauge) Add(v float64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.v += v
}
func (g *SimpleGauge) Sub(v float64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.v -= v
}
func (g *SimpleGauge) Value() float64 {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.v
}

type SimpleHistogram struct {
	mu      sync.Mutex
	buckets []float64
	counts  []int64
	total   int64
	sum     float64
}

func NewHistogram(buckets []float64) *SimpleHistogram {
	sorted := append([]float64(nil), buckets...)
	sort.Float64s(sorted)
	return &SimpleHistogram{
		buckets: sorted,
		counts:  make([]int64, len(sorted)+1),
	}
}

func (h *SimpleHistogram) Observe(v float64) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.total++
	h.sum += v
	for i, b := range h.buckets {
		if v <= b {
			h.counts[i]++
			return
		}
	}
	h.counts[len(h.counts)-1]++
}

type SimpleTimer struct{}

func NewTimer() *SimpleTimer { return &SimpleTimer{} }

func (t *SimpleTimer) Start() *Timing {
	return &Timing{start: time.Now()}
}
