package more

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gotopia/more/config"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-client-go/zipkin"
	"github.com/uber/jaeger-lib/metrics"
	"google.golang.org/grpc/metadata"
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

func injectHeadersIntoMetadata(ctx context.Context, req *http.Request) metadata.MD {
	otHeaders := []string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
		"uber-trace-id",
	}
	pairs := []string{}
	for _, h := range otHeaders {
		if v := req.Header.Get(h); len(v) > 0 {
			pairs = append(pairs, h, v)
		}
	}
	return metadata.Pairs(pairs...)
}
