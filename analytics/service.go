package analytics

import (
	"fmt"
	"bytes"
	"strings"
	"github.com/zamedic/go2hal/alert"
	"encoding/json"
	"github.com/zamedic/go2hal/chef"
)

type Service interface {
	SendAnalyticsAlert(message string)
}

type service struct {
	alert     alert.Service
	chefStore chef.Store

}

func NewService(alertService alert.Service, chefStore chef.Store) Service{
	return &service{alertService,chefStore}
}

func (s *service) SendAnalyticsAlert(message string){
	if !s.checkSend(message) {
		return
	}
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(message), &dat); err != nil {
		s.alert.SendError(fmt.Errorf("Error unmarshalling: %s", message))
		return
	}

	attachments := dat["attachments"].([]interface{})
	//Loop though the attachmanets, there should be only 1

	var buffer bytes.Buffer
	buffer.WriteString("*analytics Event*\n")
	buffer.WriteString(dat["text"].(string))
	buffer.WriteString("\n")
	getfield(attachments, &buffer)

	s.alert.SendAlert(buffer.String())
}

func (s *service)checkSend(message string) bool {
	message = strings.ToUpper(message)
	recipes, err := s.chefStore.GetRecipes()
	if err != nil {
		s.alert.SendError(err)
		return false
	}
	for _, recipe := range recipes {
		check := "RECIPE["+strings.ToUpper(recipe.Recipe)+"]"
		if strings.Contains(message,check) {
			return true
		}
	}
	return false;
}

func getfield(attachments []interface{}, buffer *bytes.Buffer) {
	for _, attachment := range attachments {
		attachmentI := attachment.(map[string]interface{})
		fields := attachmentI["fields"].([]interface{})

		//Loop through the fields
		for _, field := range fields {
			fieldI := field.(map[string]interface{})
			buffer.WriteString("*")
			buffer.WriteString(fieldI["title"].(string))
			buffer.WriteString("* ")
			buffer.WriteString(fieldI["value"].(string))
			buffer.WriteString("\n")
		}
	}
}
