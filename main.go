package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"accelone-api/controller"
	"accelone-api/repository"
	"accelone-api/usecase"

	"github.com/gorilla/mux"
)

func main() {
	// repositories
	contactRepo := repository.NewContact()

	// usecases
	contactUsecase := usecase.NewContact(contactRepo)

	// controllers
	muxRouter := mux.NewRouter()
	contactController := controller.NewContact(contactUsecase)

	httpRouter := controller.Setup(
		contactController,
		muxRouter,
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
