package controller

import (
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPassword := "password"

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
	}

}
