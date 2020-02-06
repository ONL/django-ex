package main

import (
	"net/http"
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"github.com/gorilla/mux"
	"log"
	"context"
	"syscall"
	"strconv"
	"time"
	"errors"
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
	args map[string]*string
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
    
    c := &controller{
		logger: logger,
		nextRequestID: func() string { return strconv.FormatInt(time.Now().UnixNano(), 36) },
		}
    
    env := &Env{
		c: c,
		args: ARGS,
	}
    
    router := mux.NewRouter()
    router.Handle("/", Handler{env, homeHandler})
    router.Handle("/themen",  Handler{env, themeHandler})
    
    server := &http.Server{
		Addr:         "127.0.0.1",
		Handler:      router,
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	ctx := c.shutdown(context.Background(), server)

	logger.Printf("Server is ready to handle requests at %q\n", "127.0.0.1")
	atomic.StoreInt64(&c.healthy, time.Now().UnixNano())

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", "127.0.0.1", err)
	}
	<-ctx.Done()
	logger.Printf("Server stopped\n")	
}


// homeHandler handles all requests, as other handlers redirect here with added
// parameters in env.
func homeHandler(env *Env, w http.ResponseWriter, r *http.Request) error {
    
    w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

    username, password, authOK := r.BasicAuth()
    
    if false == authOK {
        return errors.New(http.StatusText(http.StatusUnauthorized))
    }

    if password != *env.args["Password"] { {
        return errors.New(http.StatusText(http.StatusUnauthorized))
    }

    fmt.Fprintf(w, "%+v", username)
    w.Header().Set("X-Forwarded-User", username)
    return nil
}
