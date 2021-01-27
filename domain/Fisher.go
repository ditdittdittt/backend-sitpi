package domain

import (
	"context"
	"time"
)

// Fisher type
type Fisher struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Nik       string    `json:"nik"`
	Name      string    `json:"name"`
	Address   string    `json:"address,omitempty"`
	ShipType  string    `json:"ship_type"`
	AbkTotal  int       `json:"abk_total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	UserName string `json:"user_name,omitempty"`
}

type FisherUsecase interface {
	Fetch(ctx context.Context) (res []Fisher, err error)
	GetByID(ctx context.Context, id int64) (res Fisher, err error)
	Update(ctx context.Context, id int64, request *UpdateFisherRequest) (err error)
	Store(ctx context.Context, request *StoreFisherRequest) (err error)
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

type StoreFisherRequest struct {
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	ShipType string `json:"ship_type"`
	AbkTotal int    `json:"abk_total"`
}

type UpdateFisherRequest struct {
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	ShipType string `json:"ship_type"`
	AbkTotal int    `json:"abk_total"`
}
