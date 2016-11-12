package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SteveAzz/lemonade-stand/controllers/util"
	"github.com/SteveAzz/lemonade-stand/viewmodels"
)

type profileController struct {
	template *template.Template
}

func (controller *profileController) handle(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	vm := viewmodels.GetProfile()
	responseWriter.Header().Add("Content-Type", "text/html")

	if req.Method == "POST" {
		vm.User.Email = req.FormValue("email")
		vm.User.FirstName = req.FormValue("firstName")
		vm.User.LastName = req.FormValue("lastName")
		vm.User.Stand.Address = req.FormValue("standAddress")
		vm.User.Stand.City = req.FormValue("standCity")
		vm.User.Stand.State = req.FormValue("standState")
		vm.User.Stand.Zip = req.FormValue("standState")
	}

	err := controller.template.Execute(responseWriter, vm)
	if err != nil {
		log.Fatal(err)
	}
}
