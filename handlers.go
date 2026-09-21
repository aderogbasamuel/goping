package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)

type Handler struct{
	storage Storage
}
func (h Handler) shortenHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var request ShortenRequest
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			json.NewEncoder(w).Encode(err)
			return
		}
		if request.URL == "" {
			http.Error(w, "URL is required", http.StatusBadRequest)
			return
		}
		if request.CustomCode == "" {
			http.Error(w, "Provide a custom Code", http.StatusBadRequest)
			return
		}
		_, getErr := h.storage.Get(request.CustomCode)

		if getErr== nil {
			http.Error(w, "Code has been used already", http.StatusConflict)
			return
		}
		
		// shortCode := shortenURL(request.URL)
		h.storage.Save(request.CustomCode, request.URL)
		var response ShortenResponse
		response.ShortCode = request.CustomCode
		json.NewEncoder(w).Encode(&response)
	}
}

func (h Handler) redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		shortCode := r.URL.Path

		if shortCode == "/" {
			fmt.Fprintln(w, "Url Shortner api running")
			return
		}
		shortCodeTrimmed := strings.TrimPrefix(shortCode, "/")

		originalURL, err := h.storage.Get(shortCodeTrimmed)

		if err!=nil {
			http.NotFound(w, r)
			return
		}
		http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
	}
}