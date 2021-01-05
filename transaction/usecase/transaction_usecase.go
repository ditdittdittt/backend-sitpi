package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
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

func (uc *transactionUsecase) Fetch(ctx context.Context) (res []domain.Transaction, err error) {

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, err = uc.transactionRepo.Fetch(ctx)
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

func (uc *transactionUsecase) Update(ctx context.Context, t *domain.Transaction) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedTransaction, err := uc.transactionRepo.GetByID(ctx, t.ID)
	if err != nil {
		return
	}
	if existedTransaction == (domain.Transaction{}) {
		return domain.ErrNotFound
	}

	t.TpiID = 1
	t.OfficerID = 1
	t.CreatedAt = existedTransaction.CreatedAt
	t.UpdatedAt = time.Now()

	err = uc.transactionRepo.Update(ctx, t)
	return
}

func (uc *transactionUsecase) Store(ctx context.Context, t *domain.Transaction) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	t.TpiID = 1
	t.OfficerID = 1
	t.UpdatedAt = time.Now()
	t.CreatedAt = time.Now()

	err = uc.transactionRepo.Store(ctx, t)
	if err != nil {
		return
	}

	err = uc.auctionRepo.UpdateStatus(ctx, t.AuctionID)
	if err != nil {
		return
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
