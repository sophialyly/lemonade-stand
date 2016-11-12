package controllers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/SteveAzz/lemonade-stand/controllers/util"
	"github.com/SteveAzz/lemonade-stand/viewmodels"
)

type standLocatorController struct {
	template *template.Template
}

func (controller *standLocatorController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocator()
	responseWriter.Header().Add("Content-Type", "text/html")

	err := controller.template.Execute(responseWriter, vm)

	if err != nil {
		log.Fatal(err)
	}
}

func (controller *standLocatorController) apiSearch(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetStandLocations()
	responseWriter.Header().Add("Content-Type", "application/json")

	data, err := json.Marshal(vm)

	if err != nil {
		log.Fatal(err)
	}

	responseWriter.Write(data)
}
