package models

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Gender    string    `json:"gender"`
}

type CreateUserReq struct {
	Username string `json:"username" validate:"required,max=10"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
	Email    string `json:"email" validate:"required,email,max=30"`
}

type GetUserByIdReq struct {
	ID int `json:"id" validate:"required"`
}

type DeleteUserReq struct {
	ID int `json:"id" validate:"required"`
}

type ListUsersReq struct {
	Page     int `query:"page" validate:"required"`
	PageSize int `query:"page_size" validate:"required"`
}

type UpdateUserReq struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required,max=10"`
	Avatar   string `json:"avatar" validate:"omitempty,max=255"`
	Email    string `json:"email" validate:"required,email,max=30"`
	Gender   string `json:"gender" validate:"required,oneof=male female"`
}
