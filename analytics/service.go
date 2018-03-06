package analytics

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/chef"
	"github.com/weAutomateEverything/go2hal/util"
	"golang.org/x/net/context"
	"strings"
)

type Service interface {
	SendAnalyticsAlert(ctx context.Context, message string)
}

type service struct {
	alert     alert.Service
	chefStore chef.Store
}

func NewService(alertService alert.Service, chefStore chef.Store) Service {
	return &service{alertService, chefStore}
}

func (s *service) SendAnalyticsAlert(ctx context.Context, message string) {
	if !s.checkSend(ctx, message) {
		return
	}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(message), &dat); err != nil {
		s.alert.SendError(ctx, fmt.Errorf("Error unmarshalling: %s", message))
		return
	}

	attachments := dat["attachments"].([]interface{})
	//Loop though the attachmanets, there should be only 1

	var buffer bytes.Buffer
	buffer.WriteString("*analytics Event*\n")
	buffer.WriteString(dat["text"].(string))
	buffer.WriteString("\n")
	util.Getfield(attachments, &buffer)

	s.alert.SendAlert(ctx, buffer.String())
}

func (s *service) checkSend(ctx context.Context, message string) bool {
	message = strings.ToUpper(message)
	recipes, err := s.chefStore.GetRecipes()
	if err != nil {
		s.alert.SendError(ctx, err)
		return false
	}
	for _, recipe := range recipes {
		check := "RECIPE[" + strings.ToUpper(recipe.Recipe) + "]"
		if strings.Contains(message, check) {
			return true
		}
	}
	return false
}
