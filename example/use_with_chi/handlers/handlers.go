package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	service "example/use_with_chi"
)

func init() {
	service.ChiMux.MustGet().Add()

	go func() {
		service.ChiMux.MustGet().Capture(func(mux *chi.Mux) error {
			mux.Mount("/handlers", handlers())
			return nil
		})
		service.ChiMux.MustGet().Done()
	}()
}

func handlers() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})

	return r
}
