package usecase

import (
	"context"

	"github.com/engine/models"
	"go.uber.org/zap"
)

type RepositoryI interface {
	Get(ctx context.Context) (models.Response, error)
	GetForEachRowGoroutines(ctx context.Context) (models.Response, error)
	GetForEachRow(ctx context.Context) (models.Response, error)
	GetRow(ctx context.Context) (models.Response, error)
	Insert(ctx context.Context, req models.AccountRequest) (models.Response, error)
	Update(ctx context.Context) (models.Response, error)
	Delete(ctx context.Context) (models.Response, error)
}

type usecase struct {
	repo RepositoryI
	log  *zap.Logger
}

func NewUsecase(repo RepositoryI, log *zap.Logger) *usecase {
	return &usecase{
		repo: repo,
		log:  log,
	}
}

func (u *usecase) Get(ctx context.Context) (models.Response, error) {
	return u.repo.Get(ctx)
}

func (u *usecase) GetForEachRowGoroutines(ctx context.Context) (models.Response, error) {
	return u.repo.GetForEachRowGoroutines(ctx)
}

func (u *usecase) GetForEachRow(ctx context.Context) (models.Response, error) {
	return u.repo.GetForEachRow(ctx)
}

func (u *usecase) GetRow(ctx context.Context) (models.Response, error) {
	return u.repo.GetRow(ctx)
}

func (u *usecase) Insert(ctx context.Context, req models.AccountRequest) (models.Response, error) {
	return u.repo.Insert(ctx, req)
}

func (u *usecase) Update(ctx context.Context) (models.Response, error) {
	return u.repo.Update(ctx)
}

func (u *usecase) Delete(ctx context.Context) (models.Response, error) {
	return u.repo.Delete(ctx)
}
