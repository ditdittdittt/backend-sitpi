package domain

import (
	"context"
	"time"
)

// Auction type
type Auction struct {
	ID          int64     `json:"id"`
	TpiID       int64     `json:"tpi_id"`
	FisherID    int64     `json:"fisher_id"`
	OfficerID   int64     `json:"officer_id"`
	FishTypeID  int64     `json:"fish_type_id"`
	Weight      float64   `json:"weight"`
	WeightUnit  string    `json:"weight_unit"`
	FishingGear string    `json:"fishing_gear"`
	FishingArea string    `json:"fishing_area"`
	Status      int       `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	FisherName string `json:"fisher_name"`
	FishType   string `json:"fish_type"`
	StatusName string `json:"status_name"`
}

type AuctionUsecase interface {
	Fetch(ctx context.Context) (res []Auction, err error)
	GetByID(ctx context.Context, id int64) (res Auction, err error)
	Update(ctx context.Context, a *Auction) (err error)
	Store(ctx context.Context, a *Auction) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type AuctionRepository interface {
	Fetch(ctx context.Context) (res []Auction, err error)
	GetByID(ctx context.Context, id int64) (res Auction, err error)
	Update(ctx context.Context, a *Auction) (err error)
	Store(ctx context.Context, a *Auction) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
