package services

import (
	"context"
	"shopping-mono/app/models"
	"shopping-mono/pkg/utils/pagination"
	"shopping-mono/platform/database/postgres"
)

func (s *Service) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.User, error) {
	u, err := s.queries.CreateUser(ctx, postgres.CreateUserParams{
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}, nil
}

func (s *Service) GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.User, error) {
	u, err := s.queries.GetUserById(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}, nil
}

func (s *Service) GetUserByName(ctx context.Context, username string) (*models.User, error) {
	u, err := s.queries.GetUserByName(ctx, username)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		Role:      u.Role,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}, nil
}

func (s *Service) UpdateUserById(ctx context.Context, req *models.UpdateUserReq) (*models.User, error) {
	u, err := s.queries.UpdateUserById(ctx, postgres.UpdateUserByIdParams{
		ID:       req.ID,
		Username: req.Username,
		Email:    req.Email,
		Password: req.Password,
		Avatar:   req.Avatar,
	})
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		DeletedAt: u.DeletedAt,
	}, nil
}

func (s *Service) DeleteUserById(ctx context.Context, req *models.DeleteUserReq) error {
	return s.queries.DeleteUserById(ctx, req.ID)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	list, err := s.queries.GetAllUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*models.User, len(list))
	for i, u := range list {
		users[i] = &models.User{
			ID:        u.ID,
			Username:  u.Username,
			Avatar:    u.Avatar,
			Email:     u.Email,
			Gender:    u.Gender,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt}
	}
	return users, nil
}

func (s *Service) ListUsers(ctx context.Context, req *models.ListUsersReq) ([]*models.User, int64, error) {
	total, err := s.queries.CountUsers(ctx)
	if err != nil {
		return nil, 0, err
	}
	list, err := s.queries.ListUsers(ctx, postgres.ListUsersParams{
		Offset: pagination.Offset(req.Page, req.PageSize),
		Limit:  req.PageSize,
	})
	if err != nil {
		return nil, 0, err
	}
	users := make([]*models.User, len(list))
	for i, u := range list {
		users[i] = &models.User{
			ID:        u.ID,
			Username:  u.Username,
			Avatar:    u.Avatar,
			Email:     u.Email,
			Gender:    u.Gender,
			Role:      u.Role,
			CreatedAt: u.CreatedAt,
			UpdatedAt: u.UpdatedAt,
			DeletedAt: u.DeletedAt,
		}
	}
	return users, total, nil
}
