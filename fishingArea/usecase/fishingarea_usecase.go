package usecase

import (
	"context"
	"time"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type fishingAreaUsecase struct {
	fishingAreaRepo domain.FishingAreaRepository
	contextTimeout  time.Duration
}

func (uc *fishingAreaUsecase) Fetch(ctx context.Context) (res []domain.FishingArea, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fishingAreaRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *fishingAreaUsecase) GetByID(ctx context.Context, id int64) (res domain.FishingArea, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fishingAreaRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *fishingAreaUsecase) Store(ctx context.Context, request *domain.StoreFishingAreaRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	fishingArea := &domain.FishingArea{
		DistrictID:          1,
		Name:                request.Name,
		SouthLatitudeDegree: request.SouthLatitudeDegree,
		SouthLatitudeMinute: request.SouthLatitudeMinute,
		SouthLatitudeSecond: request.SouthLatitudeSecond,
		EastLongitudeDegree: request.EastLongitudeDegree,
		EastLongitudeMinute: request.EastLongitudeMinute,
		EastLongitudeSecond: request.EastLongitudeSecond,
		CreatedAt:           time.Now(),
		UpdatedAt:           time.Now(),
	}

	err = uc.fishingAreaRepo.Store(ctx, fishingArea)
	return
}

func (uc *fishingAreaUsecase) Update(ctx context.Context, id int64, request *domain.UpdateFishingAreaRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFishingArea, err := uc.fishingAreaRepo.GetByID(ctx, id)

	if err != nil {
		return
	}
	if existedFishingArea == (domain.FishingArea{}) {
		return domain.ErrNotFound
	}

	fishingArea := &domain.FishingArea{
		ID:                  id,
		DistrictID:          1,
		Name:                request.Name,
		SouthLatitudeDegree: request.SouthLatitudeDegree,
		SouthLatitudeMinute: request.SouthLatitudeMinute,
		SouthLatitudeSecond: request.SouthLatitudeSecond,
		EastLongitudeDegree: request.EastLongitudeDegree,
		EastLongitudeMinute: request.EastLongitudeMinute,
		EastLongitudeSecond: request.EastLongitudeSecond,
		CreatedAt:           existedFishingArea.CreatedAt,
		UpdatedAt:           time.Now(),
	}

	err = uc.fishingAreaRepo.Update(ctx, fishingArea)
	return
}

func (uc *fishingAreaUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFishingArea, err := uc.fishingAreaRepo.GetByID(ctx, id)

	if err != nil {
		return
	}
	if existedFishingArea == (domain.FishingArea{}) {
		return domain.ErrNotFound
	}

	err = uc.fishingAreaRepo.Delete(ctx, id)
	return
}

func NewFishingAreaUsecase(fa domain.FishingAreaRepository, timeout time.Duration) domain.FishingAreaUsecase {
	return &fishingAreaUsecase{
		fishingAreaRepo: fa,
		contextTimeout:  timeout,
	}
}
