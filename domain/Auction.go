package domain

import (
	"context"
	"time"
)

// Auction type
type Auction struct {
	ID           int64     `json:"id"`
	TpiID        int64     `json:"tpi_id,omitempty"`
	CaughtFishID int64     `json:"caught_fish_id,omitempty"`
	StatusID     int64     `json:"status,omitempty"`
	SoldAt       time.Time `json:"sold_at,omitempty"`
	CreatedAt    time.Time `json:"created_at,omitempty"`
	UpdatedAt    time.Time `json:"updated_at,omitempty"`

	Weight     float64 `json:"weight,omitempty"`
	WeightUnit string  `json:"weight_unit,omitempty"`
	FisherName string  `json:"fisher_name,omitempty"`
	FishType   string  `json:"fish_type,omitempty"`
	StatusName string  `json:"status_name,omitempty"`
}

type AuctionUsecase interface {
	Fetch(ctx context.Context, request *FetchAuctionRequest) (res []Auction, err error)
	GetByID(ctx context.Context, id int64) (res Auction, err error)
	Delete(ctx context.Context, id int64) (err error)
	Inquiry(ctx context.Context) (res []Auction, err error)
}

type AuctionRepository interface {
	Fetch(ctx context.Context, from time.Time, to time.Time, auctionID int64, fisherID int64, fishTypeID int64, statusID int64) (res []Auction, err error)
	GetByID(ctx context.Context, id int64) (res Auction, err error)
	Store(ctx context.Context, a *Auction) (err error)
	Delete(ctx context.Context, id int64) (err error)
	UpdateStatus(ctx context.Context, id int64) (err error)
	Inquiry(ctx context.Context, from time.Time, to time.Time) (res []Auction, err error)
}

type FetchAuctionRequest struct {
	From       string
	To         string
	AuctionID  string
	FisherID   string
	FishTypeID string
	StatusID   string
}
