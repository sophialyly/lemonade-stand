package main

import (
	"html/template"
	"net/http"
	"os"

	"github.com/SteveAzz/lemonade-stand/controllers"
)

func main() {
	templates := populateTemplates()

	controllers.Register(templates)

	http.ListenAndServe(":8000", nil)
}

func populateTemplates() *template.Template {
	result := template.New("templates")

	basePath := "templates"

	// Get the information from the template dir
	templateFolder, _ := os.Open(basePath)
	defer templateFolder.Close()

	// Read all of the dir
	templatePathsRaw, _ := templateFolder.Readdir(-1)

	// Create new slice for all of the templates
	templatePaths := new([]string)
	for _, pathInfo := range templatePathsRaw {
		if !pathInfo.IsDir() {
			*templatePaths = append(*templatePaths, basePath+"/"+pathInfo.Name())
		}
	}

	result.ParseFiles(*templatePaths...)

	return result
}
