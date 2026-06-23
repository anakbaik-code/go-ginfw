package user

import (
	"context"
	"database/sql"

	"go-fwgin/internal/database"
)

type RepositoryUser interface {
	Create(ctx context.Context, user *User) (int64, error)
	List(ctx context.Context, limit int32, offset int32) ([]User, error)
	UpdateProfile(ctx context.Context, user User) error
	CountUsers(ctx context.Context) (int64, error)
	Delete(ctx context.Context, id int64) error
	GetActiveUsers(ctx context.Context) ([]User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetById(ctx context.Context, id int64) (*User, error)
	GetByRefreshToken(ctx context.Context, token string) (*User, error)
	UpdateRefreshToken(ctx context.Context, id int64, refreshToken string) error
}

type repositoryUser struct {
	queries *database.Queries
}

func NewRepositoryUser(q *database.Queries) RepositoryUser {
	return &repositoryUser{queries: q}
}

func (r *repositoryUser) CountUsers(ctx context.Context) (int64, error) {
	total, err := r.queries.CountUsers(ctx)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *repositoryUser) Create(ctx context.Context, user *User) (int64, error) {
	// convert string to sql.nullstring
	var phone, address sql.NullString

	if user.Phone != "" {
		phone = sql.NullString{String: user.Phone, Valid: true}
	}
	if user.Address != "" {
		address = sql.NullString{String: user.Address, Valid: true}
	}
	result, err := r.queries.CreateUser(ctx, database.CreateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		Phone:        phone,
		Address:      address,
		Role:         user.Role,
		IsActive:     user.IsActive,
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
			ID:       int64(row.ID),
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

func (r *repositoryUser) UpdateProfile(ctx context.Context, user User) error {
	var phone, address sql.NullString
	if user.Phone != "" {
		phone = sql.NullString{String: user.Phone, Valid: true}
	}
	if user.Address != "" {
		address = sql.NullString{String: user.Address, Valid: true}
	}
	err := r.queries.UpdateUser(ctx, database.UpdateUserParams{
		ID:      uint64(user.ID),
		Name:    user.Name,
		Phone:   phone,
		Address: address,
	})

	return err
}
func (r *repositoryUser) Delete(ctx context.Context, id int64) error {
	err := r.queries.DeleteUser(ctx, uint64(id))
	return err
}

func (r *repositoryUser) GetActiveUsers(ctx context.Context) ([]User, error) {
	rows, err := r.queries.GetActiveUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]User, 0, len(rows))
	for _, row := range rows {
		users = append(users, User{
			ID:        int64(row.ID),
			Name:      row.Name,
			Email:     row.Email,
			Role:      row.Role,
			CreatedAt: row.CreatedAt,
		})
	}
	return users, nil

}
func (r *repositoryUser) GetByEmail(ctx context.Context, email string) (*User, error) {
	row, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	// Mapping
	user := &User{
		ID:           int64(row.ID),
		Name:         row.Name,
		Email:        row.Email,
		PasswordHash: row.PasswordHash,
		Role:         row.Role,
		IsActive:     row.IsActive,
		CreatedAt:    row.CreatedAt,
		UpdatedAt:    row.UpdatedAt,
	}
	if row.Phone.Valid {
		user.Phone = row.Phone.String
	}
	if row.Address.Valid {
		user.Address = row.Address.String
	}
	return user, nil
}
func (r *repositoryUser) GetById(ctx context.Context, id int64) (*User, error) {
	row, err := r.queries.GetUserByID(ctx, uint64(id))
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:        int64(row.ID),
		Name:      row.Name,
		Email:     row.Email,
		Role:      row.Role,
		IsActive:  row.IsActive,
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
	if row.Phone.Valid {
		user.Phone = row.Phone.String
	}
	if row.Address.Valid {
		user.Address = row.Address.String
	}
	return user, nil
}
func (r *repositoryUser) GetByRefreshToken(ctx context.Context, token string) (*User, error) {
	dbToken := sql.NullString{
		String: token,
		Valid:  true,
	}
	row, err := r.queries.GetUserByRefreshToken(ctx, dbToken)
	if err != nil {
		return nil, err
	}
	user := &User{
		ID:       int64(row.ID),
		Name:     row.Name,
		Email:    row.Email,
		Role:     row.Role,
		IsActive: row.IsActive,
	}
	return user, nil
}

func (r *repositoryUser) UpdateRefreshToken(ctx context.Context, id int64, refreshToken string) error {
	tokenFresh := sql.NullString{
		String: refreshToken,
		Valid:  true,
	}
	if err := r.queries.UpdateUserRefreshToken(ctx, database.UpdateUserRefreshTokenParams{
		ID:           uint64(id),
		RefreshToken: tokenFresh,
	}); err != nil {
		return err
	}
	return nil
}
