package core

import (
    "time"
)

type Loan struct {
    Id string
    Book *Book
    User *User
    //FIXME: change from string to time
    LoanStart string
    LoanEnd string
    ReturnDate string
    //LoanStart time.Time
    //LoanEnd time.Time
    //ReturnDate time.Time //accounts for late returns
    // a status makes sense?
}
