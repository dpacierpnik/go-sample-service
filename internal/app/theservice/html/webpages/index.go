package webpages

import (
	"html/template"
	"net/http"
	"strings"
)

type IndexContext struct {
	Name string
}

func NewIndex(templatesDir string) http.HandlerFunc {

	var indexTemplate = template.Must(template.New("index").ParseFiles(templatesDir+"layout.html", templatesDir+"index.html"))

	return func (w http.ResponseWriter, r * http.Request) {

		context := IndexContext{}

		name, ok := r.URL.Query()["name"]
		if ok && (len(name[0]) > 0) {
			context.Name = strings.Title(name[0])
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		if err := indexTemplate.ExecuteTemplate(w, "page", context); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
