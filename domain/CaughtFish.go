package domain

import (
	"context"
	"time"
)

// Caught fish ..
type CaughtFish struct {
	// Database model
	ID        int64 `json:"id"`
	TpiID     int64 `json:"tpi_id"`
	OfficerID int64 `json:"officer_id"`

	FisherID      int64   `json:"fisher_id"`
	FishTypeID    int64   `json:"fish_type_id"`
	WeightUnitID  int64   `json:"weight_unit_id"`
	FishingGearID int64   `json:"fishing_gear_id"`
	Weight        float64 `json:"weight"`
	FishingArea   string  `json:"fishing_area"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// Index and get by id attribute
	WeightUnit  string `json:"weight_unit,omitempty"`
	FishingGear string `json:"fishing_gear,omitempty"`
	FisherName  string `json:"fisher_name,omitempty"`
	FisherNik   string `json:"fisher_nik,omitempty"`
	FishType    string `json:"fish_type,omitempty"`

	// Other API attribute
	TotalProduction float64 `json:"total_production,omitempty"`
	TotalFisher     int     `json:"total_fisher,omitempty"`
}

type CaughtFishUsecase interface {
	Fetch(ctx context.Context) (res []CaughtFish, err error)
	GetByID(ctx context.Context, id int64) (res CaughtFish, err error)
	Update(ctx context.Context, c *CaughtFish) (err error)
	Store(ctx context.Context, c *CaughtFish, a *Auction) (err error)
	Delete(ctx context.Context, id int64) (err error)
	GetTotalProduction(ctx context.Context, from string, to string) (totalProduction float64, err error)
	GetTotalFisher(ctx context.Context, from string, to string) (totalFisher int, err error)
}

type CaughtFishRepository interface {
	Fetch(ctx context.Context) (res []CaughtFish, err error)
	GetByID(ctx context.Context, id int64) (res CaughtFish, err error)
	Update(ctx context.Context, c *CaughtFish) (err error)
	Store(ctx context.Context, c *CaughtFish) (lastID int64, err error)
	Delete(ctx context.Context, id int64) (err error)
	GetTotalProduction(ctx context.Context, from time.Time, to time.Time) (res CaughtFish, err error)
	GetTotalFisher(ctx context.Context, from time.Time, to time.Time) (res CaughtFish, err error)
}
