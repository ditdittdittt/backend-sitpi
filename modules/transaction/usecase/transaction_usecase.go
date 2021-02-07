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

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	auctionRepo     domain.AuctionRepository
	contextTimeout  time.Duration
}

func (uc *transactionUsecase) GetTotalBuyer(ctx context.Context, from string, to string) (totalBuyer int, err error) {
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

	transaction, err := uc.transactionRepo.GetTotalBuyer(ctx, timestampFrom, timestampTo)
	if err != nil {
		return 0, err
	}

	totalBuyer = transaction.TotalBuyer
	return
}

func (uc *transactionUsecase) Fetch(ctx context.Context, request *domain.FetchTransactionRequest) (res []domain.Transaction, err error) {
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
		if err != nil {
			return nil, err
		}
		timestampTo = timestampTo.Add(24 * time.Hour)
	}

	buyerID, err := strconv.ParseInt(request.BuyerID, 10, 64)
	fishTypeID, err := strconv.ParseInt(request.FishTypeID, 10, 64)

	if err != nil {
		return []domain.Transaction{}, err
	}

	res, err = uc.transactionRepo.Fetch(ctx, timestampFrom, timestampTo, buyerID, fishTypeID)
	if err != nil {
		return nil, err
	}

	return
}

func (uc *transactionUsecase) GetByID(ctx context.Context, id int64) (res domain.Transaction, err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return
	}

	return
}

func (uc *transactionUsecase) Update(ctx context.Context, id int64, request *domain.UpdateTransactionRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedTransaction, err := uc.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedTransaction == (domain.Transaction{}) {
		return domain.ErrNotFound
	}

	transaction := &domain.Transaction{
		ID:               id,
		UserID:           1,
		TpiID:            1,
		AuctionID:        existedTransaction.AuctionID,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
		Price:            request.Price,
		CreatedAt:        existedTransaction.CreatedAt,
		UpdatedAt:        time.Now(),
	}

	err = uc.transactionRepo.Update(ctx, transaction)
	return
}

func (uc *transactionUsecase) Store(ctx context.Context, request *domain.StoreTransactionRequest) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	transaction := &domain.Transaction{
		UserID:           1,
		TpiID:            1,
		BuyerID:          request.BuyerID,
		DistributionArea: request.DistributionArea,
		CreatedAt:        time.Now(),
		UpdatedAt:        time.Now(),
	}

	for _, transactionData := range request.TransactionData {
		transaction.AuctionID = transactionData.AuctionID
		transaction.Price = transactionData.Price

		err = uc.transactionRepo.Store(ctx, transaction)
		if err != nil {
			return
		}

		err = uc.auctionRepo.UpdateStatus(ctx, transaction.AuctionID)
		if err != nil {
			return
		}
	}

	return
}

func (uc *transactionUsecase) Delete(ctx context.Context, id int64) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedTransaction, err := uc.transactionRepo.GetByID(ctx, id)
	if err != nil {
		return
	}
	if existedTransaction == (domain.Transaction{}) {
		return domain.ErrNotFound
	}

	err = uc.transactionRepo.Delete(ctx, id)
	return
}

func NewTransactionUsecase(t domain.TransactionRepository, a domain.AuctionRepository, timeout time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: t,
		auctionRepo:     a,
		contextTimeout:  timeout,
	}
}
