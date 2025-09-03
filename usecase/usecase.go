package usecase

import (
	"contacts-api/model"
	"context"
	"errors"
)

var (
	ErrInvalidData = errors.New("invalid-data")
)

// KEEP ALL INTERFACES HERE
//
//go:generate mockery --all --output=./mocks --case underscore
type ContactRepository interface {
	Create(ctx context.Context, contact model.Contact) (*model.Contact, error)
	Get(ctx context.Context, contactID string) (*model.Contact, error)
	Delete(ctx context.Context, contactID string) error
	Update(ctx context.Context, contactID string, data model.Contact) (*model.Contact, error)
}
