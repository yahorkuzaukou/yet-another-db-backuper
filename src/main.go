package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)
	http.ListenAndServe(":3000", mux)
}
