package core

import (
    "time"
    //"github.com/CheckmateBasilisk/library-api/db"
    //"github.com/CheckmateBasilisk/library-api/server"
)
// User <-> API <-> core + datatypes <-> db <-> actual db
// User>api>core>db>core [compute]>api>User
// core has access to datatypes. Api and db deal in strings only
// core has business logic to manipulate datatypes
// whats the difference between mvc and this?
//  I think this way the project is vertical instead of horizontal ?
//  each type is not spread through all layers, like in MVC ???

// core represents "actions", verbs. Loaning a book is here, but Book and User are somewhere else...
//  this way, core has FUNCIONAILITIES... should functionalities be in other .go files in this package and core just orchestrate them?

// functionality... business logic!
//  also deal with exceptions and failures
func makeLoan(user *User, book *Book, start time.Time, end time.Time) Loan {
    id := "laliullelo" // think about this...
    return Loan{id, book, user, start, end, time.Time{}}// end time == zero
}

