package presenters

import (

)

// Presenter defines the contract for presenters, either html or json.
type Presenter interface {
	Index() string
	GetUserFile(publicURL string) string
}