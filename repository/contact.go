package repository

import (
	"accelone-api/model"
	"context"

	"github.com/google/uuid"
)

type Contact struct {
}

// NewContact returns a Contact repository
func NewContact() *Contact {
	return &Contact{}
}

// Non persistent storage
var (
	contacts = make(map[string]model.Contact, 0)
)

func (u *Contact) Create(ctx context.Context, contact model.Contact) (*model.Contact, error) {
	// assign uuid as id for contact
	contact.ID = uuid.New().String()
	contacts[contact.ID] = contact

	created := contacts[contact.ID]

	return &created, nil
}

func (u *Contact) Get(ctx context.Context, contactID string) (*model.Contact, error) {

	value, ok := contacts[contactID]
	if !ok {
		return nil, ErrKeyNotFound
	}

	return &value, nil
}

func (u *Contact) Delete(ctx context.Context, contactID string) error {
	_, ok := contacts[contactID]
	if !ok {
		return ErrKeyNotFound
	}

	delete(contacts, contactID)
	return nil
}

func (u *Contact) Update(ctx context.Context, contactID string, data model.Contact) (*model.Contact, error) {
	_, ok := contacts[contactID]
	if !ok {
		return nil, ErrKeyNotFound
	}
	
	data.ID = contactID
	contacts[contactID] = data

	return &data, nil
}
