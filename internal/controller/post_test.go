package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zfirdavs/another-rest/internal/entity"
	"github.com/zfirdavs/another-rest/internal/repository"
	"github.com/zfirdavs/another-rest/internal/service"
)

const (
	postID    int64  = 123
	postTitle string = "Title 1"
	postText  string = "text 1"
)

func TestAddPost(t *testing.T) {
	// create post request
	jsonRequest := []byte(`{"title": "Title 1", "text": "text 1"}`)
	request := httptest.NewRequest("POST", "/posts", bytes.NewBuffer(jsonRequest))

	// assign http handler

	postRepo := repository.NewSQLiteRepository("posts.db")

	err := postRepo.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = postRepo.CreatePostsTable()
	if err != nil {
		log.Fatal(err)
	}

	postService := service.NewPost(postRepo)
	postController := NewHandler(postService)
	handler := http.HandlerFunc(postController.AddPost)

	// record http response (httptest)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, request)

	status := w.Code

	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v, want: %v", status, http.StatusOK)
	}

	// decode response body
	post := entity.Post{}
	json.NewDecoder(w.Body).Decode(&post)

	// assertions
	assert.NotNil(t, post.ID)
	assert.Equal(t, postTitle, post.Title)
	assert.Equal(t, postText, post.Text)

}

func TestGetPosts(t *testing.T) {
	request := httptest.NewRequest("GET", "/posts", nil)

	postRepo := repository.NewSQLiteRepository("posts.db")

	err := postRepo.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = postRepo.CreatePostsTable()
	if err != nil {
		log.Fatal(err)
	}

	// create a post
	post := entity.Post{
		ID:    postID,
		Title: postTitle,
		Text:  postText,
	}
	postRepo.Save(context.Background(), &post)

	// assign
	postService := service.NewPost(postRepo)
	postController := NewHandler(postService)
	handler := http.HandlerFunc(postController.GetPosts)

	// record http response (httptest)
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, request)

	status := w.Code

	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v, want: %v", status, http.StatusOK)
	}

	var posts []entity.Post
	json.NewDecoder(w.Body).Decode(&posts)

	// assertions
	assert.NotNil(t, post.ID)
	assert.Equal(t, postTitle, posts[0].Title)
	assert.Equal(t, postText, posts[0].Text)
}
