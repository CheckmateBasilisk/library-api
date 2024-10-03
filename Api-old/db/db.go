package db

import (
    //"gorm.io/gorm"
    //"gorm.io/driver/mysql"
    "github.com/CheckmateBasilisk/library-api/core"
)

//TODO: implement this for reals..
func Connect() string {
    return "[mock] connecting..."
}


//===============MOCK DATA===================

var MockUsers = []*core.User {
    {
        Id: "99990",
        Name: "Ana Fantasy",
        Birthday: "10/01/2000",
    },
    {
        Id: "99991",
        Name: "Bob Sci Fi",
        Birthday: "29/01/1990",
    },
    {
        Id: "99992",
        Name: "Charlie Horror",
        Birthday: "31/12/2017",
    },

}

var MockBooks = []*core.Book {
    {
        Id: "001",
        Isbn: "9781611748864",
        Name: "The Lord of the Rings",
        Author: "J R R Tolkien",
        Year: 1954,
        Synopsis: "When Mr. Bilbo Baggins of Bag End announced that he would shortly be celebrating his eleventy-first birthday with a party of special magnificence, there was much talk and excitement in Hobbiton.",
    },
    {
        Id: "002",
        Isbn: "9780887484322",
        Name: "Blindsight",
        Author: "Peter Watts",
        Year: 2006,
        Synopsis: "Sara Linton leaned back in her chair, mumbling a soft \"Yes, Mama\" into the telephone.",
    },
    {
        Id: "003",
        Isbn: "0932096379",
        Name: "I, Robot",
        Author: "Issac Azimov",
        Year: 1950,
        Synopsis: "I looked at my notes and I didn't like them.",
    },
    {
        Id: "004",
        Isbn: "9780982751169",
        Name: "Frankenstein or The Modern Prometheus",
        Author: "Mary Shelley",
        Year: 1818,
        Synopsis: "Starting from fish-shape Paumanok, where I was born",
    },
}

var MockLoans = []*core.Loan {
    {
        Id: "1234",
        Book: MockBooks[0],
        User: MockUsers[0],
        LoanStart: "01/10/2024",
        LoanEnd: "08/10/2024",
        ReturnDate: "",
    },
}


