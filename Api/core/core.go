package core

import (
	"context"
	"fmt"

	"github.com/CheckmateBasilisk/library-api/server"
	"github.com/CheckmateBasilisk/library-api/routes" //FIXME: should routes be a subpcakage of server?
)

type Core struct {
	// port, baseUrl, database details, router...
    // FIXME: should this have substructs? idk... ill find out eventually
    // server + routes and handlers
    baseUrl string
    port int
    // database
}

//instantiate new core
// fill it with config (injecting deps?)
func NewCore(baseUrl string, port int) *Core {
    core := &Core{
        baseUrl: baseUrl,
        port: port,
    }
    return core
}

func (core *Core) Run(ctx context.Context) error { //FIXME: why ctx??
    server := server.NewServer(core.baseUrl, core.port, routes.BuildRouter())

    err := server.ListenAndServe()
    if err != nil {
        return fmt.Errorf("error serving %w", err)
    }
    return nil // if everything is ok
}
