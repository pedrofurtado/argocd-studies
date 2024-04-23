package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "My ArgoCD studies app v5")
	})

	http.ListenAndServe(":8080", nil)
}
