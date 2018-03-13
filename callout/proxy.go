package callout

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"golang.org/x/net/context"
	"os"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/sony/gobreaker"
	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/gokit"
	"golang.org/x/time/rate"
	"time"
)

type calloutProxy struct {
	requestCallout endpoint.Endpoint
}

type SendCalloutRequest struct {
	title   string
	message string
}

func NewCalloutProxy() Service {
	if getHalUrl() == "" {
		panic("No Alert Endpoint set. Please set the environment variable ALERT_ENDPOINT with the http address of the alert service")
	}
	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	return newProxy("", logger)

}

func NewKubernetesCalloutProxy(namespace string) Service {
	fieldKeys := []string{"method"}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = level.NewFilter(logger, level.AllowAll())
	logger = log.With(logger, "ts", log.DefaultTimestamp)

	service := newProxy(namespace, logger)
	service = NewLoggingService(log.With(logger, "component", "alert_proxy"), service)
	service = NewInstrumentService(kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "proxy",
		Subsystem: "callout_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys),
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "proxy",
			Subsystem: "alert_service",
			Name:      "error_count",
			Help:      "Number of errors.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "proxy",
			Subsystem: "callout_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys), service)

	return service
}

func newProxy(namespace string, logger log.Logger) Service {
	callout := makeCalloutHttpProxy(namespace, logger)
	callout = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{}))(callout)
	callout = ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 10))(callout)

	return &calloutProxy{requestCallout: callout}
}

func (s *calloutProxy) InvokeCallout(ctx context.Context, title, message string) error {
	_, err := s.requestCallout(ctx, &SendCalloutRequest{message: message, title: title})
	return err
}

func (calloutProxy) getFirstCall(ctx context.Context) (name string, number string, err error) {
	panic("Not implemented")
}

func getHalUrl() string {
	return os.Getenv("HAL_ENDPOINT")
}

func makeCalloutHttpProxy(namespace string, logger log.Logger) endpoint.Endpoint {
	return http.NewClient(
		"POST",
		alert.GetURL(namespace, "callout/"),
		gokit.EncodeJsonRequest,
		gokit.DecodeResponse,
		gokit.GetClientOpts(logger)...,
	).Endpoint()
}
