package domain

import (
	"context"
	"time"
)

type Buyer struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Nik       string    `json:"nik"`
	Name      string    `json:"name"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserName string `json:"user_name,omitempty"`
}

type BuyerUsecase interface {
	Fetch(ctx context.Context) (res []Buyer, err error)
	GetByID(ctx context.Context, id int64) (res Buyer, err error)
	Update(ctx context.Context, id int64, request *UpdateBuyerRequest) (err error)
	Store(ctx context.Context, request *StoreBuyerRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
	Inquiry(ctx context.Context) (res []Buyer, err error)
}

type BuyerRepository interface {
	Fetch(ctx context.Context) (res []Buyer, err error)
	GetByID(ctx context.Context, id int64) (res Buyer, err error)
	Update(ctx context.Context, b *Buyer) (err error)
	Store(ctx context.Context, b *Buyer) (err error)
	Delete(ctx context.Context, id int64) (err error)
	Inquiry(ctx context.Context) (res []Buyer, err error)
}

type StoreBuyerRequest struct {
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type UpdateBuyerRequest struct {
	Nik     string `json:"nik"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
