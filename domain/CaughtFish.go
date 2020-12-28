package domain

import (
	"context"
	"time"
)

// Caught fish ..
type CaughtFish struct {
	ID        int64 `json:"id"`
	TpiID     int64 `json:"tpi_id"`
	OfficerID int64 `json:"officer_id"`

	FisherID    int64   `json:"fisher_id"`
	FishTypeID  int64   `json:"fish_type_id"`
	Weight      float64 `json:"weight"`
	WeightUnit  string  `json:"weight_unit"`
	FishingGear string  `json:"fishing_gear"`
	FishingArea string  `json:"fishing_area"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	FisherName string `json:"fisher_name,omitempty"`
	FisherNik  string `json:"fisher_nik,omitempty"`
	FishType   string `json:"fish_type,omitempty"`
}

type CaughtFishUsecase interface {
	Fetch(ctx context.Context) (res []CaughtFish, err error)
	GetByID(ctx context.Context, id int64) (res CaughtFish, err error)
	Update(ctx context.Context, c *CaughtFish) (err error)
	Store(ctx context.Context, c *CaughtFish, a *Auction) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type CaughtFishRepository interface {
	Fetch(ctx context.Context) (res []CaughtFish, err error)
	GetByID(ctx context.Context, id int64) (res CaughtFish, err error)
	Update(ctx context.Context, c *CaughtFish) (err error)
	Store(ctx context.Context, c *CaughtFish) (lastID int64, err error)
	Delete(ctx context.Context, id int64) (err error)
}
