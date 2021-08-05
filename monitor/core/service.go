package core

type Service interface {
	Id() string
	Name() string
	Description() string
}

type DefaultService struct {
	id          string
	name        string
	description string
}

func NewService(id string, name string, description string) Service {
	return &DefaultService{
		id:          id,
		name:        name,
		description: description,
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
