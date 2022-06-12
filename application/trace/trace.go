package trace

import (
	"context"
	"io"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkresource "go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

type ProviderConfig struct {
	Url            string
	ServiceName    string
	ServiceVersion string
	Environment    string
	// set this to `true` if you want to disable tracing completely
	Disabled bool
}

// Provider represents the tracer provider. Depending on the `config.Disabled`
// parameter, it will either use a "live" provider or a "no operations" version.
// The "no operations" means, tracing will be globally disabled.
type Provider struct {
	provider trace.TracerProvider
}

func NewProvider(config ProviderConfig) (Provider, error) {
	if config.Disabled {
		return Provider{provider: trace.NewNoopTracerProvider()}, nil
	}
	// implement exporter to display the trace data
	//exp, err := jaeger.NewRawExporter(
	//	jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(config.JaegerEndpoint)),
	//)
	//if err != nil {
	//	return Provider{}, err
	//}
	exp, _ := newExporter(os.Stdout)
	prv := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(sdkresource.NewWithAttributes(
			"",
			semconv.ServiceNameKey.String(config.ServiceName),
			semconv.ServiceVersionKey.String(config.ServiceVersion),
			semconv.DeploymentEnvironmentKey.String(config.Environment),
		)),
	)
	otel.SetTracerProvider(prv)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)
	return Provider{provider: prv}, nil
}

func (p Provider) Close(ctx context.Context) error {
	if prv, ok := p.provider.(*sdktrace.TracerProvider); ok {
		return prv.Shutdown(ctx)
	}
	return nil
}

func newExporter(w io.Writer) (sdktrace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		// Use human-readable output.
		//stdouttrace.WithPrettyPrint(),
		// Do not print timestamps for the demo.
		//stdouttrace.WithoutTimestamps(),
	)
}
