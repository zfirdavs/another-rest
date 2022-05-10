package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type chiRouter struct {
	m *chi.Mux
}

func NewChiMux() *chiRouter {
	return &chiRouter{m: chi.NewRouter()}
}

func (r *chiRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.m.Get(uri, f)
}

func (r *chiRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.m.Post(uri, f)
}

func (r *chiRouter) Serve(port string) {
	log.Println("Chi http server listening on port", port)
	http.ListenAndServe(port, r.m)
}
