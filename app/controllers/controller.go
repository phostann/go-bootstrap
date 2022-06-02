package controllers

import (
	"shopping-mono/app/services"
	"shopping-mono/pkg/configs"
)

type Controller struct {
	service *services.Service
	cfg     *configs.Config
}

// NewController creates a new Controller
func NewController(service *services.Service, cfg *configs.Config) *Controller {
	return &Controller{
		service: service,
		cfg:     cfg,
	}
}
