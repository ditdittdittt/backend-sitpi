package domain

import (
	"context"
	"time"
)

// Transaction type
type Transaction struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id,omitempty"`
	TpiID            int64     `json:"tpi_id,omitempty"`
	AuctionID        int64     `json:"auction_id,omitempty"`
	BuyerID          int64     `json:"buyer_id,omitempty"`
	DistributionArea string    `json:"distribution_area,omitempty"`
	Price            int64     `json:"price,omitempty"`
	CreatedAt        time.Time `json:"created_at,omitempty"`
	UpdatedAt        time.Time `json:"updated_at,omitempty"`

	BuyerName  string  `json:"buyer_name,omitempty"`
	FisherName string  `json:"fisher_name,omitempty"`
	FishType   string  `json:"fish_type,omitempty"`
	Weight     float64 `json:"weight,omitempty"`
	WeightUnit string  `json:"weight_unit,omitempty"`

	TotalBuyer int `json:"total_buyer,omitempty"`
}

type TransactionUsecase interface {
	Fetch(ctx context.Context, request *FetchTransactionRequest) (res []Transaction, err error)
	GetByID(ctx context.Context, id int64) (res Transaction, err error)
	Update(ctx context.Context, id int64, request *UpdateTransactionRequest) (err error)
	Store(ctx context.Context, request *StoreTransactionRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
	GetTotalBuyer(ctx context.Context, from string, to string) (totalBuyer int, err error)
}

type TransactionRepository interface {
	Fetch(ctx context.Context, from time.Time, to time.Time, buyerID int64, fishTypeID int64) (res []Transaction, err error)
	GetByID(ctx context.Context, id int64) (res Transaction, err error)
	Update(ctx context.Context, t *Transaction) (err error)
	Store(ctx context.Context, t *Transaction) (err error)
	Delete(ctx context.Context, id int64) (err error)
	GetTotalBuyer(ctx context.Context, from time.Time, to time.Time) (res Transaction, err error)
}

type FetchTransactionRequest struct {
	From       string
	To         string
	BuyerID    string
	FishTypeID string
}

type StoreTransactionRequest struct {
	BuyerID          int64  `json:"buyer_id" validate:"required"`
	DistributionArea string `json:"distribution_area" validate:"required"`
	TransactionData  []struct {
		AuctionID int64 `json:"auction_id" validate:"required"`
		Price     int64 `json:"price" validate:"required"`
	} `json:"transaction_data"`
}

type UpdateTransactionRequest struct {
	BuyerID          int64  `json:"buyer_id" validate:"required"`
	DistributionArea string `json:"distribution_area" validate:"required"`
	Price            int64  `json:"price" validate:"required"`
}
