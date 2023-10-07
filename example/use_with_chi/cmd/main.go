package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	service "example/use_with_chi"
	_ "example/use_with_chi/handlers"
)

func main() {
	go func() {
		service.ChiMux.MustGet().Wait()

		service.ChiMux.MustGet().Capture(func(mux *chi.Mux) error {
			log.Fatal(http.ListenAndServe(":8080", mux))
			return nil
		})
	}()

	select {}
}
