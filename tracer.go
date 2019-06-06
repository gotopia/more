package more

import (
	"fmt"
	"net/http"

	"github.com/gotopia/more/config"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/zipkin"
	"github.com/uber/jaeger-lib/metrics"
)

func init() {
	if config.Tracer.Enabled() {
		name := config.Tracer.Name()
		if name != "jaeger" {
			err := errors.New("only jaeger is supported now")
			panic(err)
		}

		localAgentHostPort := fmt.Sprintf("%v:%v", config.Tracer.Agent.Host(), config.Tracer.Agent.Port())

		var cfg jaegercfg.Configuration
		if config.Production() {
			cfg = jaegercfg.Configuration{
				Reporter: &jaegercfg.ReporterConfig{
					LocalAgentHostPort: localAgentHostPort,
				},
			}
		} else {
			cfg = jaegercfg.Configuration{
				Sampler: &jaegercfg.SamplerConfig{
					Type:  jaeger.SamplerTypeConst,
					Param: 1,
				},
				Reporter: &jaegercfg.ReporterConfig{
					LogSpans:           true,
					LocalAgentHostPort: localAgentHostPort,
				},
			}
		}

		jLogger := jaegerlog.StdLogger
		jMetricsFactory := metrics.NullFactory
		zipkinPropagator := zipkin.NewZipkinB3HTTPHeaderPropagator()

		_, err := cfg.InitGlobalTracer(
			config.Tracer.ServiceName(),
			jaegercfg.Logger(jLogger),
			jaegercfg.Metrics(jMetricsFactory),
			jaegercfg.Injector(opentracing.HTTPHeaders, zipkinPropagator),
			jaegercfg.Extractor(opentracing.HTTPHeaders, zipkinPropagator),
			jaegercfg.ZipkinSharedRPCSpan(true),
		)
		if err != nil {
			err = errors.Wrap(err, "could not initialize jaeger tracer")
			panic(err)
		}
	}
}

var grpcGatewayTag = opentracing.Tag{Key: string(ext.Component), Value: "grpc-gateway"}

func tracingWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		parentSpanContext, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(r.Header))
		if err == nil || err == opentracing.ErrSpanContextNotFound {
			serverSpan := opentracing.GlobalTracer().StartSpan(
				r.URL.String(),
				// this is magical, it attaches the new span to the parent parentSpanContext, and creates an unparented one if empty.
				ext.RPCServerOption(parentSpanContext),
				grpcGatewayTag,
			)
			r = r.WithContext(opentracing.ContextWithSpan(r.Context(), serverSpan))
			defer serverSpan.Finish()
		}
		h.ServeHTTP(w, r)
	})
}
