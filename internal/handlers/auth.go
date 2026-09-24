package handlers

import "net/http"

type AuthHandler struct {
	hashedPasswords map[string]string 
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) Authorization(w http.ResponseWriter, r *http.Request) {
	user, password, ok := r.BasicAuth()	
	if !ok {
		http.Error(w, "Unauthorizes", http.StatusUnauthorized)
		return
	}
	if hashedPassword, ok := h.hashedPasswords[user]; !ok {
		http.Error(w, "User not foung", http.StatusNotFound)
		return
	} else if ok {
		if password == hashedPassword {
			w.Write([]byte("Успешная авторизация!"))
		} else {
			http.Error(w, "Forbidden", http.StatusUnauthorized)
		}
	}
}