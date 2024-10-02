package main

import (
	"net/http"
	"github.com/kiruba-reddy/GoWebapp"
)

const portNo = ":8080"


func main() {
	http.HandleFunc("/", Home)
	_ = http.ListenAndServe(portNo, nil)
}
