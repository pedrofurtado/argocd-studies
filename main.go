package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "My ArgoCD app")
	})

	http.ListenAndServe(":8080", nil)
}
