package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person is the JSON structure I would like to export
type Person struct {
	ID         string `json:"Id"`
	Name       string `json:"Name"`
	Occupation string `json:"Occupation"`
}

// People - the fake dataset I will use in place of a database.
var People []Person

// home function handles the homepage request and returns a string to the page.
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

// AllPeople is a function which is called when the endpoint /allPeople is requested. AllPeople returns the encoded JSON of the People object.
func allPeople(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GET all people")
	json.NewEncoder(w).Encode(People)
}

// returnPerson function utilizes the mux router and the query parameter to select and return the proper person by ID.
func returnPerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// _ can be used as a throw away variable for iteration. Iterate through the array of people and find the one that matches the query parameter ID.
	for _, person := range People {
		if person.ID == key {
			json.NewEncoder(w).Encode(person)
		}
	}
}

func createPerson(w http.ResponseWriter, r *http.Request) {
	// use ioutil to parse the POST request body.
	reqBody, _ := ioutil.ReadAll(r.Body)
	var person Person
	json.Unmarshal(reqBody, &person)

	People = append(People, person)

	json.NewEncoder(w).Encode(person)
}

// handleRequests handles HTTP requests by use of the gorilla/mux router instead of the http/net.
func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", home)
	myRouter.HandleFunc("/allPeople", allPeople)
	myRouter.HandleFunc("/person", createPerson).Methods("POST")
	myRouter.HandleFunc("/person/{id}", returnPerson)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")

	People = []Person{
		Person{ID: "1", Name: "Alec", Occupation: "Software Engineer"},
		Person{ID: "2", Name: "Jack", Occupation: "Chef"},
		Person{ID: "3", Name: "Tom", Occupation: "Cashier"},
	}

	handleRequests()
}
