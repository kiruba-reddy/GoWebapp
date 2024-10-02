package renderers

import (
	"html/template"
	"net/http"
)

func renderPage(w http.ResponseWriter, path string) {
	parsedHtml, _ := template.ParseFiles("./templates/" + path)
	_ = parsedHtml.Execute(w, nil)
}