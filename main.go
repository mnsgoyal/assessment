package main

import (
	"context"
	"emp-details/employee"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

const (
	serverPort = "localhost:8080"
)

func init() {
	log.Info("Loading database with default employee data")
	defaultEmpList := employee.GetDefaultEmployeeList()
	for _, v := range defaultEmpList {
		employee.EmpIDCounter++
		employee.EmpDetailsDB[v.ID] = v

	}

}

func main() {
	handler := setUpServer()
	srv := &http.Server{Addr: serverPort, Handler: handler}
	go func() {
		log.Info("Starting server")

		err := srv.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				log.Info("Server shut down. Waiting for connections to drain.")
			} else {
				log.WithError(err).
					WithField("server_port", srv.Addr).
					Fatal("failed to start server")
			}
		}
	}()

	// Wait for an interrupt
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)    // interrupt signal sent from terminal
	signal.Notify(sigint, syscall.SIGTERM) // sigterm signal sent from system
	<-sigint

	log.Info("Shutting down server")

	attemptGracefulShutdown(srv)
}

func setUpServer() *mux.Router {

	r := mux.NewRouter()

	//Methods
	r.HandleFunc("/employee/{ID}", employee.GetEmployeeDetailHandler).Methods("GET")
	r.HandleFunc("/employee/{page}/{limit}", employee.ListEmployeeDetailHandler).Methods("GET")
	r.HandleFunc("/employee", employee.AddEmployeeDetailHandler).Methods("POST")
	r.HandleFunc("/employee", employee.UpdateEmployeeDetailHandler).Methods("PATCH")
	r.HandleFunc("/employee/{ID}", employee.DeleteEmployeeDetailHandler).Methods("DELETE")

	return r
}

func attemptGracefulShutdown(srv *http.Server) {
	if err := shutdownServer(srv, 25*time.Second); err != nil {
		log.WithError(err).Error("failed to shutdown server")
	}
}

func shutdownServer(srv *http.Server, maximumTime time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), maximumTime)
	defer cancel()
	return srv.Shutdown(ctx)
}
