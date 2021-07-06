package gormotel

import (
oteltrace "go.opentelemetry.io/otel/trace"
)

type config struct {
	dbName         string
	tracerProvider oteltrace.TracerProvider
}

// Option is used to configure the client.
type Option func(*config)

// WithTracerProvider specifies a tracer provider to use for creating a tracer.
// If none is specified, the global provider is used.
func WithTracerProvider(provider oteltrace.TracerProvider) Option {
	return func(cfg *config) {
		cfg.tracerProvider = provider
	}
}

// WithDBName specified the database name to be used in span names
// since its not possible to extract this information from gorm
func WithDBName(name string) Option {
	return func(cfg *config) {
		cfg.dbName = name
	}
}
