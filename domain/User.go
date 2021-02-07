package domain

import (
	"context"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	RoleID    int64     `json:"role_id"`
	StatusID  int64     `json:"status_id"`
	TpiID     int64     `json:"tpi_id"`
	Nik       string    `json:"nik"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserUsecase interface {
	Register(ctx context.Context, request *RegisterUserRequest) (jwtToken string, err error)
	Login(ctx context.Context, request *LoginUserRequest) (jwtToken string, err error)
	ChangePassword(ctx context.Context, request *ChangePasswordRequest) (err error)
}

type UserRepository interface {
	Register(ctx context.Context, u *User) (jwtToken string, err error)
	Login(ctx context.Context, u *User)
	GetByID(ctx context.Context, id int64) (res User, err error)
	GetByUsername(ctx context.Context, username string) (res User, err error)
	ChangePassword(ctx context.Context, newPassword string, id int64) (err error)
}

type RegisterUserRequest struct {
	Nik      string `json:"nik"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
