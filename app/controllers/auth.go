package controllers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"shopping-mono/app/models"
	"shopping-mono/pkg/middlewares"
	"shopping-mono/pkg/response"
	"shopping-mono/pkg/utils/password"
	"shopping-mono/pkg/validate"
)

func (c *Controller) Login(ctx *fiber.Ctx) error {
	req := &models.LoginReq{}
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	if err := validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(response.Error(err))
	}
	user, err := c.service.GetUserByName(ctx.Context(), req.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	ok := password.CheckPassordHash(req.Password, user.Password)
	if !ok {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(errors.New("invalid username or password")))
	}
	accessClaims := &middlewares.CustomClaims{
		Username: user.Username,
		Role:     user.Role,
		Type:     middlewares.AccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			Issuer:    c.cfg.JWT.Issuer,
		},
	}
	refreshClaims := &middlewares.CustomClaims{
		Username: user.Username,
		Role:     user.Role,
		Type:     middlewares.RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			Issuer:    c.cfg.JWT.Issuer,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(c.cfg.JWT.Secret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(c.cfg.JWT.Secret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(fiber.Map{"access_token": accessTokenString, "refresh_token": refreshTokenString}))
}

func (c *Controller) GetProfile(ctx *fiber.Ctx) error {
	claims, ok := ctx.Locals("claims").(*middlewares.CustomClaims)
	if !ok || claims.Type != middlewares.AccessToken {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(errors.New("invalid access token")))
	}
	user, err := c.service.GetUserByName(ctx.Context(), claims.Username)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(user))
}

func (c *Controller) Refresh(ctx *fiber.Ctx) error {
	claims, ok := ctx.Locals("claims").(*middlewares.CustomClaims)
	if !ok || claims.Type != middlewares.RefreshToken {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(errors.New("invalid refresh token")))
	}
	accessClaims := &middlewares.CustomClaims{
		Username: claims.Username,
		Role:     claims.Role,
		Type:     middlewares.AccessToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 30)),
			Issuer:    c.cfg.JWT.Issuer,
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString([]byte(c.cfg.JWT.Secret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	refreshClaims := &middlewares.CustomClaims{
		Username: claims.Username,
		Role:     claims.Role,
		Type:     middlewares.RefreshToken,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 7)),
			Issuer:    c.cfg.JWT.Issuer,
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(c.cfg.JWT.Secret))
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(response.Error(err))
	}
	return ctx.JSON(response.Success(fiber.Map{"access_token": accessTokenString, "refresh_token": refreshTokenString}))
}
