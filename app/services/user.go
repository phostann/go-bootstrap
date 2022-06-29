package services

import (
	"context"

	"shopping-mono/app/models"
	"shopping-mono/pkg/utils/pagination"
	entUser "shopping-mono/platform/database/mysql/ent/user"
)

func (s *Service) CreateUser(ctx context.Context, req *models.CreateUserReq) (*models.User, error) {
	u, err := s.queries.DB.User.Create().SetUsername(req.Username).SetEmail(req.Email).SetPassword(req.Password).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		CreatedAt: u.CreateTime,
		UpdatedAt: u.UpdateTime,
	}, nil
}

func (s *Service) GetUserById(ctx context.Context, req *models.GetUserByIdReq) (*models.User, error) {
	u, err := s.queries.DB.User.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return &models.User{
		ID:        u.ID,
		Username:  u.Username,
		Avatar:    u.Avatar,
		Email:     u.Email,
		Gender:    u.Gender,
		CreatedAt: u.CreateTime,
		UpdatedAt: u.UpdateTime,
	}, nil
}

func (s *Service) GetUserByName(ctx context.Context, username string) (*models.User, error) {
	u, err := s.queries.DB.User.Query().Where(entUser.Username(username)).Only(ctx)
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
		CreatedAt: u.CreateTime,
		UpdatedAt: u.UpdateTime,
	}, nil
}

func (s *Service) UpdateUserById(ctx context.Context, req *models.UpdateUserReq) (*models.User, error) {
	update := s.queries.DB.User.UpdateOneID(req.ID)
	if req.Username != "" {
		update = update.SetUsername(req.Username)
	}
	if req.Email != "" {
		update = update.SetEmail(req.Email)
	}
	if req.Avatar != "" {
		update = update.SetAvatar(req.Avatar)
	}
	u, err := update.Save(ctx)
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
		CreatedAt: u.CreateTime,
		UpdatedAt: u.UpdateTime,
	}, nil
}

func (s *Service) DeleteUserById(ctx context.Context, req *models.DeleteUserReq) error {
	return s.queries.DB.User.DeleteOneID(req.ID).Exec(ctx)
}

func (s *Service) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	list, err := s.queries.DB.User.Query().All(ctx)
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
			CreatedAt: u.CreateTime,
			UpdatedAt: u.UpdateTime,
		}
	}
	return users, nil
}

func (s *Service) ListUsers(ctx context.Context, req *models.ListUsersReq) ([]*models.User, int, error) {
	total, err := s.queries.DB.User.Query().Count(ctx)
	if err != nil {
		return nil, 0, err
	}
	list, err := s.queries.DB.User.Query().Offset(pagination.Offset(req.Page, req.PageSize)).Limit(req.PageSize).All(ctx)
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
			CreatedAt: u.CreateTime,
			UpdatedAt: u.UpdateTime,
		}
	}
	return users, total, nil
}
