package core

import (
    "time"
)

type User struct {
    Id string
    Name string
    //birthday time.Time //ignore anything other than dd/mm/yyyy
    Birthday string//FIXME: this shoulnt be string
}
