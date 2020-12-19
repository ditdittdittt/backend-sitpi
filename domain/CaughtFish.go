package domain

import (
	"context"
	"time"
)

// Caught fish ..
type CaughtFish struct {
	ID          int64     `json:"id"`
	TpiID       int64     `json:"tpi_id"`
	OfficerID   int64     `json:"officer_id"`
	FisherID    int64     `json:"fisher_id"`
	FishTypeID  int64     `json:"fish_type_id"`
	Weight      float64   `json:"weight"`
	WeightUnit  string    `json:"weight_unit"`
	FishingGear string    `json:"fishing_gear"`
	FishingArea string    `json:"fishing_area"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CaughtFishUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []CaughtFish, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res CaughtFish, err error)
	Update(ctx context.Context, c *CaughtFish) (err error)
	Store(ctx context.Context, c *CaughtFish) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type CaughtFishRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []CaughtFish, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res CaughtFish, err error)
	Update(ctx context.Context, c *CaughtFish) (err error)
	Store(ctx context.Context, c *CaughtFish) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
