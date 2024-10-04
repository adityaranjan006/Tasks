package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var sso *oauth2.Config
var RandomVariable = "Ug8neo0wtc"

func init() {
	err := godotenv.Load("./.env")

	if err != nil {
		log.Fatal("Error Loading environment variables.")
	}
	sso = &oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURL:  os.Getenv("REDIRECT_URI"),
		Endpoint:     google.Endpoint,
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
}
func main() {
	http.HandleFunc("/auth/dreamsleep", handleAuth)
	fmt.Println("Authorization server listening on http://localhost:5000/auth/dreamsleep")
	http.ListenAndServe(":5000", nil)
}

func handleAuth(w http.ResponseWriter, r *http.Request) {
	accessToken := r.Header.Get("Authorization")
	if accessToken == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}

	if len(accessToken) < 7 || accessToken[:7] != "Bearer " {
		http.Error(w, "Invalid authorization format", http.StatusUnauthorized)
		return
	}
	tokenString := accessToken[7:]
	fmt.Println("WORKING HANDLER", tokenString)

	client := sso.Client(context.Background(), &oauth2.Token{AccessToken: tokenString})
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")

	if err != nil {
		fmt.Println("WRONG RESPONSE")
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		fmt.Printf("failed to get user info: %s", response.Status)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
