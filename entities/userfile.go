package entities

import (
	
)

// UserFile is the entity which describes a user-generated file.
type UserFile struct {
	Name string // Name must be unique.
	Data []byte
}

// UserFileService manages business logic related to 
// instances of UserFile.
type UserFileService struct {
	Repository UserFileRepository
}

// NewUserFile creates a new UserFile instance.
func (s *UserFileService) NewUserFile(name string, data []byte) UserFile {
	return UserFile{
		Name: name,
		Data: data,
	}
}

// UserFileRepository manages the storage of UserFile instances.
type UserFileRepository interface {
	Add(userFile *UserFile)
	Get(name string) UserFile
	GetPublicURL(name string) string
}
