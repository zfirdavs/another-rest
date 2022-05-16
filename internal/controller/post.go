package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/zfirdavs/another-rest/internal/entity"
	"github.com/zfirdavs/another-rest/internal/service"
	"github.com/zfirdavs/another-rest/pkg/errors"
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	AddPost(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service service.PostService
}

func NewHandler(service service.PostService) *handler {
	return &handler{service}
}

func (h *handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.service.FindAll(r.Context())
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error getting posts"})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (h *handler) AddPost(w http.ResponseWriter, r *http.Request) {
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error unmarshalling post"})
		return
	}

	ctx := r.Context()

	err = h.service.Validate(ctx, &post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err.Error()})
		return
	}

	err = h.service.Create(ctx, &post)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "error creating post"})
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}
