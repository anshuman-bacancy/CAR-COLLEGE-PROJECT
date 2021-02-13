package controller

import (
	"go/build"
	"net/http"
	"project/data/service"
	"text/template"
)

var registeremail bool

//HomePage is....
func HomePage(w http.ResponseWriter, r *http.Request) {
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.ParseGlob(path))
	tpl.ExecuteTemplate(w, "index.html", nil)
}

//CustomerRegister is...
func CustomerRegister(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if registeremail {
		hasmessge = true
		message = "Email is already taken"
		registeremail = false
	}
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.ParseGlob(path))
	tpl.ExecuteTemplate(w, "Register.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

//CustomerRegisterPOST is...
func CustomerRegisterPOST(w http.ResponseWriter, r *http.Request) {
	emailcheck, err := service.SaveCustomer(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	if emailcheck {
		registeremail = true
		http.Redirect(w, r, "/Registration", http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, "/Login", http.StatusSeeOther)
}

//CustomerLogin is...
func CustomerLogin(w http.ResponseWriter, r *http.Request) {
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.ParseGlob(path))
	tpl.ExecuteTemplate(w, "login.html", nil)
}
