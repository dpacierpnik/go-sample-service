package html

import (
	"html/template"
	"net/http"
	"strings"
)

const templatesDir = "html/templates/"
var indexTemplate = template.Must(template.New("index").ParseFiles(templatesDir+"layout.html", templatesDir+"index.html"))

type IndexContext struct {
	Name string
}

func Index(w http.ResponseWriter, r *http.Request) {

	context := IndexContext{}

	name, ok := r.URL.Query()["name"]
	if ok && (len(name[0]) > 0) {
		context.Name = strings.Title(name[0])
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := indexTemplate.ExecuteTemplate(w, "layout", context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
