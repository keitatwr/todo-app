package usecases

import (
	"context"
	"time"

	"github.com/keitatwr/todo-app/domain"
)

type SignupUsecase struct {
	userRepository domain.UserRepository
	contextTimeout time.Duration
}

func NewSignupUsecase(ur domain.UserRepository, timeout time.Duration) domain.SignupUsecase {
	return &SignupUsecase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}

func (su *SignupUsecase) Create(ctx context.Context, name, email, password string) error {
	user := &domain.User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
	}
	err := su.userRepository.Create(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (su *SignupUsecase) GetUserByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := su.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}