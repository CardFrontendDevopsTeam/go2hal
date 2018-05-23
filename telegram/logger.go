package telegram

import (
	"context"
	"github.com/go-kit/kit/log"
	"time"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) SendMessage(ctx context.Context, chatID int64, message string, messageID int) (msgid int, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendMessage",
			"chatId", chatID,
			"message", message,
			"messageId", messageID,
			"msgid", msgid,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendMessage(ctx, chatID, message, messageID)
}
func (s *loggingService) SendMessagePlainText(ctx context.Context, chatID int64, message string, messageID int) (msgid int, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendMessagePlainText",
			"chatId", chatID,
			"message", message,
			"messageId", messageID,
			"msgid", msgid,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendMessagePlainText(ctx, chatID, message, messageID)
}
func (s *loggingService) SendImageToGroup(ctx context.Context, image []byte, group int64) (err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendImageToGroup",
			"imageBytes", len(image),
			"groupId", group,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendImageToGroup(ctx, image, group)
}
func (s *loggingService) SendKeyboard(ctx context.Context, buttons []string, text string, chat int64) (msgid int, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "SendKeyboard",
			"buttons", buttons,
			"text", text,
			"chat", chat,
			"msgid", msgid,
			"error", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return s.Service.SendKeyboard(ctx, buttons, text, chat)
}
func (s *loggingService) RegisterCommand(command Command) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "RegisterCommand",
			"command", command,
			"took", time.Since(begin),
		)
	}(time.Now())
	s.Service.RegisterCommand(command)
}
func (s *loggingService) RegisterCommandLet(commandlet Commandlet) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "RegisterCommandLet",
			"commandlet", commandlet,
			"took", time.Since(begin),
		)
	}(time.Now())
	s.Service.RegisterCommandLet(commandlet)
}
