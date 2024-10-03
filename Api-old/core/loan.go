package core

import (
    //"time"
)

type Loan struct {
    Id string
    Book *Book
    User *User
    //FIXME: change from string to time
    LoanStart string
    LoanEnd string
    ReturnDate string
    //loanStart time.Time
    //loanEnd time.Time
    //returnDate time.Time //accounts for late returns
    // a status makes sense?
}
