package main

import (
	"context"
	"net/http"

	"website/components"
)

var ctx = context.Background()

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		components.Page().Render(ctx, w)
	})
	http.ListenAndServe(":3000", mux)
}
