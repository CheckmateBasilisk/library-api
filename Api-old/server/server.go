package server

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Welcome(port int) {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

    //config server
    server := &http.Server{
        Addr: fmt.Sprintf(":%d", port),//like :3000
        Handler: router,
    }
    //starting server
	err := server.ListenAndServe()
    if err != nil {
        fmt.Errorf("failed to start server: %w", err) // FIXME: use logger or appropriate error print
    }
}

//endpoints
// books
//  create
//  read
//  update
//  delete
// users
// loans

