package controller

import (
	"fmt"
	"go/build"
	"net/http"
	"project/data/service"
	"text/template"

	"github.com/gorilla/sessions"
)

var registeremail bool
var customernotexits bool
var storecustomer = sessions.NewCookieStore([]byte("t0p-s3cr3tcus"))

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
	var message string
	var hasmessge bool
	if customernotexits {
		hasmessge = true
		message = "Username or Password is wrong"
		customernotexits = false
	}
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.ParseGlob(path))
	tpl.ExecuteTemplate(w, "login.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

//CustomerLoginPost is...
func CustomerLoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usercustomer := r.PostForm.Get("username")
	passcustomer := r.PostForm.Get("password")
	customers := service.GetAllCustomer(r)
	fmt.Println(customers)
	if len(customers) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var access bool
	for _, customer := range customers {
		if customer.Email == usercustomer && customer.Password == passcustomer {
			access = true
			session, _ := storecustomer.Get(r, "customerusername")
			session.Values["customer"] = customer.Email
			session.Save(r, w)
			access = true
			break
		}
	}
	if !access {
		customernotexits = true
		http.Redirect(w, r, "/Login", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/customer/index", http.StatusSeeOther)
}

//CustomerIndexPage is...
func CustomerIndexPage(w http.ResponseWriter, r *http.Request) {
	path := build.Default.GOPATH + "/src/project/template/customer/*"
	tpl := template.Must(template.ParseGlob(path))
	vehicles := service.GetAllVehicle()
	tpl.ExecuteTemplate(w, "index.html", vehicles)
}

//AuthenticationCustomer is..
func AuthenticationCustomer(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := storecustomer.Get(r, "customerusername")
		_, ok := session.Values["customer"]
		if !ok {

			fmt.Println("not exits")
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

//CustomerLogout is....
func CustomerLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	session.Options.MaxAge = -1
	delete(session.Values, "customer")
	session.Save(r, w)
	//fmt.Fprintf(w, "username is cleared")
	http.Redirect(w, r, "/", http.StatusSeeOther)
	//return
}
