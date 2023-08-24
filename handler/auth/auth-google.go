package auth

import (
	"encoding/json"
	"fmt"
	"forum/data/models"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	clientID     = "889533868443-q0ih7c2vah44pbdn5ouag0437pfeb478.apps.googleusercontent.com"
	clientSecret = "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"
	redirectURL  = "https://localhost:8085/callback"
)

var (
	authURL     = "https://accounts.google.com/o/oauth2/auth"
	tokenURL    = "https://accounts.google.com/o/oauth2/token"
	scope       = "https://www.googleapis.com/auth/userinfo.profile"
	userInfoURL = "https://www.googleapis.com/oauth2/v2/userinfo"
)

// Handle requests to initiate Google Sign-In
func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	loginURL := fmt.Sprintf("%s?client_id=%s&redirect_uri=%s&scope=%s&response_type=code",
		authURL, clientID, url.QueryEscape(redirectURL), url.QueryEscape(scope))
	http.Redirect(w, r, loginURL, http.StatusFound)
}

// Handle the callback from Google after the user signs in
func HandleCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	// Exchange authorization code for access token
	tokenURL := fmt.Sprintf("%s?code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=authorization_code",
		tokenURL, code, clientID, clientSecret, url.QueryEscape(redirectURL))

	resp, err := http.PostForm(tokenURL, url.Values{})
	if err != nil {
		http.Error(w, "Failed to exchange code for token", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	respbody, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, "Failed to get data", http.StatusInternalServerError)
		return
	}

	//fmt.Println(string(respbody))

	var AccessToken Token

	json.Unmarshal(respbody, &AccessToken)

	//fmt.Println(AccessToken.AccessToken)

	accessToken := AccessToken.AccessToken // Extract access token from response JSON

	userInfoResp, err := http.Get(userInfoURL + "?access_token=" + accessToken)
	if err != nil {
		fmt.Println("Error fetching user info:", err)
		http.Error(w, "Error fetching user info", http.StatusInternalServerError)
		return
	}
	defer userInfoResp.Body.Close()

	userInfo, err := ioutil.ReadAll(userInfoResp.Body)
	if err != nil {
		fmt.Println("Error reading user info response:", err)
		http.Error(w, "Error reading user info response", http.StatusInternalServerError)
		return
	}

	var Data GoogleUser

	json.Unmarshal(userInfo, &Data)

	user := models.User{}

	user.ID = Data.ID
	user.Username = Data.Name
	user.AvatarURL = Data.ImageURL
	user.Role = models.RoleUser

	if _, exist := models.UserRepo.IsExisted(Data.ID); !exist {
		err := models.UserRepo.CreateUser(&user)
		if err != nil {
			log.Fatalf("❌ Failed to created account %v", err)
		}

		models.NewSessionToken(w, user.ID, user.Username)

		http.Redirect(w, r, "/", http.StatusSeeOther)
		log.Println("✅ Account created with success")
	} else {
		models.NewSessionToken(w, user.ID, user.Username)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		fmt.Println("❌ User already exist")
	}
}

type GoogleUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	ImageURL string `json:"picture"`
}

type Token struct {
	AccessToken string `json:"access_token"`
}
