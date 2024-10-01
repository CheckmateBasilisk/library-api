package core

import (
    "time"
)

type User struct {
    id string
    name string
    birthday time.Time //ignore anything other than dd/mm/yyyy
}
