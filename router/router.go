package router

import (
	"contacts-api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
)

// Setup returns router instance which is used in main package to register handlers.
func Setup(
	contact *controller.Contact,
	r *echo.Echo,
) http.Handler {

	router := r.Group("/api/v1")

	// endpoints for contact
	router.POST("/contacts", contact.Create)
	router.GET("/contacts/:contact_id", contact.Get)
	router.PUT("/contacts/:contact_id", contact.Update)
	router.DELETE("/contacts/:contact_id", contact.Delete)

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
