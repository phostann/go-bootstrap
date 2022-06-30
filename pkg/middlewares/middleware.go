package middlewares

import "shopping-mono/pkg/configs"

type Middleware struct {
	cfg configs.Config
}

// NewMiddleware creates a new Middleware
func NewMiddleware(cfg configs.Config) *Middleware {
	return &Middleware{
		cfg: cfg,
	}
}
