package user

import (
	domain "godo/internal/domain/user"
)

type UseCase struct {
	service *domain.Service
}

func NewUseCase(service *domain.Service) *UseCase {
	return &UseCase{service: service}
}

func (usecase *UseCase) RegisterUser(name, email string) (*domain.User, error) {
	return usecase.service.Register(name, email)
}
