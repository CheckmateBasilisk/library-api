package server

import (
    "fmt"
    "net/http"
	"github.com/CheckmateBasilisk/library-api/routes"
)

//TODO: think about this... maybe I should defer the server creation to server.go instead of core.
//      but core has the configs, right? Should core.go call server.go or main.go should call server.go??
//      sending the router from main to the core feels weird... Maybe sever should know about this and inject it into core. This way rework only happens at the borders of the system

//FIXME: should server implicitly get the routes from the project or should  core give it to it???
//  routes seems to belong in /server/routes but refitting isnt made at core (by changing which routes...). Sounds like hard coupling / hardcoding
func NewServer(baseUrl string, port int, router http.Handler) *http.Server {
    server := &http.Server{
        Addr: fmt.Sprintf("%s:%d", baseUrl, port),
        Handler: routes.BuildRouter(),
    }

    return server
}
