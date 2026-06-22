package user

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *CreateRequest) (*User, error)
	List(ctx context.Context, page int, limit int) ([]User, error)
	UpdateProfile(ctx context.Context, user User) error
}

type serviceUser struct {
	repo RepositoryUser
}

func NewServiceUser(repo RepositoryUser) UserService {
	return &serviceUser{repo: repo}
}

func (s *serviceUser) Register(ctx context.Context, req *CreateRequest) (*User, error) {
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
	user.ID = id
	return user, nil

}

func (s *serviceUser) List(ctx context.Context, page int, limit int) ([]User, error) {
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

	return s.repo.List(
		ctx,
		int32(limit),
		int32(offset),
	)
}
func (s *serviceUser) UpdateProfile(ctx context.Context, user User) error {
	if err := s.repo.UpdateProfile(ctx, user); err != nil {
		return err
	}
	return nil
}
