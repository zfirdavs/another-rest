package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type gorillaMuxRouter struct {
	m *mux.Router
}

func NewGorillaMux() *gorillaMuxRouter {
	return &gorillaMuxRouter{m: mux.NewRouter()}
}

func (r *gorillaMuxRouter) Get(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.m.HandleFunc(uri, f).Methods("GET")
}

func (r *gorillaMuxRouter) Post(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.m.HandleFunc(uri, f).Methods("POST")
}

func (r *gorillaMuxRouter) Serve(port string) {
	log.Println("Server listening on port", port)
	http.ListenAndServe(port, r.m)
}
