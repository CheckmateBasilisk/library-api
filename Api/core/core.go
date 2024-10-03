package core

import (
	"fmt"
	"net/http"
    "context"

	//"github.com/CheckmateBasilisk/library-api/routes"

    //FIXME: the goal is to make core unaware of the innards of server/router creation, it just has to run ListenAndServe
    //  is it possible??
)

type Core struct {
    // port, baseUrl, database details, router...
    // FIXME: should this have substructs? idk... ill find out eventually
    // server + routes and handlers
    baseUrl string
    port int
    router http.Handler
    // database
}

//instantiate new core
// fill it with config (injecting deps)
//func NewCore(router http.Handler) *Core { //FIXME: move the router creation to routes.go and inject it through main during core creation
func NewCore(baseUrl string, port int, router http.Handler) *Core {
    core := &Core{
        baseUrl: baseUrl,
        port: port,
        router: router,
        //FIXME: remove the route loading from here, defer it to main?
        //      anyways, routes.go should be able to create the router and this should get it from there, not create it by itself -> remove loadRoutes?
        //router: routes // plugs the router defined @/routes/routes.go here
    }
    return core
}

func (core *Core) Run(ctx context.Context) error {
    server := &http.Server{
        Addr: fmt.Sprintf(":%d", core.port),
        Handler: core.router,
    }

    err := server.ListenAndServe()
    if err != nil {
        return fmt.Errorf("error serving %w", err)
    }
    return nil // if everything is ok
}
