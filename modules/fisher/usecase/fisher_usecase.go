package usecase

import (
	"context"
	"time"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

type fisherUsecase struct {
	fisherRepo     domain.FisherRepository
	contextTimeout time.Duration
}

func (uc *fisherUsecase) Fetch(ctx context.Context) (res []domain.Fisher, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fisherRepo.Fetch(ctx)
	if err != nil {
		return nil, err
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

func (uc *fisherUsecase) Update(ctx context.Context, id int64, request *domain.UpdateFisherRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedFisher, err := uc.fisherRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedFisher == (domain.Fisher{}) {
		return domain.ErrNotFound
	}

	fisher := &domain.Fisher{
		ID:          id,
		UserID:      1,
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		ShipType:    request.ShipType,
		AbkTotal:    request.AbkTotal,
		Status:      request.Status,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   existedFisher.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	err = uc.fisherRepo.Update(ctx, fisher)
	return
}

func (uc *fisherUsecase) Store(ctx context.Context, request *domain.StoreFisherRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	fisher := &domain.Fisher{
		UserID:      1,
		Nik:         request.Nik,
		Name:        request.Name,
		Address:     request.Address,
		ShipType:    request.ShipType,
		AbkTotal:    request.AbkTotal,
		Status:      request.Status,
		PhoneNumber: request.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = uc.fisherRepo.Store(ctx, fisher)
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

func (uc *fisherUsecase) Inquiry(ctx context.Context) (res []domain.Fisher, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.fisherRepo.Inquiry(ctx)

	if err != nil {
		return nil, err
	}

	return
}

func NewFisherUsecase(f domain.FisherRepository, timeout time.Duration) domain.FisherUsecase {
	return &fisherUsecase{
		fisherRepo:     f,
		contextTimeout: timeout,
	}
}
