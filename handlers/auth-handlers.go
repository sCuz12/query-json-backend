package handlers

import (
	"encoding/json"
	"fmt"
	"jql-server/data"
	"net/http"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserRegisterRequest struct {
	Username string `validate:"required"`
	Email string `validate:"required"`
	Password string `validate:"required"`
	 
}

func (h *Handlers) Register(w http.ResponseWriter, r *http.Request){
    var req UserRegisterRequest

    err := json.NewDecoder(r.Body).Decode(&req)

    if err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
    }

    validate := validator.New(validator.WithRequiredStructEnabled())

    err = validate.Struct(&req)

    if err != nil {
        if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
        h.App.ErrorLog.Printf("Validation Error %v",err)
        http.Error(w,"Error validatation data", http.StatusBadRequest)
	}

    //hash the password 
    hashedPassword , err := bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
    
    if err != nil {
        h.App.ErrorLog.Printf("Error while encrypting password %v" ,err)
        http.Error(w,"Something went wrong", 500)
        return;
    }

    user := data.User {
        Username: req.Username,
        Email: req.Email,
        Password: string(hashedPassword),
    }

    fmt.Println(user)
    _,err = h.Models.Users.Insert(user)

    
    if err != nil {
        h.App.ErrorLog.Printf("Error while creating user %v" ,err)
        http.Error(w,"Something went wrong", 500)
        return; 
    }
	fmt.Println("Hello world")
}


// PostUserLogin attempts to log a user in
func (h *Handlers) PostUserLogin(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	email := r.Form.Get("email")
	password := r.Form.Get("password")

	user, err := h.Models.Users.GetByEmail(email)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	matches, err := user.PasswordMatches(password)
	if err != nil {
		w.Write([]byte("Error validating password"))
		return
	}

	if !matches {
		w.Write([]byte("Invalid password!"))
		return
	}



	h.App.Session.Put(r.Context(), "userID", user.ID)

	http.Redirect(w, r, "/", http.StatusSeeOther)

}
