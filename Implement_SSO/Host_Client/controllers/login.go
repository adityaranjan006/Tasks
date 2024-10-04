package controllers

import (
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

func Login(w http.ResponseWriter, r *http.Request) {
	url := sso.AuthCodeURL(RandomVariable)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)

	// 	resp, err := http.NewRequest("GET", "http://localhost:5000/auth/dreamsleep", nil)
	// 	if err != nil {
	// 		log.Fatal("Error", err)
	// 	}
	// 	client := &http.Client{}
	// 	res, err := client.Do(resp)
	// 	if err != nil {
	// 		log.Fatal("Not Validated")
	// 		return
	// 	}

	// 	data, err := io.ReadAll(res.Body)
	// 	if err != nil {
	// 		log.Fatal("Error in response Data")
	// 		return
	// 	}

	// 	fmt.Printf("Data : %s", data)

	// 	fmt.Println(resp)
}
