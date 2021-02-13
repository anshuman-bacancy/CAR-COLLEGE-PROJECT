package controller

import (
	"go/build"
	"net/http"
	"text/template"
)

//HomePage is....
func HomePage(w http.ResponseWriter, r *http.Request) {
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.ParseGlob(path))
	tpl.ExecuteTemplate(w, "index.html", nil)
}
