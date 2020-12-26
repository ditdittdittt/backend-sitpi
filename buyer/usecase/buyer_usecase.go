package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type buyerUsecase struct {
	buyerRepo      domain.BuyerRepository
	contextTimeout time.Duration
}

func (uc *buyerUsecase) Fetch(ctx context.Context) (res []domain.Buyer, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.buyerRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *buyerUsecase) GetByID(ctx context.Context, id int64) (res domain.Buyer, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.buyerRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *buyerUsecase) Update(ctx context.Context, b *domain.Buyer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	b.UpdatedAt = time.Now()
	err = uc.buyerRepo.Update(ctx, b)
	return
}

func (uc *buyerUsecase) Store(ctx context.Context, b *domain.Buyer) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err = uc.buyerRepo.Store(ctx, b)
	return
}

func (uc *buyerUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedBuyer, err := uc.buyerRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedBuyer == (domain.Buyer{}) {
		return domain.ErrNotFound
	}

	err = uc.buyerRepo.Delete(ctx, id)
	return
}

func NewBuyerUsecase(b domain.BuyerRepository, timeout time.Duration) domain.BuyerUsecase {
	return &buyerUsecase{
		buyerRepo:      b,
		contextTimeout: timeout,
	}
}
