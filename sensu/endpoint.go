package sensu

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/weAutomateEverything/go2hal/gokit"
)

type SensuMessageRequest struct {
	Text        string            `json:"text"`
	IconURL     string            `json:"icon_url"`
	Attachments []sensuAttachment `json:"attachments"`
}

type sensuAttachment struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func makeSensuEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(SensuMessageRequest)
		s.handleSensu(ctx, gokit.GetChatId(ctx), req)
		return nil, nil
	}
}
