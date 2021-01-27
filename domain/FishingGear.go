package domain

import (
	"context"
	"time"
)

type FishingGear struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FishingGearUsecase interface {
	Fetch(ctx context.Context) (res []FishingGear, err error)
	GetByID(ctx context.Context, id int64) (res FishingGear, err error)
	Update(ctx context.Context, id int64, request *UpdateFishingGearRequest) (err error)
	Store(ctx context.Context, request *StoreFishingGearRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type FishingGearRepository interface {
	Fetch(ctx context.Context) (res []FishingGear, err error)
	GetByID(ctx context.Context, id int64) (res FishingGear, err error)
	Update(ctx context.Context, fg *FishingGear) (err error)
	Store(ctx context.Context, fg *FishingGear) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type StoreFishingGearRequest struct {
	Name string `json:"name"`
}

type UpdateFishingGearRequest struct {
	Name string `json:"name"`
}
