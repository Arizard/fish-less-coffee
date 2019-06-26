package infrastructure

import (
	"github.com/arizard/fish-less-coffee/entities"
	"cloud.google.com/go/storage"
	"context"
)

// GCSUserFileRepository implements the repository model using google
// cloud storage
type GCSUserFileRepository struct {
	Context context.Context
	Bucket  *storage.BucketHandle
}

// Add submits a new UserFile into the repository for persistence.
func (repo GCSUserFileRepository) Add(userFile entities.UserFile) {
	obj := repo.Bucket.Object(userFile.Name)
	w := obj.NewWriter(repo.Context)
	defer w.Close()

	w.Write(userFile.Data)
}

// Get retrieves the UserFile entity of a UserFile Name
func (repo GCSUserFileRepository) Get(name string) entities.UserFile {
	return entities.UserFile{}
}