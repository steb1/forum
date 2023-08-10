package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/gofrs/uuid"
)

var maxSize int64 = 20 * 1024 * 1024 // 20 MB

func Slugify(input string) string {
	input = strings.ToLower(input)
	re := regexp.MustCompile("[^a-z0-9]+")
	input = re.ReplaceAllString(input, "-")
	input = strings.Trim(input, "-")

	return input
}

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

func UploadImage(req *http.Request) string {
	image, header, err := req.FormFile("image")
	if err != nil {
		log.Println("❌ Request doesn't contain image", err)
		return ""
	}
	defer image.Close()

	if header.Size > maxSize {
		log.Println("❌ File size exceeds limit")
		return ""
	}

	if !isValidFileType(header.Header.Get("Content-Type")) {
		log.Println("❌ Invalid file type")
		return ""
	}

	uploads := "uploads" // Use "uploads" without the leading slash
	imageURL := filepath.Join(uploads, generateUniqueFilename(header.Filename))
	filePath := filepath.Join(".", imageURL) // Use "." to denote the current directory
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("❌ Error when creating the file", err)
		return ""
	}
	defer file.Close()
	_, err = io.Copy(file, image)
	if err != nil {
		fmt.Println("❌ Error when copying data", err)
		return ""
	}

	return imageURL
}

func generateUniqueFilename(filename string) string {
	ext := filepath.Ext(filename)
	randomName, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("❌ Failed to generate UUID: %v", err)
	}
	newFilename := randomName.String() + ext
	return newFilename
}

func isValidFileType(contentType string) bool {
	switch contentType {
	case "image/jpeg", "image/png", "image/gif":
		return true
	}
	return false
}
