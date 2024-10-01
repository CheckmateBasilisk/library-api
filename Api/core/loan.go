package core

import (
    "time"
)

type Loan struct {
    id string
    book *Book
    user *User
    loanStart time.Time
    loanEnd time.Time
    returnDate time.Time //accounts for late returns
    // a status makes sense?
}
