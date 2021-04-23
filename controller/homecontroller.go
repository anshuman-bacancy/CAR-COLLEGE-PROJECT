package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"net/smtp"
	"project/services"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

var (
	admintpl *template.Template
	hometpl  *template.Template
	custtpl  *template.Template
)

func init() {
	admintpl = template.Must(template.New("").Funcs(Fm).ParseGlob(("./template/admin/*")))
	hometpl = template.Must(template.New("").Funcs(Fm).ParseGlob(("./template/home/*")))
	custtpl = template.Must(template.New("").Funcs(Fm).ParseGlob(("./template/customer/*")))
}

var registeremail bool
var customernotexits bool
var storecustomer = sessions.NewCookieStore([]byte("t0p-s3cr3tcus"))
var emailfound bool
var emailnotextits bool

// home page
func HomePage(w http.ResponseWriter, r *http.Request) {
	hometpl.ExecuteTemplate(w, "index.html", nil)
}

// view customer registration
func CustomerRegister(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if registeremail {
		hasmessge = true
		message = "Email is already taken"
		registeremail = false
	}
	hometpl.ExecuteTemplate(w, "Register.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

// successful customer registration
func CustomerRegisterPOST(w http.ResponseWriter, r *http.Request) {
	emailcheck, err := services.SaveCustomer(r)
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

// view customer login
func CustomerLogin(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if customernotexits {
		hasmessge = true
		message = "Wrong Username or Password"
		customernotexits = false
	}
	hometpl.ExecuteTemplate(w, "login.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

// successful customer login
func CustomerLoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	usercustomer := r.PostForm.Get("username")
	passcustomer := r.PostForm.Get("password")
	customers := services.GetAllCustomer(r)
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

// customer index page
func CustomerIndexPage(w http.ResponseWriter, r *http.Request) {
	brand := services.GetAllBrand(r)
	// fmt.Println(brand)
	custtpl.ExecuteTemplate(w, "index.html", brand)
}

// customer authentication
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

// customer logout
func CustomerLogout(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	session.Options.MaxAge = -1
	delete(session.Values, "customer")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// customer forgot password
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
		message = "Email does not exists!"
		emailnotextits = false
	}

	hometpl.ExecuteTemplate(w, "forgotpassword.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

// customer validate email
func CustomerValidateEmail(w http.ResponseWriter, r *http.Request) {
	emailfound = false
	email := r.FormValue("email")
	customers := services.GetAllCustomer(r)
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
	// from := "autogradingsystem99999@gmail.com"
	from := "anshumanaich99@gmail.com"
	password := "anshumanaich32"
	to := []string{
		email,
	}
	Receivermail := email

	customer := services.GetOneCustomerBYemail(email)
	customerid := strconv.Itoa(int(customer.ID))
	subjet := "FORGOT YOUR PASSSWORD EMAIL"
	body := "http://localhost:8084/customer/setpassword/" + customerid
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

// customer set new password
func CustomerSetForgotPasswordPage(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	id := mux.Vars(r)["id"]
	email, ok := session.Values["emailid"]
	if !ok {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	customer := services.GetOneCustomerBYemail(email)
	if email != customer.Email {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	hometpl.ExecuteTemplate(w, "setforgotpassword.html", id)
}

// customer successful login
func CustomerSuccess(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	delete(session.Values, "emailid")
	session.Save(r, w)
	hometpl.ExecuteTemplate(w, "success.html", nil)
}
