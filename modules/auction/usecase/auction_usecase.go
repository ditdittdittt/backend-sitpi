package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/ditdittdittt/backend-sitpi/domain"
)

const (
	layoutISO = "2006-01-02"
)

type auctionUsecase struct {
	auctionRepo    domain.AuctionRepository
	contextTimeout time.Duration
}

func (uc *auctionUsecase) Fetch(ctx context.Context, request *domain.FetchAuctionRequest) (res []domain.Auction, err error) {
	var timestampFrom time.Time
	var timestampTo time.Time

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if request.From == "" {
		dateNowString := time.Now().Format("2006-01-02")
		timestampFrom, err = time.Parse(layoutISO, dateNowString)
	} else {
		timestampFrom, err = time.Parse(layoutISO, request.From)
	}

	if request.To == "" {
		timestampTo = time.Now()
	} else {
		timestampTo, err = time.Parse(layoutISO, request.To)
		timestampTo = timestampTo.Add(24 * time.Hour)
	}

	auctionID, err := strconv.ParseInt(request.AuctionID, 10, 64)
	fisherID, err := strconv.ParseInt(request.FisherID, 10, 64)
	fishTypeID, err := strconv.ParseInt(request.FishTypeID, 10, 64)
	statusID, err := strconv.ParseInt(request.StatusID, 10, 64)

	if err != nil {
		return []domain.Auction{}, err
	}

	res, err = uc.auctionRepo.Fetch(ctx, timestampFrom, timestampTo, auctionID, fisherID, fishTypeID, statusID)
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
	var timestampFrom time.Time
	var timestampTo time.Time

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	dateNowString := time.Now().Format("2006-01-02")
	timestampFrom, err = time.Parse(layoutISO, dateNowString)

	timestampTo = time.Now()

	res, err = uc.auctionRepo.Inquiry(ctx, timestampFrom, timestampTo)
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
