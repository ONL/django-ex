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
    
    w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

    username, password, authOK := r.BasicAuth()
    
    if false == authOK {
        return errors.New(http.StatusText(http.StatusUnauthorized))
    }

    if password != *env.args["Password"] {
        return errors.New(http.StatusText(http.StatusUnauthorized))
    }

    fmt.Fprintf(w, "%+v", username)
    w.Header().Set("X-Forwarded-User", username)
    return renderTemplate(w, "index", "base", nil)
}
