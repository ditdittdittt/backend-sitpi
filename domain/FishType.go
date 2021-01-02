package domain

import (
	"context"
	"time"
)

type FishType struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FishTypeUsecase interface {
	Fetch(ctx context.Context) (res []FishType, err error)
	GetByID(ctx context.Context, id int64) (res FishType, err error)
	Update(ctx context.Context, ft *FishType) (err error)
	Store(ctx context.Context, ft *FishType) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type FishTypeRepository interface {
	Fetch(ctx context.Context) (res []FishType, err error)
	GetByID(ctx context.Context, id int64) (res FishType, err error)
	Update(ctx context.Context, ft *FishType) (err error)
	Store(ctx context.Context, ft *FishType) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
