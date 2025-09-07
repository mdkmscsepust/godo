package user

import "time"

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Register(name, emailStr string) (*User, error) {
	email, err := NewEmail(emailStr)
	if err != nil {
		return nil, err
	}

	user := &User{
		Id:        "sw",
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}

	err = s.repo.Add(user)
	return user, err
}
