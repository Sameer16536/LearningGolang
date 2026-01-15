// service/user_service.go
package service

import (
	"fmt"
	domain "learninggolang/interfaces/domain"
)

// UserService contains business logic.
// It depends ONLY on an interface, not a concrete DB.
type UserService struct {
	repo domain.UserRepository
}

// NewUserService injects dependencies (dependency injection)
func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

// CreateUser handles validation + business rules
func (s *UserService) CreateUser(name string) error {
	if name == "" {
		return fmt.Errorf("name cannot be empty")
	}

	user := domain.User{Name: name}
	return s.repo.Create(user)
}
