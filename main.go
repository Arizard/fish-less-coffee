package main

// GOOGLE_APPLICATION_CREDENTIALS="/Users/arie/go/src/github.com/arizard/fish-less-coffee/arie-personal-arie-images-service.json"

import (
	"context"
	"fmt"
	"cloud.google.com/go/storage"
	// "io/ioutil"
	"github.com/arizard/fish-less-coffee/infrastructure"
	"github.com/arizard/fish-less-coffee/presenters"
	"github.com/arizard/fish-less-coffee/handlers"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Printf("Setting up infrastructure...\n")

	ctx := context.Background()
	bucketName := "arie-images"
	client, _ := storage.NewClient(ctx)

	bucket := client.Bucket(bucketName)

	userFileRepo := infrastructure.GCSUserFileRepository{
		Context: ctx,
		Bucket: bucket,
	}

	fmt.Printf("Setting up layers...\n")

	htmlPresenter := presenters.HTMLPresenter{}

	r := mux.NewRouter().StrictSlash(false)
	HTMLHandler := handlers.Handler{
		userFileRepo,
		htmlPresenter,
	}

	r.HandleFunc("/", HTMLHandler.Index).Methods("GET")

	r.HandleFunc("/look/{name}", HTMLHandler.GetPublicURL).Methods("GET")
	r.HandleFunc("/give", HTMLHandler.UploadUserFile).Methods("POST")

	http.ListenAndServe(":8080", r)

}