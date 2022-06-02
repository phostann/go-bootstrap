package services

import "shopping-mono/platform/database/postgres"

type Service struct {
	queries *postgres.Queries
}

// NewService creates a new Service
func NewService(queries *postgres.Queries) *Service {
	return &Service{
		queries: queries,
	}
}
