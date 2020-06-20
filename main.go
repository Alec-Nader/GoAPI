package main

import (
    "fmt"
    "log"
	"net/http"
	"encoding/json"
)
type Person struct {
    Name string `json:"Name"`
    Occupation string `json:"Occupation"`
}

var People []Person

func allPeople(w http.ResponseWriter, r *http.Request){
    fmt.Println("GET all people")
    json.NewEncoder(w).Encode(People)
}

func home(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Hello World!")
    fmt.Println("home")
}

func handleRequests() {
	http.HandleFunc("/", home)
	http.HandleFunc("/allPeople", allPeople)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {

	People = []Person{
        Person{Name: "Alec", Occupation: "Software Engineer"},
		Person{Name: "Jack", Occupation: "Chef"},
		Person{Name: "Tom",  Occupation: "Cashier"},
	}
	
    handleRequests()
}