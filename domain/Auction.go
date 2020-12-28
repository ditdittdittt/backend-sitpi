package domain

import (
	"context"
	"time"
)

// Auction type
type Auction struct {
	ID           int64 `json:"id"`
	TpiID        int64 `json:"tpi_id"`
	OfficerID    int64 `json:"officer_id"`
	CaughtFishID int64 `json:"caught_fish_id"`

	Weight     float64 `json:"weight"`
	WeightUnit string  `json:"weight_unit"`
	Status     int     `json:"status"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	FisherName string `json:"fisher_name,omitempty"`
	FishType   string `json:"fish_type,omitempty"`
	StatusName string `json:"status_name,omitempty"`
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
	UpdateStatus(ctx context.Context, id int64) (err error)
}
