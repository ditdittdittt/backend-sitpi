package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type fisherUsecase struct {
	fisherRepo     domain.FisherRepository
	contextTimeout time.Duration
}

func (uc *fisherUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Fisher, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, nextCursor, err = uc.fisherRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (uc *fisherUsecase) GetByID(ctx context.Context, id int64) (res domain.Fisher, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fisherRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *fisherUsecase) Update(ctx context.Context, f *domain.Fisher) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	f.UpdatedAt = time.Now()
	err = uc.fisherRepo.Update(ctx, f)
	return
}

func (uc *fisherUsecase) Store(ctx context.Context, f *domain.Fisher) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err = uc.fisherRepo.Store(ctx, f)
	return
}

func (uc *fisherUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFisher, err := uc.fisherRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedFisher == (domain.Fisher{}) {
		return domain.ErrNotFound
	}

	err = uc.fisherRepo.Delete(ctx, id)
	return
}

func NewFisherUsecase(f domain.FisherRepository, timeout time.Duration) domain.FisherUsecase {
	return &fisherUsecase{
		fisherRepo:     f,
		contextTimeout: timeout,
	}
}
