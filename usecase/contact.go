package usecase

import (
	"contacts-api/model"
	"context"
)

type Contact struct {
	contactRepo ContactRepository
}

// NewContact returns a Contact usecase
func NewContact(
	contactRepo ContactRepository,
) *Contact {
	return &Contact{
		contactRepo,
	}
}

func (u *Contact) Get(ctx context.Context, contactID string) (*model.Contact, error) {

	return u.contactRepo.Get(ctx, contactID)
}

func (u *Contact) Create(ctx context.Context, contact model.Contact) (*model.Contact, error) {
	// Validate data
	if contact.Name == "" || contact.Phone == "" {
		return nil, ErrInvalidData

	}
	return u.contactRepo.Create(ctx, contact)
}

func (u *Contact) Update(ctx context.Context, id string, contact model.Contact) (*model.Contact, error) {
	// Validate data
	if contact.Name == "" || contact.Phone == "" {
		return nil, ErrInvalidData

	}
	return u.contactRepo.Update(ctx, id, contact)
}

func (u *Contact) Delete(ctx context.Context, contactID string) error {

	return u.contactRepo.Delete(ctx, contactID)
}
