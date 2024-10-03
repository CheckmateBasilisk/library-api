package main

import (
    "fmt"
    //"github.com/CheckmateBasilisk/library-api/core"
    "github.com/CheckmateBasilisk/library-api/db"
    "github.com/CheckmateBasilisk/library-api/server"
)

func main() {
    fmt.Printf("starting database...\n")
    fmt.Printf("starting server...\n")
    fmt.Printf("done!...\n")

    fmt.Println(db.Connect())

    server.Welcome(6969)
}
