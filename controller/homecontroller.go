package controller

import (
	"fmt"
	"go/build"
	"net/http"
	"net/smtp"
	"project/data/service"
	"strconv"
	"text/template"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var registeremail bool
var customernotexits bool
var storecustomer = sessions.NewCookieStore([]byte("t0p-s3cr3tcus"))
var emailfound bool
var emailnotextits bool

//HomePage is....
func HomePage(w http.ResponseWriter, r *http.Request) {
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
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
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
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
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
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
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	brand := service.GetAllBrand(r)
	tpl.ExecuteTemplate(w, "index.html", brand)
}

//AuthenticationCustomer is..
func AuthenticationCustomer(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := storecustomer.Get(r, "customerusername")
		_, ok := session.Values["customer"]
		if !ok {

			fmt.Println("not exits")
			http.Redirect(w, r, "/Login", http.StatusSeeOther)
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

//CustomerForgotPassword is..
func CustomerForgotPassword(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if emailfound {
		hasmessge = true
		message = "Check Your Email"
		emailfound = false
	}

	if emailnotextits {
		hasmessge = true
		message = "Email is does not Exits"
		emailnotextits = false
	}

	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	tpl.ExecuteTemplate(w, "forgotpassword.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

//CustomerValidateEmail is...
func CustomerValidateEmail(w http.ResponseWriter, r *http.Request) {
	emailfound = false
	email := r.FormValue("email")
	customers := service.GetAllCustomer(r)
	for _, customer := range customers {
		if customer.Email == email {
			emailfound = true
			break
		}
	}
	if emailfound {
		session, _ := storecustomer.Get(r, "customerusername")
		session.Values["emailid"] = email
		session.Save(r, w)
		sendemail(email)
		http.Redirect(w, r, "/customer/forgotpassword", http.StatusSeeOther)
		return
	}
	emailnotextits = true
	http.Redirect(w, r, "/customer/forgotpassword", http.StatusSeeOther)
}

func sendemail(email string) {
	// Sender data.

	from := "autogradingsystem99999@gmail.com"
	password := "nedlsjaxafqmlnms"
	to := []string{
		email,
	}
	// Receiver email address.
	Receivermail := email

	customer := service.GetOneCustomerBYemail(email)
	customerid := strconv.Itoa(int(customer.ID))
	subjet := "FORGOT YOUR PASSSWORD EMAIL"
	body := "http://localhost:8084/customer/setpassword/" + customerid
	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("To: " + Receivermail + "\r\n" +
		"Subject: " + subjet + "!\r\n" +
		"\r\n" + "Set your forgot password from below link\r\n" +
		body + "\r\n")

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}

//CustomerSetForgotPasswordPage is...
func CustomerSetForgotPasswordPage(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	id := mux.Vars(r)["id"]
	email, ok := session.Values["emailid"]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	customer := service.GetOneCustomerBYemail(email)
	if email != customer.Email {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	tpl.ExecuteTemplate(w, "setforgotpassword.html", id)
}

//CustomerSuccess is...
func CustomerSuccess(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	delete(session.Values, "emailid")
	session.Save(r, w)
	path := build.Default.GOPATH + "/src/project/template/home/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	tpl.ExecuteTemplate(w, "success.html", nil)
}
