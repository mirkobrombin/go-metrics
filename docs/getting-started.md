# Getting Started

`go-metrics` keeps the metrics contract intentionally small so application code can remain decoupled from any specific metrics backend.

## Included pieces

- `metrics.Counter` defines the contract.
- `metrics.SimpleCounter` provides an in-memory implementation.
- `metrics.NewCounter` returns a ready-to-use counter value.

## Integrating with other systems

The `Counter` interface is deliberately simple, making it straightforward to implement thin adapters for Prometheus or OpenTelemetry without leaking those dependencies into the rest of your application.
