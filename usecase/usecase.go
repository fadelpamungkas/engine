package usecase

import (
	"context"

	"github.com/engine/models"
	"github.com/engine/repository"
)

type UsecaseI interface {
	Get(ctx context.Context) (models.Response, error)
}

type Usecase struct {
	repo repository.RepositoryI
}

func NewUsecase(repo repository.RepositoryI) *Usecase {
	return &Usecase{
		repo: repo,
	}
}

func (u *Usecase) Get(ctx context.Context) (models.Response, error) {
	return u.repo.Get(ctx)
}
