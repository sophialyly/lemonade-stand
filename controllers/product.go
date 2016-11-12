package controllers

import  (
	"html/template"
	"net/http"
 	"strconv"
 	"log"
	 
 	"github.com/gorilla/mux"
 	"github.com/SteveAzz/lemonade-stand/viewmodels"
 	"github.com/SteveAzz/lemonade-stand/controllers/util"
)

type productController struct {
	template *template.Template
}

func (controller *productController) get(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	idRaw := vars["id"]

	id, err := strconv.Atoi(idRaw)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	}

	vm := viewmodels.GetProduct(id)

	w.Header().Add("Content-Type", "text/html")
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()
	tmplErr := controller.template.Execute(responseWriter, vm)

	if tmplErr != nil {
		log.Fatal(tmplErr)
	}
}
