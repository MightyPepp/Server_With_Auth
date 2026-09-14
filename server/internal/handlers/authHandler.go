package handlers

import "net/http"

func (h *myHandler) AuthHandler(w http.ResponseWriter, r *http.Request) {
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