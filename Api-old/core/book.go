package core

type Book struct {
    Id string // this is a leakage from the db concepts all the way through the API... dont know if there's a fix
    Isbn string
    Name string
    Author string
    Year int //FIXME: does it make sense to use time? idk...
    Synopsis string
}

// FIXME: is this needed? probably not since I can instantiante structs with Name{fields...}
//  i think this way of thinking is an artifact of my time with C... I had to make constructors and destructor for structs by hand
func newBook(id string, name string, isbn string, author string, year int, synopsis string) *Book {
    //validate fields
    //fill instace
    return &Book{id, name, isbn, author, year, synopsis}
}

