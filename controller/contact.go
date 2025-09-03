package controller

import (
	"contacts-api/model"
	"contacts-api/repository"
	"contacts-api/usecase"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

// Contact controller struct
type Contact struct {
	u *usecase.Contact
}

// NewContact returns a Contact controller
func NewContact(u *usecase.Contact) *Contact {
	return &Contact{u}
}

func (c *Contact) Get(cx echo.Context) (err error) {
	ctx := cx.Request().Context()
	contactID := cx.Param("contact_id")

	contact, err := c.u.Get(ctx, contactID)
	if err != nil && errors.Is(err, repository.ErrKeyNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return cx.JSON(http.StatusOK, contact)

}

func (c *Contact) Create(cx echo.Context) (err error) {
	ctx := cx.Request().Context()
	contact := new(model.Contact)
	if err = cx.Bind(contact); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	createdContact, err := c.u.Create(ctx, *contact)
	if err != nil && errors.Is(err, usecase.ErrInvalidData) {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return cx.JSON(http.StatusCreated, createdContact)
}

func (c *Contact) Update(cx echo.Context) (err error) {
	ctx := cx.Request().Context()
	contactID := cx.Param("contact_id")

	contact := new(model.Contact)
	if err = cx.Bind(contact); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	contactUpdated, err := c.u.Update(ctx, contactID, *contact)
	if err != nil && errors.Is(err, repository.ErrKeyNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}
	if err != nil && errors.Is(err, usecase.ErrInvalidData) {

		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return cx.JSON(http.StatusOK, contactUpdated)
}

func (c *Contact) Delete(cx echo.Context) (err error) {
	ctx := cx.Request().Context()
	contactID := cx.Param("contact_id")

	err = c.u.Delete(ctx, contactID)
	if err != nil && errors.Is(err, repository.ErrKeyNotFound) {
		return echo.NewHTTPError(http.StatusNotFound, err.Error())
	}

	return cx.NoContent(http.StatusNoContent)
}
