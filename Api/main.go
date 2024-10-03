package main

import (
    "fmt"
    "context"

    "github.com/CheckmateBasilisk/library-api/core"
)

//FIXME: REORGANIZE PACKAGES. JOIN RELEVANT ONES!!!!
func main() {
    // config
    // TODO: get these from params or env variables. Easy to send to application from docker or cli!
    port := 6969
    baseUrl := "localhost"

    // TODO: remove this discussion
    //  it should be something like ? INSIDE CORE!!!
    //     main should only give configs, not define funcionality!
    /*
    server = server.newServer(config)
    routes = routes.newRoutes(config)
    db = db.newDb(config)
    core.NewCore()
    core.Run()
    */

    app := core.NewCore(baseUrl, port)
    err := app.Run(context.TODO()) //FIXME: change todo to actual context

    if err != nil {
        fmt.Println("failed to start app: ", err)
    }

}



