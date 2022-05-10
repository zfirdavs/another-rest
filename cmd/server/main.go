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
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	projectID := os.Getenv("FIRESTORE_PROJECT_ID")
	collectionName := os.Getenv("FIRESTORE_COLLECTON_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	repo := repository.NewFirestoreRepository(projectID, collectionName)
	s := service.NewPost(repo)
	handler := controller.NewHandler(s)

	r := router.NewGorillaMux()
	// Another implementation of our Router interface
	// r := router.NewChiMux()

	r.Get("/posts", handler.GetPosts)
	r.Post("/posts", handler.AddPost)

	r.Serve(serverPort)
}
