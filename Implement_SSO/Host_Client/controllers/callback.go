package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

func Callback(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	code := r.FormValue("code")

	if state != RandomVariable {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	token, err := sso.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		log.Println("Token not validated or verified:", err)
		return
	}
	url := "http://localhost:5000/auth/dreamsleep"
	bearer := "Bearer " + token.AccessToken

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Failed to create request:", err)
		return
	}

	req.Header.Add("Authorization", bearer)
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Error on response:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Successfully authenticated by dreamsleep server")
	} else {
		fmt.Printf("Authentication failed: %s\n", resp.Status)
	}
}
