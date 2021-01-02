package domain

import (
	"context"
	"time"
)

type WeightUnit struct {
	ID        int64     `json:"id"`
	Unit      string    `json:"unit"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WeightUnitUsecase interface {
	Fetch(ctx context.Context) (res []WeightUnit, err error)
	GetByID(ctx context.Context, id int64) (res WeightUnit, err error)
	Update(ctx context.Context, wu *WeightUnit) (err error)
	Store(ctx context.Context, wu *WeightUnit) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type WeightUnitRepository interface {
	Fetch(ctx context.Context) (res []WeightUnit, err error)
	GetByID(ctx context.Context, id int64) (res WeightUnit, err error)
	Update(ctx context.Context, wu *WeightUnit) (err error)
	Store(ctx context.Context, wu *WeightUnit) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
