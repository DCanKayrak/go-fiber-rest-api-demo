package query

import (
	"context"
	"errors"
	"golang-rest-api-demo/internal/application/repository"
	"golang-rest-api-demo/internal/domain"
)

type IUserQueryService interface {
	GetById(ctx context.Context, id string) (*domain.User, error)
	Get(ctx context.Context) ([]*domain.User, error)
}

type userQueryService struct {
	userRepository repository.IUserRepository
}

func NewUserQueryService(userRepository repository.IUserRepository) IUserQueryService {
	return &userQueryService{
		userRepository: userRepository,
	}
}

func (q *userQueryService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := q.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (q *userQueryService) Get(ctx context.Context) ([]*domain.User, error) {
	users, err := q.userRepository.Get(ctx)
	if err != nil {
		return nil, err
	}

	if users == nil {
		return nil, errors.New("users not found")
	}

	return users, nil
}
