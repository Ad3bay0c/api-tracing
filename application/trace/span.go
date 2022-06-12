package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// NewSpan returns a new span from the global tracer. This can either be plain or customized
// depending on the custom argument.
func NewSpan(ctx context.Context, name string, custom SpanCustomiser) (context.Context, trace.Span) {
	if custom == nil {
		return otel.Tracer("").Start(ctx, name)
	}
	return otel.Tracer("").Start(ctx, name, custom.customise()...)
}

// SpanFromContext returns the current span from a context. Instead of creating child spans
// for each operation, this helps rely on the parent span.
func SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}

// AddSpanTags adds a new tags to the span. It will appear under "Tags" section
// of the selected span.
func AddSpanTags(span trace.Span, tags map[string]string) {
	attributes := make([]attribute.KeyValue, len(tags))

	var i int
	for k, v := range tags {
		attributes[i] = attribute.Key(k).String(v)
		i++
	}
	span.SetAttributes(attributes...)
}

// AddSpanEvents adds a new events to the span. It will appear under the "Logs"
// section of the selected span.
func AddSpanEvents(span trace.Span, name string, events map[string]string) {
	list := make([]trace.EventOption, len(events))

	var i int
	for k, v := range events {
		list[i] = trace.WithAttributes(attribute.Key(k).String(v))
		i++
	}

	span.AddEvent(name, list...)
}

// AddSpanError adds a new event to the span. It will appear under the "Logs"
// section of the selected span. This is not going to flag the span as "failed".
func AddSpanError(span trace.Span, err error) {
	span.RecordError(err)
}

// FailSpan flags the span as "failed" and adds "error" label on listed trace.
func FailSpan(span trace.Span, msg string) {
	span.SetStatus(codes.Error, msg)
}

// SpanCustomiser is used to enforce custom span options. Any custom concrete
// span customiser type must implement this interface.
type SpanCustomiser interface {
	customise() []trace.SpanStartOption
}
