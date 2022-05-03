package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	http.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
		res, err := httputil.DumpRequest(request, true)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
		fmt.Fprintf(writer, string(res))
	})

	http.ListenAndServe(":8080", nil)
}
