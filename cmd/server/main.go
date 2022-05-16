package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/zfirdavs/another-rest/internal/controller"
	"github.com/zfirdavs/another-rest/internal/repository"
	"github.com/zfirdavs/another-rest/internal/service"
	"github.com/zfirdavs/another-rest/pkg/http/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	// projectID := os.Getenv("FIRESTORE_PROJECT_ID")
	// collectionName := os.Getenv("FIRESTORE_COLLECTON_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	// repo := repository.NewFirestoreRepository(projectID, collectionName)
	repo := repository.NewSQLiteRepository("posts.db")
	defer repo.Close()

	err = repo.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = repo.CreatePostsTable()
	if err != nil {
		log.Fatal(err)
	}

	s := service.NewPost(repo)
	handler := controller.NewHandler(s)

	r := router.NewGorillaMux()
	// Another implementation of our Router interface
	// r := router.NewChiMux()

	r.Get("/posts", handler.GetPosts)
	r.Post("/posts", handler.AddPost)

	r.Serve(serverPort)
}
