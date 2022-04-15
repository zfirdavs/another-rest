package route

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"

	"github.com/zfirdavs/another-rest/internal/entity"
	"github.com/zfirdavs/another-rest/internal/repository"
)

type handler struct {
	repo repository.PostRepository
}

func NewHandler(repo repository.PostRepository) *handler {
	return &handler{repo}
}

func (h *handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := h.repo.FindAll(r.Context())
	if err != nil {
		http.Error(w, "error getting posts", http.StatusInternalServerError)
		log.Println(err)
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
		http.Error(w, "error decoding data", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	post.ID = rand.Int63()

	err = h.repo.Save(r.Context(), &post)
	if err != nil {
		http.Error(w, "error saving posts", http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)

}
