package usecase

import (
	"context"
	"github.com/ditdittdittt/backend-sitpi/domain"
	"time"
)

type transactionUsecase struct {
	transactionRepo domain.TransactionRepository
	contextTimeout  time.Duration
}

func (uc *transactionUsecase) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Transaction, nextCursor string, err error) {
	if num == 0 {
		num = 10
	}

	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	res, nextCursor, err = uc.transactionRepo.Fetch(ctx, cursor, num)
	if err != nil {
		return nil, "", err
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

	t.UpdatedAt = time.Now()
	err = uc.transactionRepo.Update(ctx, t)
	return
}

func (uc *transactionUsecase) Store(ctx context.Context, t *domain.Transaction) (err error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err = uc.transactionRepo.Store(ctx, t)
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

func NewTransactionUsecase(t domain.TransactionRepository, timeout time.Duration) domain.TransactionUsecase {
	return &transactionUsecase{
		transactionRepo: t,
		contextTimeout:  timeout,
	}
}
