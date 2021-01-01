package domain

import (
	"context"
	"time"
)

// Fisher type
type Fisher struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Nik       string    `json:"nik"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type FisherUsecase interface {
	Fetch(ctx context.Context) (res []Fisher, err error)
	GetByID(ctx context.Context, id int64) (res Fisher, err error)
	Update(ctx context.Context, f *Fisher) (err error)
	Store(ctx context.Context, f *Fisher) (err error)
	Delete(ctx context.Context, id int64) (err error)
	Inquiry(ctx context.Context) (res []Fisher, err error)
}

type FisherRepository interface {
	Fetch(ctx context.Context) (res []Fisher, err error)
	GetByID(ctx context.Context, id int64) (res Fisher, err error)
	Update(ctx context.Context, f *Fisher) (err error)
	Store(ctx context.Context, f *Fisher) (err error)
	Delete(ctx context.Context, id int64) (err error)
	Inquiry(ctx context.Context) (res []Fisher, err error)
}
