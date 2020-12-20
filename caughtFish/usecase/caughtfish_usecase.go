package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type caughtFishUsecase struct {
	caughtFishRepo domain.CaughtFishRepository
	contextTimeout time.Duration
}

func (uc *caughtFishUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.CaughtFish, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, nextCursor, err = uc.caughtFishRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
	}

	return
}

func (uc *caughtFishUsecase) GetByID(ctx context.Context, id int64) (res domain.CaughtFish, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.caughtFishRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *caughtFishUsecase) Update(ctx context.Context, cf *domain.CaughtFish) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	cf.UpdatedAt = time.Now()
	err = uc.caughtFishRepo.Update(ctx, cf)
	return
}

func (uc *caughtFishUsecase) Store(ctx context.Context, cf *domain.CaughtFish) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err = uc.caughtFishRepo.Store(ctx, cf)
	return
}

func (uc *caughtFishUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedCaughtFish, err := uc.caughtFishRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedCaughtFish == (domain.CaughtFish{}) {
		return domain.ErrNotFound
	}

	err = uc.caughtFishRepo.Delete(ctx, id)
	return
}

func NewCaughtFishUsecase(c domain.CaughtFishRepository, timeout time.Duration) domain.CaughtFishUsecase {
	return &caughtFishUsecase{
		caughtFishRepo: c,
		contextTimeout: timeout,
	}
}
