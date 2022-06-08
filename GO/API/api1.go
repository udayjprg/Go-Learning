package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main(){
	//just test1 in main file
	router :=mux.NewRouter()
	router.HandleFunc("/submitInvoice", InvoiceHandler).Methods(http.MethodGet)
}
