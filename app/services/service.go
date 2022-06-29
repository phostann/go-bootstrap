package services

import (
	"shopping-mono/platform/database/mysql"
)

type Service struct {
	queries *mysql.Queries
}

// NewService creates a new Service
func NewService(queries *mysql.Queries) *Service {
	return &Service{
		queries: queries,
	}
}
