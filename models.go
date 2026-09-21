package main

type ShortenRequest struct {
	URL        string `json:"url"`
	CustomCode string `json:"customCode"`
}

type ShortenResponse struct {
	ShortCode string `json:"shortCode"`
}

