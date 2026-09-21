package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
)

var urls = make(map[string]string)

type ShortenRequest struct {
	URL string `json:"url"`
}

type ShortenResponse struct{
	ShortCode string `json:"shortCode"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Url Shortner")
}

func shortenHandler(w http.ResponseWriter, r *http.Request){
	if r.Method == "POST" {
		var request ShortenRequest
		json.NewDecoder(r.Body).Decode(&request)
		shortCode:= shortenURL(request.URL);

		var response ShortenResponse
		response.ShortCode= shortCode
		json.NewEncoder(w).Encode(&response)
	}
}
func generateCode() string {
	number := rand.Intn(1000000)
	code := fmt.Sprint(number)

	_, exists := urls[code]

	if exists {
		return generateCode()
	}

	return code
}
func shortenURL(url string) string {
	code := generateCode()

	urls[code] = url

	return code
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/shorten", shortenHandler)

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
	shortCode1 := shortenURL("https://github.com/aderogbasamuel")
	shortCode2 := shortenURL("https://google.com")

	fmt.Println(shortCode1)
	fmt.Println(urls[shortCode1])
	fmt.Println(shortCode2)
	fmt.Println(urls[shortCode2])
}
