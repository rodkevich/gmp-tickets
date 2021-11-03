package user

import (
	"context"
	"github.com/google/uuid"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, arg *User) (id string, err error)
	List(ctx context.Context, f *Filter) ([]*User, error)
	Read(ctx context.Context, userID uuid.UUID) (*User, error)
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, userID uuid.UUID) error
	Search(ctx context.Context, req *Filter) ([]*User, error)
}
