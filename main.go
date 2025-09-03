package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"contacts-api/controller"
	"contacts-api/repository"
	"contacts-api/router"
	"contacts-api/usecase"

	"github.com/labstack/echo/v4"
)

func main() {
	// repositories
	contactRepo := repository.NewContact()

	// usecases
	contactUsecase := usecase.NewContact(contactRepo)

	// controllers
	contactController := controller.NewContact(contactUsecase)

	// router
	echoRouter := echo.New()
	httpRouter := router.Setup(
		contactController,
		echoRouter,
	)

	server := http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%s", "8080"),
		Handler:           httpRouter,
		WriteTimeout:      15 * time.Second,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 15 * time.Second,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(
		stop,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
	)

	// Launch server start in a separate goroutine
	go func() {
		fmt.Printf("Server running in %s", server.Addr)
		err := server.ListenAndServe()
		if err != nil {
			fmt.Printf("starting server:%v ", err)
		}
	}()

	// Catch application's interrupt signals (Kill, Hang up and Interrupt)
	<-stop

}
