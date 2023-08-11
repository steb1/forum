package auth

import (
	"context"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// Create an OAuth2 configuration using your client ID and secret.
var (
	googleOAuthConfig = oauth2.Config{
		ClientID:     "889533868443-q0ih7c2vah44pbdn5ouag0437pfeb478.apps.googleusercontent.com",         // Replace with your actual client ID
		ClientSecret: "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",     // Replace with your actual client secret
		RedirectURL:  "http://localhost:8080/callback", // Replace with your actual redirect URI
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email"}, // Request specific scopes
		Endpoint: google.Endpoint, // Google's OAuth2 endpoint
	}
)

// func main() {
// 	// Set up HTTP routes and server
// 	http.HandleFunc("/", handleHome)
// 	http.HandleFunc("/login", handleLogin)
// 	http.HandleFunc("/callback", handleCallback)
// 	http.ListenAndServe(":8080", nil) // Start the server on port 8080
// }


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

	// Parse the response and use user information
	// In a real application, you would process the response to get user details

	fmt.Fprint(w, "Logged in successfully!")
}