package controllers

import (
	"context"

	"github.com/gofiber/fiber/v2"

	"shopping-mono/app/models"
	"shopping-mono/pkg/response"
	"shopping-mono/pkg/utils/password"
	"shopping-mono/pkg/validate"
)

func (c *Controller) CreateUser(ctx *fiber.Ctx) error {
	req := &models.CreateUserReq{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	hash, err := password.HashPassword(req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	req.Password = hash
	u, err := c.service.CreateUser(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(u))
}

func (c *Controller) GetUserById(ctx *fiber.Ctx) error {
	req := &models.GetUserByIdReq{}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.ID = int64(id)
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	user, err := c.service.GetUserById(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(user))
}

// UpdateUserById updates user by id
func (c *Controller) UpdateUserById(ctx *fiber.Ctx) error {
	req := &models.UpdateUserReq{}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.ID = int64(id)
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	hash, err := password.HashPassword(req.Password)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	req.Password = hash
	user, err := c.service.UpdateUserById(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(user))
}

//DeleteUserById deletes user by id
func (c *Controller) DeleteUserById(ctx *fiber.Ctx) error {
	req := &models.DeleteUserReq{}
	id, err := ctx.ParamsInt("id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	req.ID = int64(id)
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	err = c.service.DeleteUserById(ctx.Context(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(nil))
}

// ListAllUsers	returns all users
func (c *Controller) ListAllUsers(ctx *fiber.Ctx) error {
	users, err := c.service.GetAllUsers(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(users))
}

func (c *Controller) ListUsers(ctx *fiber.Ctx) error {
	req := &models.ListUsersReq{}
	err := ctx.QueryParser(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	err = validate.Struct(req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	users, total, err := c.service.ListUsers(context.Background(), req)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.SuccessPage(users, req.Page, req.PageSize, total))
}
