package main

import (
	"time"

	"context"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"
	"io"
	"log"
	"net/http"
)

var exp1 sdktrace.SpanExporter
var exp2 sdktrace.SpanExporter

func init() {
	var err error
	exp1, err = stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
		return
	}
	exp2, err = otlptracehttp.New(context.TODO(),
		otlptracehttp.WithEndpoint("otel-collector.k8s.sf"),
		otlptracehttp.WithInsecure(),
	)
	if err != nil {
		log.Panicf("failed to initialize stdouttrace exporter %v\n", err)
		return
	}

	r, _ := resource.New(context.TODO(),
		resource.WithOS(),
		resource.WithOSDescription(),
		resource.WithProcess(),
		resource.WithHost(),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(semconv.ServiceNameKey.String("golang")),
	)
	p := sdktrace.NewTracerProvider(
		sdktrace.WithResource(r),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(sdktrace.NewSimpleSpanProcessor(exp1)),
		sdktrace.WithSpanProcessor(sdktrace.NewSimpleSpanProcessor(exp2)),
	)
	otel.SetTracerProvider(p)
}
func demo() {
	tracer := otel.Tracer("main.go")
	ctx, span := tracer.Start(context.TODO(), "span1")
	time.Sleep(time.Second / 10)

	span.AddEvent("span1 start", trace.WithAttributes(attribute.String("abc", "def")))
	{
		var pp propagation.TraceContext
		pp.Extract(context.TODO(), propagation.HeaderCarrier(http.Header{}))
		_, span := tracer.Start(ctx, "span1.1")
		span.AddEvent("span1.1 start")
		rsp, err := otelhttp.Get(ctx, "http://qq.com")
		if err == nil {
			io.Copy(io.Discard, rsp.Body)
			span.SetStatus(codes.Ok, "")
		} else {
			span.RecordError(err, trace.WithStackTrace(true))
			log.Println(err)
		}
		span.End()
	}
	span.End()

	exp1.Shutdown(context.TODO())
	exp2.Shutdown(context.TODO())
}

func main() {
	demo()
}
