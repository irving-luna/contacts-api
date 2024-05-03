package controller

import (
	"accelone-api/model"
	"accelone-api/repository"
	"accelone-api/usecase"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

// Contact controller struct
type Contact struct {
	u *usecase.Contact
}

// NewContact returns a Contact controller
func NewContact(u *usecase.Contact) *Contact {
	return &Contact{u}
}

func (c *Contact) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactID := mux.Vars(r)["contact_id"]

	contact, err := c.u.Get(ctx, contactID)
	if err != nil && errors.Is(err, repository.ErrKeyNotFound) {
		crender.JSON(w, http.StatusNotFound, map[string]string{
			"status": "404",
			"error":  err.Error(),
		})
		return
	}
	crender.JSON(w, http.StatusOK, contact)
}

func (c *Contact) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	body, err := readBodyAs[model.Contact](r.Body)
	if err != nil {
		crender.JSON(w, http.StatusBadRequest, map[string]string{
			"status": "400",
			"error":  err.Error(),
		})
		return
	}
	contact, err := c.u.Create(ctx, *body)
	if err != nil && errors.Is(err, usecase.ErrInvalidData) {
		crender.JSON(w, http.StatusBadRequest, map[string]string{
			"status": "400",
			"error":  err.Error(),
		})
		return
	}

	crender.JSON(w, http.StatusOK, contact)
}

func (c *Contact) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactID := mux.Vars(r)["contact_id"]

	body, err := readBodyAs[model.Contact](r.Body)
	if err != nil {
		crender.JSON(w, http.StatusBadRequest, map[string]string{
			"status": "400",
			"error":  err.Error(),
		})
		return
	}

	contact, err := c.u.Update(ctx, contactID, *body)
	if err != nil && errors.Is(err, repository.ErrKeyNotFound) {
		crender.JSON(w, http.StatusNotFound, map[string]string{
			"status": "404",
			"error":  err.Error(),
		})
		return
	}
	if err != nil && errors.Is(err, usecase.ErrInvalidData) {
		crender.JSON(w, http.StatusBadRequest, map[string]string{
			"status": "400",
			"error":  err.Error(),
		})
		return
	}
	crender.JSON(w, http.StatusOK, contact)
}

func (c *Contact) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	contactID := mux.Vars(r)["contact_id"]

	err := c.u.Delete(ctx, contactID)
	if err != nil && errors.Is(err, repository.ErrKeyNotFound) {
		crender.JSON(w, http.StatusNotFound, map[string]string{
			"status": "404",
			"error":  err.Error(),
		})
		return
	}
	crender.JSON(w, http.StatusNoContent, nil)
}
