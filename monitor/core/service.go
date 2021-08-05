package core

type Service interface {
	// Id is the globally unique identifier of the serivce
	Id() string
	// Name is a human-readable name of the service
	Name() string
	// Description is a human-readable description of the service
	Description() string
	// Origin is the node from which the service originates
	Origin() string
}

type DefaultService struct {
	id          string
	name        string
	description string
	origin      string
}

func NewService(id string, name string, description string, origin string) Service {
	return &DefaultService{
		id:          id,
		name:        name,
		description: description,
		origin:      origin,
	}
}

func (service *DefaultService) Id() string {
	return service.id
}

func (service *DefaultService) Name() string {
	return service.name
}

func (service *DefaultService) Description() string {
	return service.description
}

func (service *DefaultService) Origin() string {
	return service.origin
}
