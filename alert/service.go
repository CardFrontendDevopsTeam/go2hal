package alert

import (
	"github.com/weAutomateEverything/go2hal/telegram"
	"golang.org/x/net/context"
	"gopkg.in/kyokomi/emoji.v1"
	"os"
	"strconv"
)

/*
Service interface
*/
type Service interface {
	SendAlert(ctx context.Context, chatId uint32, message string) error
	SendImageToAlertGroup(ctx context.Context, chatid uint32, image []byte) error

	SendError(ctx context.Context, err error) error
	SendErrorImage(ctx context.Context, image []byte) error
}

type service struct {
	telegram telegram.Service
	store    telegram.Store
}

/*
NewService returns a new Alert Service
*/
func NewService(t telegram.Service, store telegram.Store) Service {

	return &service{
		telegram: t,
		store:    store,
	}
}

//IMPL

func (s *service) SendAlert(ctx context.Context, chatid uint32, message string) error {
	group, err := s.store.GetRoomKey(chatid)
	if err != nil {
		return err
	}
	return s.telegram.SendMessage(ctx, group, message, 0)
}

func (s *service) SendImageToAlertGroup(ctx context.Context, chatid uint32, image []byte) error {

	group, err := s.store.GetRoomKey(chatid)
	if err != nil {
		return err
	}

	return s.telegram.SendImageToGroup(ctx, image, group)
}

func (s *service) SendError(ctx context.Context, err error) error {
	g := os.Getenv("ERROR_GROUP")
	group, errs := strconv.ParseInt(g, 10, 64)
	if errs != nil {
		return errs
	}
	return s.telegram.SendMessagePlainText(ctx, group, emoji.Sprintf(":poop: %s", err.Error()), 0)
}

func (s service) SendErrorImage(ctx context.Context, image []byte) error {
	g := os.Getenv("ERROR_GROUP")
	group, err := strconv.ParseUint(g, 10, 32)
	if err != nil {
		return err
	}
	return s.SendImageToAlertGroup(ctx, uint32(group), image)
}
