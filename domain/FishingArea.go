package domain

import (
	"context"
	"time"
)

type FishingArea struct {
	ID                  int64     `json:"id"`
	DistrictID          int64     `json:"district_id"`
	Name                string    `json:"name"`
	SouthLatitudeDegree string    `json:"south_latitude_degree"`
	SouthLatitudeMinute string    `json:"south_latitude_minute"`
	SouthLatitudeSecond string    `json:"south_latitude_second"`
	EastLongitudeDegree string    `json:"east_longitude_degree"`
	EastLongitudeMinute string    `json:"east_longitude_minute"`
	EastLongitudeSecond string    `json:"east_longitude_second"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}

type FishingAreaUsecase interface {
	Fetch(ctx context.Context) (res []FishingArea, err error)
	GetByID(ctx context.Context, id int64) (res FishingArea, err error)
	Store(ctx context.Context, request *StoreFishingAreaRequest) (err error)
	Update(ctx context.Context, id int64, request *UpdateFishingAreaRequest) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type FishingAreaRepository interface {
	Fetch(ctx context.Context) (res []FishingArea, err error)
	GetByID(ctx context.Context, id int64) (res FishingArea, err error)
	Store(ctx context.Context, fishingArea *FishingArea) (err error)
	Update(ctx context.Context, fishingArea *FishingArea) (err error)
	Delete(ctx context.Context, id int64) (err error)
}

type StoreFishingAreaRequest struct {
	Name                string `json:"name"`
	SouthLatitudeDegree string `json:"south_latitude_degree"`
	SouthLatitudeMinute string `json:"south_latitude_minute"`
	SouthLatitudeSecond string `json:"south_latitude_second"`
	EastLongitudeDegree string `json:"east_longitude_degree"`
	EastLongitudeMinute string `json:"east_longitude_minute"`
	EastLongitudeSecond string `json:"east_longitude_second"`
}

type UpdateFishingAreaRequest struct {
	Name                string `json:"name"`
	SouthLatitudeDegree string `json:"south_latitude_degree"`
	SouthLatitudeMinute string `json:"south_latitude_minute"`
	SouthLatitudeSecond string `json:"south_latitude_second"`
	EastLongitudeDegree string `json:"east_longitude_degree"`
	EastLongitudeMinute string `json:"east_longitude_minute"`
	EastLongitudeSecond string `json:"east_longitude_second"`
}
