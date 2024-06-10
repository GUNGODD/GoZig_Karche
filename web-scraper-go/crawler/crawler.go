package crawler

import (
	"fmt"

	"github.com/gorilla/mux"
)

func CrawlerForEngine(w http.ResponseWritter, r *http.Request) {
	vars := mux.Vars(r)
	vin := vars["vin"]

	fmt.Fprintf(w, "you useed this vin: %s", vin)
}
