// Package callout provides a mechanism to invoke various different forms of callout depending on the services
// linked when creating the service.
package callout

import (
	"fmt"

	"github.com/weAutomateEverything/go2hal/alert"
	"github.com/weAutomateEverything/go2hal/firstCall"
	"github.com/weAutomateEverything/go2hal/halaws"
	"github.com/weAutomateEverything/go2hal/jira"
	"github.com/weAutomateEverything/go2hal/snmp"
	"golang.org/x/net/context"
)

//Service interface for the Callout Service
type Service interface {
	//InvokeCallout will invoke snmp if configured, then create a jira ticket if configured.
	InvokeCallout(ctx context.Context, title, message string, variables map[string]string) error
}

type service struct {
	alert     alert.Service
	snmp      snmp.Service
	jira      jira.Service
	alexa     halaws.Service
	firstcall firstCall.Service
}

// NewService creates a new Callout Service. Parameters can be passed in as Nil should they not be required.
// any items that are nil will simply not be invoked.
func NewService(alert alert.Service, firstcall firstCall.Service, snmp snmp.Service, jira jira.Service, alexa halaws.Service) Service {
	return &service{
		snmp:      snmp,
		alert:     alert,
		alexa:     alexa,
		firstcall: firstcall,
		jira:      jira,
	}
}

// InvokeCallout will invoke snmp if configured, then create a jira ticket if configured, finally it will invoke a phone
// call via alexa connect, if configured.
func (s *service) InvokeCallout(ctx context.Context, title, message string, variables map[string]string) error {
	s.alert.SendError(ctx, fmt.Errorf("invoking callout for: %s, %s", title, message))
	if s.snmp != nil {
		s.snmp.SendSNMPMessage(ctx)
	}
	name, phone, err := s.firstcall.GetFirstCall(ctx)
	if err != nil {
		s.alert.SendError(ctx, err)
		name = "DEFAULT"
	}
	if s.jira != nil {
		s.jira.CreateJira(ctx, title, message, name)
	}
	if s.alexa != nil {
		return s.alexa.SendAlert(ctx, phone, name, variables)
	}
	return nil
}
