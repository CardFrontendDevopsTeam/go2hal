package appdynamics

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type AddAppdynamicsQueueEndpointRequest struct {
	Name        string
	Application string
	Metricpath  string
}

type AddAppdynamicsEndpointRequest struct {
	Endpoint string
}

type ExecuteAppDynamicsCommandRequest struct {
	CommandName, NodeID, ApplicationID string
}

type BusinessAlertRequest struct {
	Severity, Type, DisplayName, SummaryMessage string
}

func makeAppDynamicsAlertEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(string)
		s.sendAppdynamicsAlert(ctx, ctx.Value("CHAT-ID").(uint32), req)
		return nil, nil
	}
}

func makeAddAppdynamicsEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddAppdynamicsEndpointRequest)
		return nil, s.addAppdynamicsEndpoint(req.Endpoint)

	}
}
func makeAddAppdynamicsQueueEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(AddAppdynamicsQueueEndpointRequest)
		return nil, s.addAppDynamicsQueue(ctx.Value("CHAT-ID").(uint32), req.Name, req.Application, req.Metricpath)
	}
}
func makExecuteCommandFromAppdynamics(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ExecuteAppDynamicsCommandRequest)
		return nil, s.executeCommandFromAppd(ctx, ctx.Value("CHAT-ID").(uint32), req.CommandName, req.ApplicationID, req.NodeID)

	}
}
