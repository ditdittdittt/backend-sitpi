package domain

import (
	"context"
	"time"
)

// Fisher type
type Fisher struct {
	ID        int64     `json:"id"`
	Nik       string    `json:"nik"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FisherUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Fisher, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res Fisher, err error)
	Update(ctx context.Context, f *Fisher) (err error)
	Store(ctx context.Context, f *Fisher) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type FisherRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Fisher, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res Fisher, err error)
	Update(ctx context.Context, f *Fisher) (err error)
	Store(ctx context.Context, f *Fisher) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
