package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type auctionUsecase struct {
	auctionRepo    domain.AuctionRepository
	contextTimeout time.Duration
}

func (uc *auctionUsecase) Fetch(ctx context.Context) (res []domain.Auction, err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.auctionRepo.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *auctionUsecase) GetByID(ctx context.Context, id int64) (res domain.Auction, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.auctionRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *auctionUsecase) Update(ctx context.Context, a *domain.Auction) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedAuction, err := uc.auctionRepo.GetByID(ctx, a.ID)
	if err != nil {
		return
	}
	if existedAuction == (domain.Auction{}) {
		return domain.ErrNotFound
	}

	a.TpiID = existedAuction.TpiID
	a.OfficerID = existedAuction.OfficerID
	a.CaughtFishID = existedAuction.CaughtFishID
	a.Status = existedAuction.Status
	a.CreatedAt = existedAuction.CreatedAt
	a.UpdatedAt = time.Now()

	err = uc.auctionRepo.Update(ctx, a)
	return
}

func (uc *auctionUsecase) Store(ctx context.Context, a *domain.Auction) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	a.TpiID = 1
	a.OfficerID = 1
	a.Status = 1
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	err = uc.auctionRepo.Store(ctx, a)
	return
}

func (uc *auctionUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedAuction, err := uc.auctionRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedAuction == (domain.Auction{}) {
		return domain.ErrNotFound
	}

	err = uc.auctionRepo.Delete(ctx, id)
	return
}

func (uc *auctionUsecase) Inquiry(ctx context.Context) (res []domain.Auction, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.auctionRepo.Inquiry(ctx)
	if err != nil {
		return nil, err
	}

	return
}

func NewAuctionUsecase(a domain.AuctionRepository, timeout time.Duration) domain.AuctionUsecase {
	return &auctionUsecase{
		auctionRepo:    a,
		contextTimeout: timeout,
	}
}
