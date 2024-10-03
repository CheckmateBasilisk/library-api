package main

import (
    "fmt"
    "context"

    "github.com/CheckmateBasilisk/library-api/core"
    "github.com/CheckmateBasilisk/library-api/routes"
)

//FIXME: REORGANIZE PACKAGES. JOIN RELEVANT ONES!!!!
func main() {
    // config
    port := 6969
    baseUrl := "localhost"
    router := routes.BuildRouter()

    //FIXME: defer server creation to server.go ??
    //  it should be something like ?
    /*
    server = server.newServer(config)
    routes = routes.newRoutes(config)
    db = db.newDb(config)

    core.NewCore()
    core.Run()

    */
    app := core.NewCore(baseUrl, port, router)
    err := app.Run(context.TODO()) //FIXME: change todo to actual context

    if err != nil {
        fmt.Println("failed to start app: ", err)
    }

}



