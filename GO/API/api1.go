package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/submitInvoice", InvoiceHandler).Methods(http.MethodGet)
}
