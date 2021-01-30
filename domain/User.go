package domain

import "time"

type User struct {
	ID        int64     `json:"id"`
	RoleID    int64     `json:"role_id"`
	StatusID  int64     `json:"status_id"`
	TpiID     int64     `json:"tpi_id"`
	Nik       string    `json:"nik"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
