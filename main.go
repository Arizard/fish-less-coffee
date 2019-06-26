package main

// GOOGLE_APPLICATION_CREDENTIALS="/Users/arie/go/src/github.com/arizard/fish-less-coffee/arie-personal-arie-images-service.json"

import (
	"context"
	"fmt"
	"cloud.google.com/go/storage"
	"io/ioutil"
	"github.com/arizard/fish-less-coffee/entities"
	"github.com/arizard/fish-less-coffee/infrastructure"
)

func main() {
	fmt.Printf("Setting up infrastructure...\n")

	ctx := context.Background()
	bucketName := "arie-images"
	client, err := storage.NewClient(ctx)

	bucket := client.Bucket(bucketName)

	userFileRepo := infrastructure.GCSUserFileRepository{
		Context: ctx,
		Bucket: bucket,
	}

	fmt.Printf("Setting up services...\n")

	userFileService := entities.UserFileService{
		Repository: userFileRepo,
	}

	fmt.Printf("Adding a new file to GCS...\n")

	fileName := "test-image-2.png"

	imgBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}

	newUserFile := userFileService.NewUserFile(fileName, imgBytes)

	userFileRepo.Add(newUserFile)

	fmt.Printf(
		"Link: https://storage.googleapis.com/%s/%s\n",
		bucketName,
		fileName,
	)

}