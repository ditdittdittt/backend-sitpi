package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type weightUnitUsecase struct {
	weightUnitRepo domain.WeightUnitRepository
	contextTimeout time.Duration
}

func (uc *weightUnitUsecase) Fetch(ctx context.Context) (res []domain.WeightUnit, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.weightUnitRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *weightUnitUsecase) GetByID(ctx context.Context, id int64) (res domain.WeightUnit, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.weightUnitRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return

}

func (uc *weightUnitUsecase) Update(ctx context.Context, wu *domain.WeightUnit) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedWeightUnit, err := uc.weightUnitRepo.GetByID(ctx, wu.ID)

	if err != nil {
		return
	}
	if existedWeightUnit == (domain.WeightUnit{}) {
		return domain.ErrNotFound
	}

	wu.CreatedAt = existedWeightUnit.CreatedAt
	wu.UpdatedAt = time.Now()

	err = uc.weightUnitRepo.Update(ctx, wu)
	return
}

func (uc *weightUnitUsecase) Store(ctx context.Context, wu *domain.WeightUnit) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	wu.CreatedAt = time.Now()
	wu.UpdatedAt = time.Now()
	err = uc.weightUnitRepo.Store(ctx, wu)
	return
}

func (uc *weightUnitUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedWeightUnit, err := uc.weightUnitRepo.GetByID(ctx, id)

	if err != nil {
		return
	}
	if existedWeightUnit == (domain.WeightUnit{}) {
		return domain.ErrNotFound
	}

	err = uc.weightUnitRepo.Delete(ctx, id)
	return

}

func NewWeightUnitUsecase(wu domain.WeightUnitRepository, timeout time.Duration) domain.WeightUnitUsecase {
	return &weightUnitUsecase{
		weightUnitRepo: wu,
		contextTimeout: timeout,
	}
}
