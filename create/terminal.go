package main

import (
	"encoding/json"
	"net/http"
)

func createTerminal(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	token, err := ValidateJWT(tokenString)
	if err != nil || !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Terminal created successfully"})
}
