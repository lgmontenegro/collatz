package server

import (
	"fmt"
	"lgmontenegro/collantz/internal/service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server(port int) {
	portToServe := fmt.Sprintf(":%d", port)
	r := mux.NewRouter()

	r.HandleFunc("/collatz/{factor}", handlers.CollatzHandler).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(portToServe, r))
	return
}
