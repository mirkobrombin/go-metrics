package metrics_test

import (
	"sync"
	"testing"

	"github.com/mirkobrombin/go-metrics/pkg/metrics"
)

func TestCounterTracksValues(t *testing.T) {
	counter := metrics.NewCounter()
	counter.Inc()
	counter.Add(4)

	if counter.Value() != 5 {
		t.Fatalf("Value() = %d, want %d", counter.Value(), 5)
	}
}

func TestCounterIsConcurrencySafe(t *testing.T) {
	counter := metrics.NewCounter()

	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			counter.Inc()
		}()
	}
	wg.Wait()

	if counter.Value() != 32 {
		t.Fatalf("Value() = %d, want %d", counter.Value(), 32)
	}
}
