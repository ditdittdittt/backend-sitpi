package usecase

import (
	"context"
	"time"

	"github.com/ditdittdittt/backend-sitpi/domain"
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

func (uc *buyerUsecase) Update(ctx context.Context, id int64, request *domain.UpdateBuyerRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedBuyer, err := uc.buyerRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedBuyer == (domain.Buyer{}) {
		return domain.ErrNotFound
	}

	buyer := &domain.Buyer{
		ID:          id,
		UserID:      1,
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		Status:      request.Status,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   existedBuyer.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	err = uc.buyerRepo.Update(ctx, buyer)
	return
}

func (uc *buyerUsecase) Store(ctx context.Context, request *domain.StoreBuyerRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	buyer := &domain.Buyer{
		UserID:      1,
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		Status:      request.Status,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = uc.buyerRepo.Store(ctx, buyer)
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

func (uc *buyerUsecase) Inquiry(ctx context.Context) (res []domain.Buyer, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.buyerRepo.Inquiry(ctx)
	if err != nil {
		return
	}

	return
}

func NewBuyerUsecase(b domain.BuyerRepository, timeout time.Duration) domain.BuyerUsecase {
	return &buyerUsecase{
		buyerRepo:      b,
		contextTimeout: timeout,
	}
}
