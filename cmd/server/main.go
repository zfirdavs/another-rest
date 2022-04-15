package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/zfirdavs/another-rest/internal/repository"
	"github.com/zfirdavs/another-rest/internal/route"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	projectID := os.Getenv("FIRESTORE_PROJECT_ID")
	collectionName := os.Getenv("FIRESTORE_COLLECTON_NAME")
	serverPort := os.Getenv("SERVER_PORT")

	repo := repository.NewPostRepository(projectID, collectionName)
	handler := route.NewHandler(repo)
	router := mux.NewRouter()

	router.HandleFunc("/posts", handler.GetPosts).Methods("GET")
	router.HandleFunc("/posts", handler.AddPost).Methods("POST")

	log.Println("Server listening on port", serverPort)
	log.Fatalln(http.ListenAndServe(serverPort, router))

}
