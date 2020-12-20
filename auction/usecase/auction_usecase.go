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

func (uc *auctionUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Auction, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, nextCursor, err = uc.auctionRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
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

	a.UpdatedAt = time.Now()
	err = uc.auctionRepo.Update(ctx, a)
	return
}

func (uc *auctionUsecase) Store(ctx context.Context, a *domain.Auction) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

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

func NewAuctionUsecase(a domain.AuctionRepository, timeout time.Duration) domain.AuctionUsecase {
	return &auctionUsecase{
		auctionRepo:    a,
		contextTimeout: timeout,
	}
}
