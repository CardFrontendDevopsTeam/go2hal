package alert

import (
	"github.com/weAutomateEverything/go2hal/telegram"
	"golang.org/x/net/context"
	"gopkg.in/kyokomi/emoji.v1"
	"log"
)

/*
Service interface
*/
type Service interface {
	SendAlert(ctx context.Context, message string) error
	SendNonTechnicalAlert(ctx context.Context, message string) error
	SendHeartbeatGroupAlert(ctx context.Context, message string) error
	SendImageToAlertGroup(ctx context.Context, image []byte) error
	SendImageToHeartbeatGroup(ctx context.Context, image []byte) error
	SendError(ctx context.Context, err error) error
	SendAlertKeyboardRecipe(ctx context.Context, buttons []string) error
	SendAlertEnvironment(ctx context.Context, buttons []string) error
	SendAlertNodes(ctx context.Context, nodes []string) error
}

type service struct {
	telegram telegram.Service
	store    Store
}

/*
NewService returns a new Alert Service
*/
func NewService(t telegram.Service, store Store) Service {

	return &service{
		telegram: t,
		store:    store,
	}
}

//IMPL

func (s *service) SendAlert(ctx context.Context, message string) error {
	alertGroup, err := s.store.alertGroup()
	if err != nil {
		return err
	}
	err = s.telegram.SendMessage(ctx, alertGroup, message, 0)
	return err
}

func (s *service) SendNonTechnicalAlert(ctx context.Context, message string) error {
	return nil
}

func (s *service) SendImageToAlertGroup(ctx context.Context, image []byte) error {

	alertGroup, err := s.store.alertGroup()
	if err != nil {
		s.SendError(ctx, err)
		return err
	}

	return s.telegram.SendImageToGroup(ctx, image, alertGroup)
}

func (s *service) SendImageToHeartbeatGroup(ctx context.Context, image []byte) error {
	group, err := s.store.heartbeatGroup()
	if err != nil {
		s.SendError(ctx, err)
		return err
	}

	return s.telegram.SendImageToGroup(ctx, image, group)

}

func (s *service) SendHeartbeatGroupAlert(ctx context.Context, message string) error {
	group, err := s.store.heartbeatGroup()
	if err != nil {
		s.SendError(ctx, err)
		return err
	}

	return s.telegram.SendMessage(ctx, group, message, 0)
}

func (s *service) SendError(ctx context.Context, err error) error {
	log.Println(err.Error())
	group, e := s.store.heartbeatGroup()
	if e != nil {
		return e
	}
	return s.telegram.SendMessagePlainText(ctx, group, emoji.Sprintf(":poop: %s", err.Error()), 0)
}
func (s *service) SendAlertKeyboardRecipe(ctx context.Context, buttons []string) error {
	alertGroup, err := s.store.alertGroup()
	s.telegram.SendKeyboard(ctx, buttons, "Please select the application", alertGroup)
	return err;
}
func (s *service) SendAlertEnvironment(ctx context.Context, nodes []string) error {
	alertGroup, err := s.store.alertGroup()
	s.telegram.SendKeyboard(ctx, nodes, "Please select the environment", alertGroup)
	return err;
}
func (s *service) SendAlertNodes(ctx context.Context, nodes []string) error {
	alertGroup, err := s.store.alertGroup()
	s.telegram.SendKeyboard(ctx, nodes, "please select the node", alertGroup)
	return err;
}
