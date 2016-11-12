package controllers

import (
	"html/template"
	"log"
	"net/http"

	"github.com/SteveAzz/lemonade-stand/controllers/util"
	"github.com/SteveAzz/lemonade-stand/models"
	"github.com/SteveAzz/lemonade-stand/viewmodels"
)

type homeController struct {
	template      *template.Template
	loginTemplate *template.Template
}

func (controller *homeController) get(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")
	vm := viewmodels.GetHome()

	err := controller.template.Execute(responseWriter, vm)
	if err != nil {
		log.Fatal(err)
	}
}

func (controller *homeController) login(w http.ResponseWriter, req *http.Request) {
	responseWriter := util.GetResponseWriter(w, req)
	defer responseWriter.Close()

	responseWriter.Header().Add("Content-Type", "text/html")
	vm := viewmodels.GetLogin()

	if req.Method == "POST" {
		email := req.FormValue("email")
		password := req.FormValue("password")

		member, err := models.GetMember(email, password)

		if err != nil {
			log.Fatal(err)
		}

		session, err := models.CreateSession(member)

		if err != nil {
			log.Fatal(err)
		}

		var cookie http.Cookie
		cookie.Name = "sessionID"
		cookie.Value = session.SessionID()
		responseWriter.Header().Add("Set-Cookie", cookie.String())
	}

	err := controller.loginTemplate.Execute(responseWriter, vm)

	if err != nil {
		log.Fatal(err)
	}
}
