package service

import (
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/dkotTech/startfast"
)

var ChiMux = startfast.NewLazy(initMux)

func initMux() (*tmux, error) {
	m := chi.NewRouter()

	initMiddlewares(m)

	return &tmux{mux: m}, nil
}

func initMiddlewares(mux *chi.Mux) {
	mux.Use(
		middleware.Heartbeat("/alive"),
		middleware.Recoverer,
	)
}

type tmux struct {
	wg  sync.WaitGroup
	m   sync.Mutex
	mux *chi.Mux
}

func (t *tmux) Capture(f func(*chi.Mux) error) {
	t.m.Lock()
	defer t.m.Unlock()

	err := f(t.mux)
	if err != nil {
		panic(err)
	}

}

func (t *tmux) Add() {
	t.wg.Add(1)
}

func (t *tmux) Done() {
	t.wg.Done()
}

func (t *tmux) Wait() {
	t.wg.Wait()
}
