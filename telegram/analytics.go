package telegram

import (
	"context"
	"github.com/go-kit/kit/metrics"
	"time"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	errorCount     metrics.Counter
	requestLatency metrics.Histogram
	Service
}

func NewInstrumentService(counter metrics.Counter, errorCount metrics.Counter, latency metrics.Histogram, s Service) Service {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		errorCount:     errorCount,
		Service:        s,
	}
}

func (s *instrumentingService) SendMessage(ctx context.Context, chatID int64, message string, messageID int) (msgid int, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SendMessage").Add(1)
		s.requestLatency.With("method", "SendMessage").Observe(time.Since(begin).Seconds())
		if err != nil {
			s.errorCount.With("method", "SendMessage").Add(1)
		}
	}(time.Now())
	_, err = s.Service.SendMessage(ctx, chatID, message, messageID)
	return
}
func (s *instrumentingService) SendMessagePlainText(ctx context.Context, chatID int64, message string, messageID int) (msgid int, err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SendMessagePlainText").Add(1)
		s.requestLatency.With("method", "SendMessagePlainText").Observe(time.Since(begin).Seconds())
		if err != nil {
			s.errorCount.With("method", "SendMessagePlainText").Add(1)
		}
	}(time.Now())
	_, err = s.Service.SendMessagePlainText(ctx, chatID, message, messageID)
	return
}
func (s *instrumentingService) SendImageToGroup(ctx context.Context, image []byte, group int64) (err error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SendImageToGroup").Add(1)
		s.requestLatency.With("method", "SendImageToGroup").Observe(time.Since(begin).Seconds())
		if err != nil {
			s.errorCount.With("method", "SendImageToGroup").Add(1)
		}
	}(time.Now())
	return s.Service.SendImageToGroup(ctx, image, group)
}
func (s *instrumentingService) SendKeyboard(ctx context.Context, buttons []string, text string, chat int64) (int, error) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "SendKeyboard").Add(1)
		s.requestLatency.With("method", "SendKeyboard").Observe(time.Since(begin).Seconds())
	}(time.Now())
	return s.Service.SendKeyboard(ctx, buttons, text, chat)
}
func (s *instrumentingService) RegisterCommand(command Command) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "RegisterCommand").Add(1)
		s.requestLatency.With("method", "RegisterCommand").Observe(time.Since(begin).Seconds())
	}(time.Now())
	s.Service.RegisterCommand(command)
}
func (s *instrumentingService) RegisterCommandLet(commandlet Commandlet) {
	defer func(begin time.Time) {
		s.requestCount.With("method", "RegisterCommandLet").Add(1)
		s.requestLatency.With("method", "RegisterCommandLet").Observe(time.Since(begin).Seconds())
	}(time.Now())
	s.Service.RegisterCommandLet(commandlet)
}
