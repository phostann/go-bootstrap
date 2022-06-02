package middlewares

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"

	"shopping-mono/pkg/response"
	"shopping-mono/platform/database/postgres"
)

type TokenType string

const (
	AccessToken  = "access_token"
	RefreshToken = "refresh_token"
)

type CustomClaims struct {
	Username string            `json:"username"`
	Role     postgres.RoleEnum `json:"role"`
	Type     TokenType         `json:"type"`
	jwt.RegisteredClaims
}

func (m *Middleware) JWTProtected(ctx *fiber.Ctx) error {
	token, err := getJwtToken(ctx, m.cfg.JWT.Secret)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(err))
	}
	if cliams, ok := token.Claims.(*CustomClaims); ok && token.Valid && cliams.Type == AccessToken {
		ctx.Locals("cliams", cliams)
		return ctx.Next()
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(errors.New("invalid token")))
	}
}

func (m *Middleware) JWTRefreshProtected(ctx *fiber.Ctx) error {
	token, err := getJwtToken(ctx, m.cfg.JWT.Secret)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(err))
	}
	if cliams, ok := token.Claims.(*CustomClaims); ok && token.Valid && cliams.Type == RefreshToken {
		ctx.Locals("cliams", cliams)
		return ctx.Next()
	} else {
		return ctx.Status(fiber.StatusUnauthorized).JSON(response.Error(errors.New("invalid token")))
	}
}

func getJwtToken(ctx *fiber.Ctx, secret string) (*jwt.Token, error) {
	var token *jwt.Token
	auth := ctx.Get("Authorization")
	if auth == "" {
		return nil, errors.New("authorization header not found")
	}
	splits := strings.Split(auth, " ")
	if len(splits) != 2 {
		return nil, errors.New("authorization header is not in correct format")
	}
	tokenString := splits[1]
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	return token, err
}
