package handlers

import (
	// "encoding/json"
	"github.com/arizard/fish-less-coffee/presenters"
	"github.com/arizard/fish-less-coffee/entities"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"github.com/arizard/fish-less-coffee/usecases"
)

// Handler is a struct which implements methods that take the 
// ResponseWriter and Request objects as arguments, such as from an
// HTTP request. It is used to decouple the Drivers layer from the
// Controllers and Presenters.
type Handler struct {
	UserFileRepo entities.UserFileRepository
	Presenter presenters.Presenter
}

func (handler Handler) errorHelper(
	w http.ResponseWriter,
	r *http.Request,
	rc usecases.ResponseCollector,
) {
	if rc.Error != nil {
		if rc.Error.Name == "NOT_FOUND" {
			w.WriteHeader(404)
			handler.NotFound(w, r)
		}
		if rc.Error.Name == "SEVERE_FAILURE" {
			w.WriteHeader(500)
			handler.InternalServerErrorHandler(w, r)
		}
	}
}

//NotFoundHandler handles 404s
func (handler Handler) NotFound(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.Presenter.NotFound())
}

//InternalServerErrorHandler handles 500s
func (handler Handler) InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.Presenter.InternalServerError())
}

// Index handles a request for the Index view of the presenter.
func (handler Handler) Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.Presenter.Index())
}

// GetPublicURL handles the GetPublicURL view of the presenter.
func (handler Handler) GetPublicURL(w http.ResponseWriter, r *http.Request) {
	fileName := mux.Vars(r)["name"]

	rc := usecases.ResponseCollector{}
	uc := usecases.ViewUserFile{
		FileName: fileName,
		UserFileRepo: handler.UserFileRepo,
		Presenter: handler.Presenter,
		Response: &rc,
	}

	uc.Setup()
	uc.Execute()
	
	if rc.Error != nil {
		if rc.Error.Name == "NOT_FOUND" {
			w.WriteHeader(404)
			handler.NotFound(w, r)
		}
		if rc.Error.Name == "SEVERE_FAILURE" {
			w.WriteHeader(500)
			handler.InternalServerErrorHandler(w, r)
		}
		return
	}
	fmt.Fprintf(w, "%s", rc.Response.Body)
}


func (handler Handler) UploadUserFile(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	file, fileHeader, _ := r.FormFile("file")

	rc := usecases.ResponseCollector{}
	uc := usecases.CreateUserFile{
		File: file,
		FileHeader: fileHeader,
		UserFileRepo: handler.UserFileRepo,
		Response: &rc,
	}

	uc.Setup()

	uc.Execute()

	w.Header().Set("Content-Location", fmt.Sprintf("/look/%s", rc.Response.Body))
	http.Redirect(w, r, fmt.Sprintf("/look/%s", rc.Response.Body), 301)

	handler.errorHelper(w, r, rc)
}