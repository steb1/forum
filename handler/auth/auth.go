package auth

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"forum/data/models"
	"forum/lib"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gofrs/uuid"
)

const tokenExpiration = time.Hour

type Token struct {
	UserID    uuid.UUID `json:"user_id"`
	Username  string    `json:"username"`
	ExpiresAt int64     `json:"exp"`
}

func SignupHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { //Verifier que le formulaire est bien structure
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	////////////////////////////////////////////////////////////////

	email := r.FormValue("mail")
	username := r.FormValue("username")
	password := r.FormValue("password")
	avatarURL := "assets/img/community.webp"
	Type := 3

	//////////////////////////////////////////////////////////////////

	db, _ := sql.Open("sqlite3", "./data/sql/forum.db")

	data, err := models.NewUserRepository(db).SelectAllUsers()

	if err != nil {
		return
	}

	ok := lib.CheckUsers(data, email, username)

	if ok {
		// Create a Version 4 UUID.
		ID, err := uuid.NewV4()

		if err != nil {
			log.Fatalf("failed to generate UUID: %v", err)
		}

		// Prepared Statement
		stmt, err := db.Prepare("INSERT INTO user (ID, username, email, password, avatarURL, type, token, tokenExpirationDate) VALUES(?, ?, ?, ?, ?, ?, ?, ?)")

		defer db.Close()

		if err != nil {
			log.Fatal(err)
		}

		defer stmt.Close()

		// Créer un nouveau jeton
		token := models.Token{
			UserID:    ID.String(),
			Username:  username,
			ExpiresAt: time.Now().Add(tokenExpiration * 2).Unix(),
		}

		// Encodage du jeton au format JSON
		tokenJson, err := json.Marshal(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		////////////////////////////////////////////////////////////////////
		// Set cookie

		cookie := http.Cookie{}
		cookie.Name = username
		cookie.Value = ID.String()
		cookie.Expires = time.Now().Add(2 * time.Hour)
		cookie.Secure = true
		cookie.HttpOnly = true
		http.SetCookie(w, &cookie)

		// Renvoyer le jeton encodé dans la réponse

		w.Header().Set("Content-Type", "application/json")
		w.Write(tokenJson)

		////////////////////////////////////////////////////////

		_, err = stmt.Exec(ID, email, username, password, avatarURL, Type, token.UserID, token.ExpiresAt)

		fmt.Println("Success")
	} else {
		fmt.Println("Failed")
		return
	}
}

func SigninHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil { //Verifier que le formulaire est bien structure
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	db, _ := sql.Open("sqlite3", "data\\sql\\forum.db")

	email := r.FormValue("mail")
	username := r.FormValue("username")
	password := r.FormValue("password")

	data, err := models.NewUserRepository(db).SelectAllUsers()

	if err != nil {
		log.Fatal(err)
	}

	ID, exists := lib.Isregistered(data, email, password)

	if !exists {
		fmt.Println("failed")
		return
	}

	// Créer un nouveau jeton
	token := models.Token{
		UserID:    ID,
		Username:  username,
		ExpiresAt: time.Now().Add(tokenExpiration * 2).Unix(),
	}

	// Encodage du jeton au format JSON
	tokenJson, err := json.Marshal(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	////////////////////////////////////////////////////////////////////
	// Set cookie

	cookie := http.Cookie{}
	cookie.Name = username
	cookie.Value = ID
	cookie.Expires = time.Now().Add(2 * time.Hour)
	cookie.Secure = true
	cookie.HttpOnly = true
	http.SetCookie(w, &cookie)

	// UPDATE CLIENT TOKEN

	stmt, err := db.Prepare("UPDATE user SET tokenExpirationDate = ? WHERE id = ?")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(token.ExpiresAt, ID)
	if err != nil {
		log.Fatal(err)
	}

	// Renvoyer le jeton encodé dans la réponse

	w.Header().Set("Content-Type", "application/json")
	w.Write(tokenJson)

	t, err := template.ParseFiles("templates\\user\\user.html")
	fmt.Println("signin success")

	t.Execute(w, ID)
}

// // this map stores the users sessions. For larger scale applications, you can use a database or cache for this purpose
// var sessions = map[string]session{}

// // each session contains the username of the user and the time at which it expires
// type session struct {
// 	username string
// 	expiry   time.Time
// }

// // we'll use this method later to determine if the session has expired
// func (s session) isExpired() bool {
// 	return s.expiry.Before(time.Now())
// }
