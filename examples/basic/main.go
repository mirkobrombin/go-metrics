package main

import (
	"fmt"

	"github.com/mirkobrombin/go-metrics/pkg/metrics"
)

func main() {
	counter := metrics.NewCounter()
	counter.Inc()
	counter.Add(2)

	fmt.Printf("counter = %d\n", counter.Value())
}
