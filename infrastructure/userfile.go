package infrastructure

import (
	"fmt"
	"github.com/arizard/fish-less-coffee/entities"
	"cloud.google.com/go/storage"
	"context"
	"github.com/icrowley/fake"
	"github.com/gosimple/slug"
	"path/filepath"
)

// GCSUserFileRepository implements the repository model using google
// cloud storage
type GCSUserFileRepository struct {
	Context context.Context
	Bucket  *storage.BucketHandle
}

// Add submits a new UserFile into the repository for persistence.
func (repo GCSUserFileRepository) Add(userFile *entities.UserFile) {
	ref := fmt.Sprintf(
		"%s %s the %s",
		fake.FirstName(),
		fake.LastName(),
		fake.ProductName(),
	)
	newName := slug.Make(ref) + filepath.Ext(userFile.Name)

	userFile.Name = newName

	obj := repo.Bucket.Object(userFile.Name)
	w := obj.NewWriter(repo.Context)
	defer w.Close()

	w.Write(userFile.Data)

}

// Get retrieves the UserFile entity of a UserFile Name.
func (repo GCSUserFileRepository) Get(name string) entities.UserFile {
	return entities.UserFile{}
}

// GetPublicURL returns the public access URL of a UserFile.
func (repo GCSUserFileRepository) GetPublicURL(name string) string {
	
	attrs, err := repo.Bucket.Attrs(repo.Context)
	if err != nil {
		fmt.Printf("%s\n", err)
		return ""
	}
	return fmt.Sprintf(
		"https://storage.googleapis.com/%s/%s",
		attrs.Name,
		name,
	)
}
