package handler

import (
	api "op/greetingapi"
	domain "op/greetingdomain"
	"op/internal/model"
)

type GreetingService struct {
	service *Service
}

func NewGreetingService(service *Service) *GreetingService {
	return &GreetingService{
		service: service,
	}
}

func (g *GreetingService) Greeting(id int) (*api.GreetingResponse, error) {
	mess, err := g.service.GreetingStorage.GetMessageByID(id)
	if err != nil {
		return nil, err
	}
	gtmess := toGreetingDomain(mess)
	return &api.GreetingResponse{ID: id, Data: gtmess}, nil
}

func (g *GreetingService) GreetingAll() (*api.GreetingAllResponse, error) {
	mess, err := g.service.GreetingStorage.GetMessages()
	if err != nil {
		return nil, err
	}
	messes := make([]*domain.GreetingMessage, len(mess))
	for i, m := range mess {
		messes[i] = toGreetingDomain(m)
	}
	return &api.GreetingAllResponse{Data: messes}, err
}

func toGreetingDomain(m *model.MessageData) *domain.GreetingMessage {
	return &domain.GreetingMessage{
		ID:      m.ID,
		Type:    uint8(m.Type),
		Message: m.Message,
	}
}

func (g *GreetingService) SaveMessage(
	req *api.SaveNewGreetingRequest,
) (*api.SaveNewGreetingResponse, error) {
	id, err := g.service.GreetingStorage.SaveMessage(
		req.Message,
		model.MessageType(req.Type),
	)
	if err != nil {
		return nil, err
	}
	return &api.SaveNewGreetingResponse{ID: id}, nil
}
