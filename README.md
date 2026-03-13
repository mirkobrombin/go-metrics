# Go Metrics

A tiny **counter abstraction** for Go applications that want a simple in-memory implementation and a clean adapter surface.

## Features

- **Small Interface:** Model metrics with a three-method `Counter` contract.
- **In-Memory Counter:** Use `SimpleCounter` for tests, local tools, and lightweight services.
- **Adapter Friendly:** Wrap Prometheus, OpenTelemetry, or custom backends behind the same API.
- **Concurrency Safe:** Protect counter updates with internal locking.

## Installation

```bash
go get github.com/mirkobrombin/go-metrics
```

## Quick Start

```go
package main

import (
    "fmt"

    "github.com/mirkobrombin/go-metrics/pkg/metrics"
)

func main() {
    counter := metrics.NewCounter()
    counter.Inc()
    counter.Add(4)

    fmt.Println(counter.Value())
}
```

## Documentation

- [Getting Started](docs/getting-started.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
