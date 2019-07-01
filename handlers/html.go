package handlers

import (
	"encoding/json"
	"github.com/arizard/fish-less-coffee/presenters"
	"github.com/arizard/fish-less-coffee/entities"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

// Handler is a struct which implements methods that take the 
// ResponseWriter and Request objects as arguments, such as from an
// HTTP request. It is used to decouple the Drivers layer from the
// Controllers and Presenters.
type Handler struct {
	UserFileRepo entities.UserFileRepository
	Presenter presenters.Presenter
}

// Index handles a request for the Index view of the presenter.
func (handler Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.Presenter.Index())
}

// GetPublicURL handles the GetPublicURL view of the presenter.
func (handler Handler) GetPublicURL(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["name"]
	fmt.Fprintf(w, handler.Presenter.GetUserFile(
		handler.UserFileRepo.GetPublicURL(fileName),
	))
}

func (handler Handler) UploadUserFile(w http.ResponseWriter, r *http.Request) {
	var jsonReq map[string]interface{}
	json.NewDecoder(r.Body).Decode(&jsonReq)
	fmt.Printf("%x\n", jsonReq["file"])
	fmt.Fprintf(w, `{"message": "Testing"}`)
}