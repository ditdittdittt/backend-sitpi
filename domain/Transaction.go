package domain

import (
	"context"
	"time"
)

// Transaction type
type Transaction struct {
	ID        int64 `json:"id"`
	TpiID     int64 `json:"tpi_id"`
	OfficerID int64 `json:"officer_id"`

	AuctionID        int64  `json:"auction_id"`
	BuyerID          int64  `json:"buyer_id"`
	DistributionArea string `json:"distribution_area"`
	Price            int64  `json:"price"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	BuyerName  string  `json:"buyer_name,omitempty"`
	FisherName string  `json:"fisher_name,omitempty"`
	FishType   string  `json:"fish_type,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	WeightUnit string  `json:"weight_unit,omitempty"`
}

type TransactionUsecase interface {
	Fetch(ctx context.Context) (res []Transaction, err error)
	GetByID(ctx context.Context, id int64) (res Transaction, err error)
	Update(ctx context.Context, t *Transaction) (err error)
	Store(ctx context.Context, t *Transaction) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type TransactionRepository interface {
	Fetch(ctx context.Context) (res []Transaction, err error)
	GetByID(ctx context.Context, id int64) (res Transaction, err error)
	Update(ctx context.Context, t *Transaction) (err error)
	Store(ctx context.Context, t *Transaction) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
