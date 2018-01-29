package ssh

import (
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentService(counter metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		Service:        s,
	}
}

func (s *instrumentingService)ExecuteRemoteCommand(commandName, address string) error{
	defer func(begin time.Time) {
		s.requestCount.With("method", "ExecuteRemoteCommand").Add(1)
		s.requestLatency.With("method", "ExecuteRemoteCommand").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.ExecuteRemoteCommand(commandName,address)
}

