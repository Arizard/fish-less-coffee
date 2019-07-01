package usecases

import (

)

// UseCase is an interface which standardises the methods available on
// Use Cases.
type UseCase interface {
	Setup()
	Execute()
}

