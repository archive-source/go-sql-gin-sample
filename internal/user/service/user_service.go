package service

import (
	"context"

	"github.com/core-go/core"
	"go-service/internal/user/model"
)

func NewUserService(repository core.Repository) UserService {
	return &UserUseCase{repository: repository}
}

type UserUseCase struct {
	repository core.Repository
}

func (s *UserUseCase) Load(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	ok, err := s.repository.Get(ctx, id, &user)
	if !ok {
		return nil, err
	} else {
		return &user, err
	}
}
func (s *UserUseCase) Create(ctx context.Context, user *model.User) (int64, error) {
	return s.repository.Insert(ctx, user)
}
func (s *UserUseCase) Update(ctx context.Context, user *model.User) (int64, error) {
	return s.repository.Update(ctx, user)
}
func (s *UserUseCase) Patch(ctx context.Context, user map[string]interface{}) (int64, error) {
	return s.repository.Patch(ctx, user)
}
func (s *UserUseCase) Delete(ctx context.Context, id string) (int64, error) {
	return s.repository.Delete(ctx, id)
}
