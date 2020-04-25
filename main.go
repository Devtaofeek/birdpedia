package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
"net/http")
type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

var birds []Bird
func newRouter() *mux.Router {
	r:= mux.NewRouter()
	r.HandleFunc("/hello", handler).Methods("GET")
	r.HandleFunc("/bird",getBirdHandler)
	staticFileDirectory:= http.Dir("./assets")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return  r
}

func main()  {
	r:= newRouter()
	http.ListenAndServe(":8080",r)
}

func handler(w http.ResponseWriter, r *http.Request)  {

	fmt.Fprintf(w, "Hello World!")
}

func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	//Convert the "birds" variable to json
	birds = append(birds, Bird{Species:"one specie",Description:"tolotolo"})
	birdListBytes, err := json.Marshal(birds)

	// If there is an error, print it to the console, and return a server
	// error response to the user
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// If all goes well, write the JSON list of birds to the response
	w.Write(birdListBytes)
}