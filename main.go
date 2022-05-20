package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Mukunth-arya/Ticketbooking/helpers"
	"github.com/gorilla/mux"
)

func main() {

	logger := log.New(os.Stdout, "TicketBookingApi", log.LstdFlags)

	//http.server for making server managable
	//Readtimeout for read the body and requests
	//Write Timeouts for writing the responses
	//Idle Timeouts for waiting for the new requests

	handlers_var := helpers.Datafunc(logger)
	Router := mux.NewRouter()

	Getrouter := Router.Methods(http.MethodGet).Subrouter()
	Getrouter.HandleFunc("/", handlers_var.LivenessProbe)
	Getroutevalue := Router.Methods(http.MethodGet).Subrouter()
	Getroutevalue.HandleFunc("/Locget", handlers_var.Displayget)
	Postvalue := Router.Methods(http.MethodPost).Subrouter()
	Postvalue.HandleFunc("/Insdoc", handlers_var.Dataenter)

	server := http.Server{

		Addr:         ":9090",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			os.Exit(1)
		}
	}()

	//Here we make to pass the signal to pass it to operating system
	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt)
	signal.Notify(signals, os.Kill)
	//Receive an signal to make Graceful shutdown
	Receive := <-signals

	log.Printf(" Received An Gracefull Shutdown", Receive)
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	server.Shutdown(ctx)

}
