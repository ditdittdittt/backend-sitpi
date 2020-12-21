package domain

import (
	"context"
	"time"
)

type Buyer struct {
	ID        int64     `json:"id"`
	Nik       string    `json:"nik"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BuyerUsecase interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Buyer, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res Buyer, err error)
	Update(ctx context.Context, b *Buyer) (err error)
	Store(ctx context.Context, b *Buyer) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type BuyerRepository interface {
	Fetch(ctx context.Context, cursor string, num int64) (res []Buyer, nextCursor string, err error)
	GetByID(ctx context.Context, id int64) (res Buyer, err error)
	Update(ctx context.Context, b *Buyer) (err error)
	Store(ctx context.Context, b *Buyer) (err error)
	Delete(ctx context.Context, id int64) (err error)
}
