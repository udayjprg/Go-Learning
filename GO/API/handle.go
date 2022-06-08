package main

import "net/http"

func InvoiceHandler(resp http.ResponseWriter, req *http.Request){
	resp.Write([]byte("Hello World"))
}