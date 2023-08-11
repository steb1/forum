package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Create an OAuth2 configuration using your client ID and secret.
var (
	googleOAuthConfig = oauth2.Config{
		ClientID:     "889533868443-q0ih7c2vah44pbdn5ouag0437pfeb478.apps.googleusercontent.com", // Replace with your actual client ID
		ClientSecret: "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",                                      // Replace with your actual client secret
		RedirectURL:  "http://localhost:8080/callback",                                           // Replace with your actual redirect URI
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},                 // Request specific scopes
		Endpoint:     google.Endpoint,                                                            // Google's OAuth2 endpoint
	}
)

// Handle requests to initiate Google Sign-In
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Generate the URL for Google Sign-In and redirect the user
	url := googleOAuthConfig.AuthCodeURL("", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// Handle the callback from Google after the user signs in
func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code") // Get the authorization code from the query parameter
	token, err := googleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token", http.StatusInternalServerError)
		return
	}

	// Create an authenticated HTTP client using the token
	client := googleOAuthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseData, _ := io.ReadAll(response.Body)

	var Data GoogleUser

	json.Unmarshal(responseData, &Data)

	fmt.Fprint(w, "Logged in successfully!")
}

type GoogleUser struct {
	ID       string `json: "id"`
	Email    string `json: "email"`
	Isvalid  bool   `json: "verified_email"`
	ImageURL string `json : "picture"`
}
