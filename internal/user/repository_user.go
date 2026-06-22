package user

import (
	"context"
	"database/sql"

	"go-fwgin/internal/database"
)

type RepositoryUser interface {
	Create(ctx context.Context, u *User) (int64, error)
	List(ctx context.Context, limit int32, offset int32) ([]User, error)
	UpdateProfile(ctx context.Context, u User) error
}

type repositoryUser struct {
	queries *database.Queries
}

func NewRepositoryUser(q *database.Queries) RepositoryUser {
	return &repositoryUser{queries: q}
}

func (r *repositoryUser) Create(ctx context.Context, u *User) (int64, error) {
	// convert string to sql.nullstring
	var phone, address sql.NullString

	if u.Phone != "" {
		phone = sql.NullString{String: u.Phone, Valid: true}
	}
	if u.Address != "" {
		address = sql.NullString{String: u.Address, Valid: true}
	}
	result, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Name:         u.Name,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		Phone:        phone,
		Address:      address,
		Role:         u.Role,
		IsActive:     u.IsActive,
	})
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *repositoryUser) List(ctx context.Context, limit int32, offset int32) ([]User, error) {
	rows, err := r.queries.ListUsers(ctx, database.ListUsersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	users := make([]User, 0, len(rows))
	for _, row := range rows {
		users = append(users, User{
			ID:       row.ID.Int64,
			Name:     row.Name,
			Email:    row.Email,
			Phone:    row.Phone.String,
			Address:  row.Address.String,
			Role:     row.Role,
			IsActive: row.IsActive,
		})
	}
	return users, nil
}

func (r *repositoryUser) UpdateProfile(ctx context.Context, u User) error {
	var id sql.NullInt64
	var phone, address sql.NullString
	if u.Phone != "" {
		phone = sql.NullString{String: u.Phone, Valid: true}
	}
	if u.Address != "" {
		address = sql.NullString{String: u.Address, Valid: true}
	}
	if u.ID != 0 {
		id = sql.NullInt64{Int64: u.ID, Valid: true}
	}
	err := r.queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:      id,
		Name:    u.Name,
		Phone:   phone,
		Address: address,
	})

	return err
}
