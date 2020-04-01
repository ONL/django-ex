package main

import (
	"net/http"
	"fmt"
	"errors"
)

func quellenHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, "quellen", "base", nil)
}

func aboutHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
	return renderTemplate(w, "about", "base", nil)
}

// homeHandler handles all requests, as other handlers redirect here with added
// parameters in env.
func homeHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
    return renderTemplate(w, "index", "base", nil)
}

type LoginContent struct {
  Pwfor string
  Next string
}

type AdminContent struct {
  Afrika-klima-lospw string
  Afrika-vegetation-lospw string
}

func adminHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
   	session, error := env.Store.Get(r, "admin")
	if nil != error {
		return errors.New(http.StatusText(http.StatusUnauthorized))
	}
	
	loginContent := &LoginContent {
		Pwfor: "Administration",
		Next: "/admin"}
	adminContent := &AdminContent {
		Afrika-klima-lospw: env.args["afrika-klima-lospw"]
		Afrika-vegetation-lospw: env.args["afrika-vegetation-lospw"] }
	
	if http.MethodGet == r.Method {
		if "yes" == session.Values["auth"] {
			return renderTemplate(w, "admin", "admin", adminContent)
		} else {
			return renderTemplate(w, "login", "login", loginContent)
		}
	} else {
		r.ParseForm()
		switch action := r.PostForm.Get("action"); action {
			case "login":
				if r.PostForm.Get("pw") != env.args["Password"] {
					session.Values["auth"] = "no"
					error = session.Save(r, w)
					if error != nil {
						return errors.New(http.StatusText(http.StatusUnauthorized))
					}
					return errors.New(http.StatusText(http.StatusUnauthorized))
				} else {
					session.Values["auth"] = "yes"
					error = session.Save(r, w)
					if error != nil {
						return errors.New(http.StatusText(http.StatusUnauthorized))
					}
					return renderTemplate(w, "admin", "admin", adminContent)
				}
			case "updatepw":
				if "yes" == session.Values["auth"] {
					env.args["afrika-klima-lospw"] = r.PostForm.Get("afrika-klima-lospw")
					env.args["afrika-vegetation-lospw"] = r.PostForm.Get("afrika-vegetation-lospw")
					adminContent.Afrika-klima-lospw = env.args["afrika-klima-lospw"]
					adminContent.Afrika-vegetation-lospw = env.args["afrika-vegetation-lospw"]
					return renderTemplate(w, "admin", "admin", adminContent)
				} else {
					return renderTemplate(w, "login", "login", loginContent)
				}
			default:
				return errors.New(http.StatusText(http.StatusUnauthorized))
		}
		
	}
}
