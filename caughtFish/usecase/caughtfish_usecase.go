package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type caughtFishUsecase struct {
	caughtFishRepo domain.CaughtFishRepository
	auctionRepo    domain.AuctionRepository
	contextTimeout time.Duration
}

func (uc *caughtFishUsecase) Fetch(ctx context.Context) (res []domain.CaughtFish, err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.caughtFishRepo.Fetch(ctx)
	if err != nil {
		return nil, err
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

	existedCaughtFish, err := uc.caughtFishRepo.GetByID(ctx, cf.ID)
	if err != nil {
		return
	}
	if existedCaughtFish == (domain.CaughtFish{}) {
		return domain.ErrNotFound
	}

	cf.TpiID = existedCaughtFish.TpiID
	cf.OfficerID = existedCaughtFish.OfficerID
	cf.CreatedAt = existedCaughtFish.CreatedAt
	cf.UpdatedAt = time.Now()

	err = uc.caughtFishRepo.Update(ctx, cf)
	return
}

func (uc *caughtFishUsecase) Store(ctx context.Context, cf *domain.CaughtFish, a *domain.Auction) (err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	cf.TpiID = 1
	cf.OfficerID = 1
	cf.CreatedAt = time.Now()
	cf.UpdatedAt = time.Now()

	lastID, err := uc.caughtFishRepo.Store(ctx, cf)
	if err != nil {
		return
	}

	a.TpiID = 1
	a.OfficerID = 1
	a.CaughtFishID = lastID
	a.Status = 1
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()

	err = uc.auctionRepo.Store(ctx, a)
	if err != nil {
		return
	}

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

func NewCaughtFishUsecase(c domain.CaughtFishRepository, a domain.AuctionRepository, timeout time.Duration) domain.CaughtFishUsecase {
	return &caughtFishUsecase{
		caughtFishRepo: c,
		auctionRepo:    a,
		contextTimeout: timeout,
	}
}
