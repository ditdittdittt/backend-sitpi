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

type caughtFishUsecase struct {
	caughtFishRepo domain.CaughtFishRepository
	auctionRepo    domain.AuctionRepository
	contextTimeout time.Duration
}

func (uc *caughtFishUsecase) GetTotalFisher(ctx context.Context, from string, to string) (totalFisher int, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	timestampFrom, err := time.Parse(layoutISO, from)
	if err != nil {
		return 0, err
	}

	timestampTo, err := time.Parse(layoutISO, to)
	if err != nil {
		return 0, err
	}
	timestampTo = timestampTo.Add(24 * time.Hour)

	caughtFish, err := uc.caughtFishRepo.GetTotalFisher(ctx, timestampFrom, timestampTo)
	if err != nil {
		return 0, err
	}

	totalFisher = caughtFish.TotalFisher
	return
}

func (uc *caughtFishUsecase) GetTotalProduction(ctx context.Context, from string, to string) (totalProduction float64, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	timestampFrom, err := time.Parse(layoutISO, from)
	if err != nil {
		return 0, err
	}

	timestampTo, err := time.Parse(layoutISO, to)
	if err != nil {
		return 0, err
	}
	timestampTo = timestampTo.Add(24 * time.Hour)

	caughtFish, err := uc.caughtFishRepo.GetTotalProduction(ctx, timestampFrom, timestampTo)
	if err != nil {
		return 0, err
	}

	totalProduction = caughtFish.TotalProduction
	return
}

func (uc *caughtFishUsecase) Fetch(ctx context.Context, request *domain.FetchCaughtFishRequest) (res []domain.CaughtFish, err error) {
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

	fisherID, err := strconv.ParseInt(request.FisherID, 10, 64)
	fishTypeID, err := strconv.ParseInt(request.FishTypeID, 10, 64)

	if err != nil {
		return []domain.CaughtFish{}, err
	}

	res, err = uc.caughtFishRepo.Fetch(ctx, timestampFrom, timestampTo, fisherID, fishTypeID)
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

func (uc *caughtFishUsecase) Update(ctx context.Context, id int64, request *domain.UpdateCaughtFishRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedCaughtFish, err := uc.caughtFishRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedCaughtFish == (domain.CaughtFish{}) {
		return domain.ErrNotFound
	}

	caughtFish := &domain.CaughtFish{
		ID:            id,
		UserID:        1,
		TpiID:         1,
		FisherID:      request.FisherID,
		FishTypeID:    request.FishTypeID,
		FishingGearID: request.FishingGearID,
		FishingAreaID: request.FishingAreaID,
		Weight:        request.Weight,
		TripDay:       request.TripDay,
		CreatedAt:     existedCaughtFish.CreatedAt,
		UpdatedAt:     time.Now(),
	}

	err = uc.caughtFishRepo.Update(ctx, caughtFish)
	return
}

func (uc *caughtFishUsecase) Store(ctx context.Context, request *domain.StoreCaughtFishRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	cf := &domain.CaughtFish{
		UserID:        1, // TODO change with auth
		TpiID:         1, // TODO change with auth
		FisherID:      request.FisherID,
		TripDay:       request.TripDay,
		FishingGearID: request.FishingGearID,
		FishingAreaID: request.FishingAreaID,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}

	a := &domain.Auction{
		TpiID:     1, // TODO change with auth
		StatusID:  1,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	for _, caughtFishData := range request.CaughtFishData {
		cf.FishTypeID = caughtFishData.FishTypeID
		cf.Weight = caughtFishData.Weight
		cf.WeightUnit = caughtFishData.WeightUnit

		lastID, err := uc.caughtFishRepo.Store(ctx, cf)
		if err != nil {
			return err
		}

		a.CaughtFishID = lastID

		err = uc.auctionRepo.Store(ctx, a)
		if err != nil {
			return err
		}
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
