package handlers

import (
    "fmt"
    "net/http"
)


type Datatype struct {}

//CRUD handlers for Datatype
func CreateDatatype(w http.ResponseWriter, r *http.Request) {
    fmt.Println("creating Datatype")
}
func ReadDatatype(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reading Datatype")
}
func ReadDatatypeById(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reading Datatype by ID")
}
func UpdateDatatypeById(w http.ResponseWriter, r *http.Request) {
    fmt.Println("updating Datatype by ID")
}
func DeleteDatatypeById(w http.ResponseWriter, r *http.Request) {
    fmt.Println("deleting Datatype by ID")
}

/*
//CRUD handlers for Datatype
func (d *Datatype) Create(w http.ResponseWriter, r *http.Request) {
    fmt.Println("creating Datatype")
}
func (d *Datatype) Read(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reading Datatype")
}
func (d *Datatype) ReadById(w http.ResponseWriter, r *http.Request) {
    fmt.Println("reading Datatype by ID")
}
func (d *Datatype) UpdateById(w http.ResponseWriter, r *http.Request) {
    fmt.Println("updating Datatype by ID")
}
func (d *Datatype) DeleteById(w http.ResponseWriter, r *http.Request) {
    fmt.Println("deleting Datatype by ID")
}
*/
