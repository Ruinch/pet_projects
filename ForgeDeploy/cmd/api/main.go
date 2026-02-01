package main

import (
	"log"
	"net/http"
	"os"

	"forgedeploy/internal/api"
	"forgedeploy/internal/api/handlers"
	"forgedeploy/internal/store"
)

func main() {
	db := store.NewPostgres(os.Getenv("POSTGRES_DSN"))
	repo := store.NewPipelineRepoPostgres(db)

	handler := handlers.NewPipelineHandler(repo)
	router := api.NewRouter(handler)

	log.Println("API listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
