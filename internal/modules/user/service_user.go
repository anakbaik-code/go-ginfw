package user

import (
	"context"
	"errors"
	"go-fwgin/internal/config"
	"go-fwgin/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *CreateRequest) (*User, error)
	List(ctx context.Context, page int32, limit int32) ([]User, int64, error)
	UpdateProfile(ctx context.Context, user User) error
	ListActiveUsers(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	Login(ctx context.Context, req *LoginRequest) (*User, string, string, error)
	GetById(ctx context.Context, id int64) (*User, error)
	RefreshToken(ctx context.Context, token string) (string, error)
}

type serviceUser struct {
	repo   UserRepository
	config *config.Config
}

func NewServiceUser(r UserRepository, cfg *config.Config) UserService {
	return &serviceUser{repo: r, config: cfg}
}

func (s *serviceUser) Register(ctx context.Context, req *CreateRequest) (*User, error) {
	exist, err := s.repo.GetByEmail(ctx, req.Email)
	if err == nil && exist != nil {
		return nil, errors.New("email already exists")
	}
	// bcrypt password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to process password")
	}
	// mapping Dto
	user := &User{
		Name:         req.Name,
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		Phone:        req.Phone,
		Address:      req.Address,
		Role:         "user",
		IsActive:     true,
	}

	// bussiness logic -> save to repository
	id, err := s.repo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = uint64(id)
	return user, nil

}

func (s *serviceUser) List(ctx context.Context, page int32, limit int32) ([]User, int64, error) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = 10
	}

	if limit > 100 {
		limit = 100
	}

	offset := (page - 1) * limit
	users, err := s.repo.List(ctx, limit, offset)
	if err != nil {
		return nil, 0, err
	}

	// count users
	total, err := s.repo.CountUsers(ctx)
	if err != nil {
		return nil, 0, err
	}
	return users, total, nil
}
func (s *serviceUser) UpdateProfile(ctx context.Context, user User) error {
	if err := s.repo.UpdateProfile(ctx, user); err != nil {
		return err
	}
	return nil
}

func (s *serviceUser) DeleteUser(ctx context.Context, id int64) error {
	err := s.repo.Delete(ctx, id)
	return err
}

func (s *serviceUser) ListActiveUsers(ctx context.Context) ([]User, error) {
	users, err := s.repo.ListActiveUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *serviceUser) GetByEmail(ctx context.Context, email string) (*User, error) {
	user, err := s.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *serviceUser) Login(ctx context.Context, req *LoginRequest) (*User, string, string, error) {
	u, err := s.repo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password))
	if err != nil {
		return nil, "", "", err
	}

	// new access token jwt
	accessToken, err := utils.GenerateAccessToken(
		u.ID,
		u.Role,
		s.config.JwtSecret,
		s.config.JwtAccessTokenExp,
	)
	if err != nil {
		return nil, "", "", errors.New("jwt : generate access token failed")
	}
	refreshToken, err := utils.GenerateRefreshToken(
		u.ID,
		s.config.JwtSecret,
		s.config.JwtRefreshTokenExp,
	)
	if err != nil {
		return nil, "", "", errors.New("jwt : generate refresh token failed")
	}
	// save to repo database
	if err := s.repo.UpdateRefreshToken(ctx, int64(u.ID), refreshToken); err != nil {
		return nil, "", "", errors.New("jwt : failed save session login to db")
	}
	return u, accessToken, refreshToken, nil
}

func (s *serviceUser) GetById(ctx context.Context, id int64) (*User, error) {
	u, err := s.repo.GetById(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *serviceUser) RefreshToken(ctx context.Context, token string) (string, error) {
	u, err := s.repo.GetByRefreshToken(ctx, token)
	if err != nil {
		return "", errors.New("jwt : invalid refresh token or session has expired")
	}

	if !u.IsActive {
		return "", errors.New("jwt: account is no longer active")
	}

	// refresh access token
	accessToken, err := utils.GenerateAccessToken(
		u.ID,
		u.Role,
		s.config.JwtSecret,
		s.config.JwtAccessTokenExp,
	)
	if err != nil {
		return "", errors.New("jwt : generate access token failed")
	}

	return accessToken, nil
}
