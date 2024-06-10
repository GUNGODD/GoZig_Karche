package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func crawlForEngine(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]
	fmt.Fprintf(w, " you used this Vin: %s", vin)

}

func main() {
	router := mux.NewRouter().StrictSlash(true) // StrictSlash is set to true by default what it does is it will not allow you to have a url with a trailing slash

	router.HandleFunc("/{vin}", crawlForEngine)
	fmt.Println("Listening on port 8080 {vin}")

	http.ListenAndServe(":8080", router)
}
