package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/rs/cors"
)

var (
	crender = render.New()
)

// Setup returns router instance which is used in main package to register handlers.
func Setup(
	contact *Contact,
	r *mux.Router,
) http.Handler {

	router := r.PathPrefix("/api/v1").Subrouter()

	// endpoints for contact
	router.HandleFunc(
		"/contacts", contact.Create).
		Methods(http.MethodPost).Name("create_contact")

	router.HandleFunc(
		"/contacts/{contact_id}", contact.Get).
		Methods(http.MethodGet).Name("get_contact_by_id")

	router.HandleFunc(
		"/contacts/{contact_id}", contact.Update).
		Methods(http.MethodPut).Name("update_contact")

	router.HandleFunc(
		"/contacts/{contact_id}", contact.Delete).
		Methods(http.MethodDelete).Name("delete_contact_by_id")

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowedMethods: []string{
			http.MethodHead,
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: false,
	})

	return c.Handler(r)
}
