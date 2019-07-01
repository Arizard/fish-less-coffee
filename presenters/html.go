package presenters

import (
	"bytes"
	"text/template"
	"fmt"
)

// HTMLPresenter presents data in browser-renderable html.
type HTMLPresenter struct {}

func (HTMLPresenter) Index() string {
	return "Index."
}

// GetUserFile returns the HTML pretty-formatted user file view.
func (HTMLPresenter) GetUserFile(publicURL string) string {
	var buf bytes.Buffer
	data := struct {
		PublicURL string
	}{
		PublicURL: publicURL,
	}
	t, err := template.ParseFiles("presenters/templates/getuserfile.tmpl")
	if err != nil {
		fmt.Printf("Error: %s", err)
		return "Error 1"
	}
	if err := t.Execute(&buf, data); err != nil {
		fmt.Printf("Error: %s", err)
		return "Error 2"
	}
	return buf.String()
}