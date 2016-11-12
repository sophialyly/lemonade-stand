package controllers

import (
	"html/template"
	"log"
	"net/http"

	"strconv"

	"github.com/SteveAzz/lemonade-stand/controllers/convertors"
	"github.com/SteveAzz/lemonade-stand/controllers/util"
	"github.com/SteveAzz/lemonade-stand/models"
	"github.com/SteveAzz/lemonade-stand/viewmodels"
	"github.com/gorilla/mux"
)

type categoriesController struct {
	template *template.Template
}

func (controller *categoriesController) get(w http.ResponseWriter, req *http.Request) {
	// Get data from model
	categories := models.GetCategories()

	// ViewModel Category slice
	categoriesVM := []viewmodels.Category{}

	// Convert model to view model
	for _, category := range categories {
		categoriesVM = append(categoriesVM, convertors.ConvertCategoryToViewModel(category))
	}

	// Create new view model
	vm := viewmodels.GetCategories()
	// Assigned Category slice
	vm.Categories = categoriesVM

	w.Header().Add("Content-Type", "text/html")

	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	err := controller.template.Execute(responseWriter, vm)
	if err != nil {
		log.Fatal(err)
	}
}

type categoryController struct {
	template *template.Template
}

func (controller *categoryController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	vm := viewmodels.GetProducts(id)

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	tmplErr := controller.template.Execute(responseWriter, vm)

	if tmplErr != nil {
		log.Fatal(tmplErr)
	}
}
