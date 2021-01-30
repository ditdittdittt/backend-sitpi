package domain

import "time"

type Tpi struct {
	ID         int64     `json:"id"`
	DistrictID int64     `json:"district_id"`
	Name       string    `json:"name"`
	Location   string    `json:"location"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
