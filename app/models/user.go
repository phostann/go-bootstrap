package models

import (
	"database/sql"
	"time"

	"shopping-mono/platform/database/postgres"
)

type User struct {
	ID        int64               `json:"id"`
	Username  string              `json:"username"`
	Avatar    string              `json:"avatar"`
	Email     string              `json:"email"`
	Password  string              `json:"-"`
	Role      postgres.RoleEnum   `json:"role"`
	CreatedAt time.Time           `json:"created_at"`
	UpdatedAt time.Time           `json:"updated_at"`
	DeletedAt sql.NullTime        `json:"deleted_at"`
	Gender    postgres.GenderEnum `json:"gender"`
}

type CreateUserReq struct {
	Username string `json:"username" validate:"required,max=10"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
	Email    string `json:"email" validate:"required,email,max=30"`
}

type GetUserByIdReq struct {
	ID int64 `json:"id" validate:"required"`
}

type DeleteUserReq struct {
	ID int64 `json:"id" validate:"required"`
}

type ListUsersReq struct {
	Page     int32 `query:"page" validate:"required"`
	PageSize int32 `query:"page_size" validate:"required"`
}

type UpdateUserReq struct {
	ID       int64               `json:"id" validate:"required"`
	Username string              `json:"username" validate:"required,max=10"`
	Password string              `json:"password" validate:"required,min=8,max=20"`
	Avatar   string              `json:"avatar" validate:"omitempty,max=255"`
	Email    string              `json:"email" validate:"required,email,max=30"`
	Gender   postgres.GenderEnum `json:"gender" validate:"required,oneof=male female"`
}
