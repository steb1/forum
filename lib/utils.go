package lib

import (
	"bufio"
	"fmt"
	"forum/data/models"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"text/template"

	"github.com/google/uuid"
)

func LoadEnv(path string) error {
	file, err := os.Open(path)
	if err != nil {
		log.Println("🚨 " + err.Error())
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Println("🚨 Your env file must be set")
		}
		key := parts[0]
		value := parts[1]
		os.Setenv(key, value)
	}
	return scanner.Err()
}

func ValidateRequest(req *http.Request, res http.ResponseWriter, url, method string) bool {
	if strings.Contains(url, "*") && path.Dir(url) == path.Dir(req.URL.Path) {
		return true
	}

	if req.URL.Path != url {
		res.WriteHeader(http.StatusNotFound)
		RenderPage("base", "404", nil, res)
		log.Println("404 ❌ - Page not found ", req.URL)
		return false
	}

	if req.Method != method {
		res.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(res, "%s", "Error - Method not allowed")
		log.Printf("405 ❌ - Method not allowed %s - %s on URL : %s\n", req.Method, method, url)
		return false
	}
	return true
}

func RenderPage(basePath, pagePath string, data any, res http.ResponseWriter) {
	files := []string{"templates/common/" + basePath + ".html", "templates/" + pagePath + ".html"}
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Println("🚨 " + err.Error())
	} else {
		tpl.Execute(res, data)
	}
}

func CheckUsers(data []models.User, Email, Username string) bool {
	for _, val := range data {
		//fmt.Println(val.ID)
		if val.Email == Email || val.Username == Username {
			return false
		}
	}

	return true
}

func Isregistered(data []models.User, email, password string) (string, bool) {
	for _, val := range data {
		//fmt.Println(val.ID)
		if val.Email == email && val.Password == password {
			return val.ID, true
		}
	}

	return password, false
}

// func GenerateImageUrl(url io.Reader) string {
// 	u := uuid.New()
// 	time := time.Now().Format("2006-01-02 14:24:24")

//		return u.String() + time
//	}
func UploadImage(r *http.Request) string {
	image, header, err := r.FormFile("image")
	uploads := "/uploads/"
	u := uuid.New()
	imageURL := uploads + u.String() + header.Filename
	file, err := os.Create(imageURL)
	if err != nil {
		fmt.Println("Erreur lors de la création du fichier :", err)
		return ""
	}
	defer file.Close()
	_, err = io.Copy(file, image)
	if err != nil {
		fmt.Println("Erreur lors de la copie des données :", err)
		return ""
	}

	return imageURL
}
