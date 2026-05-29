package handler


import (
	"fmt"
	"net/http"
)

type Order struct {}


func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Order Created")
}


func (o *Order) UpdateById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Order Update by ID")
}


func (o *Order) DeleteById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Delete order by ID")
}


func (o *Order) GetById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Get Order By ID")
}

func (o *Order) List(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "List Order")
}
