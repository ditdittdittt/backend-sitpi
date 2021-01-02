package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type fishingGearUsecase struct {
	fishingGearRepo domain.FishingGearRepository
	contextTimeout  time.Duration
}

func (uc *fishingGearUsecase) Fetch(ctx context.Context) (res []domain.FishingGear, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fishingGearRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *fishingGearUsecase) GetByID(ctx context.Context, id int64) (res domain.FishingGear, err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fishingGearRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *fishingGearUsecase) Update(ctx context.Context, fg *domain.FishingGear) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFishingGear, err := uc.fishingGearRepo.GetByID(ctx, fg.ID)

	if err != nil {
		return
	}
	if existedFishingGear == (domain.FishingGear{}) {
		return domain.ErrNotFound
	}

	fg.CreatedAt = existedFishingGear.CreatedAt
	fg.UpdatedAt = time.Now()

	err = uc.fishingGearRepo.Update(ctx, fg)
	return
}

func (uc *fishingGearUsecase) Store(ctx context.Context, fg *domain.FishingGear) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	fg.CreatedAt = time.Now()
	fg.UpdatedAt = time.Now()
	err = uc.fishingGearRepo.Store(ctx, fg)
	return
}

func (uc *fishingGearUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFishingGear, err := uc.fishingGearRepo.GetByID(ctx, id)

	if err != nil {
		return
	}
	if existedFishingGear == (domain.FishingGear{}) {
		return domain.ErrNotFound
	}

	err = uc.fishingGearRepo.Delete(ctx, id)
	return
}

func NewFishingGearUsecase(fg domain.FishingGearRepository, timeout time.Duration) domain.FishingGearUsecase {
	return &fishingGearUsecase{
		fishingGearRepo: fg,
		contextTimeout:  timeout,
	}
}
