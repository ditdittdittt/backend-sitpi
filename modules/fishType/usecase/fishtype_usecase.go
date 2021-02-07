package usecase

import (
	"context"
	"time"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type fishTypeUsecase struct {
	fishTypeRepo   domain.FishTypeRepository
	contextTimeout time.Duration
}

func NewFishTypeUsecase(ft domain.FishTypeRepository, timeout time.Duration) domain.FishTypeUsecase {
	return &fishTypeUsecase{
		fishTypeRepo:   ft,
		contextTimeout: timeout,
	}
}

func (uc *fishTypeUsecase) Fetch(ctx context.Context) (res []domain.FishType, err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fishTypeRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *fishTypeUsecase) GetByID(ctx context.Context, id int64) (res domain.FishType, err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fishTypeRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *fishTypeUsecase) Update(ctx context.Context, ft *domain.FishType) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFishType, err := uc.fishTypeRepo.GetByID(ctx, ft.ID)

	if err != nil {
		return
	}
	if existedFishType == (domain.FishType{}) {
		return domain.ErrNotFound
	}

	ft.CreatedAt = existedFishType.CreatedAt
	ft.UpdatedAt = time.Now()

	err = uc.fishTypeRepo.Update(ctx, ft)
	return
}

func (uc *fishTypeUsecase) Store(ctx context.Context, ft *domain.FishType) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	ft.CreatedAt = time.Now()
	ft.UpdatedAt = time.Now()
	err = uc.fishTypeRepo.Store(ctx, ft)
	return
}

func (uc *fishTypeUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFishType, err := uc.fishTypeRepo.GetByID(ctx, id)

	if err != nil {
		return
	}
	if existedFishType == (domain.FishType{}) {
		return domain.ErrNotFound
	}

	err = uc.fishTypeRepo.Delete(ctx, id)
	return
}
