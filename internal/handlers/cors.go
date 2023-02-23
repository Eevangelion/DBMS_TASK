package handlers

import "net/http"

func setupCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func OptionsHandler(w http.ResponseWriter, r *http.Request) {
	setupCors(&w)
	w.WriteHeader(http.StatusOK)
	return
}
