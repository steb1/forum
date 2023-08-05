package lib

import (
	"forum/data/models"
)

func CheckUsers(data []models.User, Email, Username string) bool {
	for _, val := range data {
		//fmt.Println(val.ID)
		if val.Email == Email || val.Username == Username {
			return false
		}
	}

	return true
}

func Isregistered(data []models.User, email, password string) (string , bool) {
	for _, val := range data {
		//fmt.Println(val.ID)
		if val.Email == email && val.Password == password {
			return val.ID, true
		}
	}

	return password, false
}
