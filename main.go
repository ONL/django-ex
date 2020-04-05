package main

import (
	"net/http"
	"html/template"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"github.com/gorilla/sessions"
	"log"
	"context"
	"syscall"
	"strconv"
	"time"
)

// controller allows logging of the server
type controller struct {
	logger        *log.Logger
	nextRequestID func() string
	healthy       int64
}

// Env hold configuration parametes, that are passed to http handlers
type Env struct {
	c   *controller
	args map[string]string
	Store *sessions.CookieStore
}

// Handler is a custom http.Handler allowing environment data to be passed to the handler functions.
type Handler struct {
	*Env
	H func(e *Env, w http.ResponseWriter, r *http.Request) error
}

// ServeHTTP allows our Handler type to satisfy http.Handler.
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.H(h.Env, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
	}
}

// shutdown shuts down the server gracefully
func (c *controller) shutdown(ctx context.Context, server *http.Server) context.Context {
	ctx, done := context.WithCancel(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer done()

		<-quit
		signal.Stop(quit)
		close(quit)

		atomic.StoreInt64(&c.healthy, 0)
		server.ErrorLog.Printf("Server is shutting down...\n")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			server.ErrorLog.Fatalf("Could not gracefully shutdown the server: %s\n", err)
		}
	}()

	return ctx
}


// main - The main logic.
func main() {
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Server is starting...")
	
	ARGS := make(map[string]string)

  	ARGS["Password"] = os.Getenv("PASSWORD")
	ARGS["afrika-klima-lospw"] = os.Getenv("PASSWORD")
	ARGS["afrika-vegetation-lospw"] = os.Getenv("PASSWORD")
    
    c := &controller{
		logger: logger,
		nextRequestID: func() string { return strconv.FormatInt(time.Now().UnixNano(), 36) },
		}
	sessionpw, _ := bcrypt.GenerateFromPassword([]byte( os.Getenv("PASSWORD") ),14)
	store := sessions.NewCookieStore( sessionpw )
    env := &Env{
		c: c,
		args: ARGS,
	    	Store: store,
	}
    
    router := mux.NewRouter()
    router.Handle("/", Handler{env, homeHandler})
    router.Handle("/about",  Handler{env, aboutHandler})
    router.Handle("/lehrer",  Handler{env, adminHandler})
    router.Handle("/quellen",  Handler{env, quellenHandler})
	router.Handle("/afrika", Handler{env, afrikaHandler})
    router.Handle("/afrika/klima", Handler{env, afrikaKlimaHandler})
    router.Handle("/afrika/vegetation", Handler{env, afrikaVegetationHandler})
    router.Handle("/afrika/klima-los", Handler{env, afrikaKlimaLosHandler})
    router.Handle("/afrika/vegetation-los", Handler{env, afrikaVegetationLosHandler})
    router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
    
    server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	ctx := c.shutdown(context.Background(), server)

	logger.Printf("Server is ready to handle requests at %q\n", ":8080")
	atomic.StoreInt64(&c.healthy, time.Now().UnixNano())

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", ":8080", err)
	}
	<-ctx.Done()
	logger.Printf("Server stopped\n")	
}

var templates map[string]*template.Template

//Render templates for the given name, template definition and data object
func renderTemplate(w http.ResponseWriter, name string, template string, viewModel interface{}) error {
	// Ensure the template exists in the map.
	tmpl, ok := templates[name]
	if !ok {
	http.Error(w, "The template does not exist.", http.StatusInternalServerError)
	}
	err := tmpl.ExecuteTemplate(w, template, viewModel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
	}
	return err
}


//Compile view templates
func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	
	templates["index"] = template.Must(template.New("index").Funcs(template.FuncMap{
        "IterateOne": func(count *int) []int {
            var i int
            var Items []int
            for i = 1; i <= (*count); i++ {
                Items = append(Items, i)
            }
            return Items
        },
		"Iterate": func(count *int) []int {
            var i int
            var Items []int
            for i = 0; i < (*count); i++ {
                Items = append(Items, i)
            }
            return Items
        },
	}).ParseFiles("templates/index.html",
		"templates/base.html"))
	templates["quellen"] = template.Must(template.ParseFiles("templates/quellen.html",
		"templates/base.html"))
	templates["about"] = template.Must(template.ParseFiles("templates/about.html",
		"templates/base.html"))
	templates["admin"] = template.Must(template.ParseFiles("templates/admin.html",
		"templates/base.html"))
	templates["login"] = template.Must(template.ParseFiles("templates/login.html",
		"templates/base.html"))
	templates["afrika"] = template.Must(template.ParseFiles("templates/afrika.html", 
		"templates/base.html"))
	templates["afrika_klima"] = template.Must(template.ParseFiles("templates/afrika-map.html", "templates/afrika-klima.html",
		"templates/base.html"))
	templates["afrika_vegetation"] = template.Must(template.ParseFiles("templates/afrika-map.html", "templates/afrika-vegetation.html",
		"templates/base.html"))
}
