package usecase

import (
	"github.com/rsmnarts/codebase-golang-backend/internal/domain"
	"time"

	"github.com/google/uuid"
)

// UserUseCase handles user business logic
type UserUseCase struct {
	userRepo domain.UserRepository
}

// NewUserUseCase creates a new user use case
func NewUserUseCase(userRepo domain.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

// CreateUser creates a new user
func (uc *UserUseCase) CreateUser(name, email string) (*domain.User, error) {
	if name == "" || email == "" {
		return nil, domain.ErrInvalidInput
	}

	user := &domain.User{
		ID:        uuid.New().String(),
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err := uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// GetUser retrieves a user by ID
func (uc *UserUseCase) GetUser(id string) (*domain.User, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	return uc.userRepo.GetByID(id)
}

// GetAllUsers retrieves all users
func (uc *UserUseCase) GetAllUsers() ([]*domain.User, error) {
	return uc.userRepo.GetAll()
}

// UpdateUser updates an existing user
func (uc *UserUseCase) UpdateUser(id, name, email string) (*domain.User, error) {
	if id == "" {
		return nil, domain.ErrInvalidInput
	}

	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if name != "" {
		user.Name = name
	}
	if email != "" {
		user.Email = email
	}
	user.UpdatedAt = time.Now()

	err = uc.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DeleteUser deletes a user by ID
func (uc *UserUseCase) DeleteUser(id string) error {
	if id == "" {
		return domain.ErrInvalidInput
	}

	return uc.userRepo.Delete(id)
}
