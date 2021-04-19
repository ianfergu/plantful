package main

import (
    "fmt"
    "log"
    "net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

    router.HandleFunc("/", homePage)
	router.HandleFunc("/plants", returnAllPlants)
	router.HandleFunc("/plants/{id}", returnSinglePlant)
    log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	Plants = []Plant{
		Plant{Id: "1", Name: "Green one", Desc: "Its fluffy", Content: "hello"},
		Plant{Id: "2", Name: "Blue", Desc: "Its small", Content: "hello smello"},
	}
	handleRequests()
}

var Plants []Plant

type Plant struct {
	Id string `json:"id"`
    Name string `json:"Name"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

type Alert struct {
	Id string `json:"id"`
	Name string `json:"Name"`
	Desc string `json:"desc"`
	Update int `json:"update"`
}


func returnAllPlants(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Plants)
}

func returnSinglePlant(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, plant := range Plants {
        if plant.Id == key {
            json.NewEncoder(w).Encode(plant)
        }
    }

}