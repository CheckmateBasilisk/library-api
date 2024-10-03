package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

    "github.com/CheckmateBasilisk/library-api/handlers"
)

// contains all routes and calls the handlers
// FIXME: think of a way to go back to core so that neither routes or handlers knows about the database in ANY way
func BuildRouter() *chi.Mux { //TODO: recieves config?
    router := chi.NewRouter()
    router.Use(middleware.Logger)

    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    })

    //creates a subrouter
    router.Route("/datatype", func(router chi.Router) {
        router.Post("/", handlers.CreateDatatype)
        router.Get("/", handlers.ReadDatatype)
        router.Get("/{id}", handlers.ReadDatatypeById)
        router.Put("/{id}", handlers.UpdateDatatypeById)
        router.Delete("/{id}", handlers.DeleteDatatypeById)
    })

    return router
}



