package presenters

import (

)

// Presenter defines the contract for presenters, either html or json.
type Presenter interface {
	NotFound() string
	InternalServerError() string
	Index() string
	GetUserFile(fileName string, publicURL string) string
}
