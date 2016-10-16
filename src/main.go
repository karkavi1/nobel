package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Winners structure
type Winners struct {
	Name        string `json:"name"`
	DateOfBirth string `json:"dob"`
	Place       string `json:"place"`
	Country     string `json:"country"`
	Category    string `json:"category"`
	Year        string `json:"year"`
	Synopsis    string `json:"synopsis"`
}

var nobelPrizeWinners = map[string]*Winners{
	"id_0": &Winners{Name: "Sir Chandrasekhara Venkata Raman", DateOfBirth: "11/07/1888", Place: "Tiruchirappalli, Tamil Nadu", Country: "India", Category: "Physics", Year: "1930", Synopsis: "for his work on the scattering of light and for the discovery of the effect named after him"},
	"id_1": &Winners{Name: "Albert Einstein", DateOfBirth: "03/14/1879", Place: "Ulm", Country: "Germany", Category: "Physics", Year: "1921", Synopsis: "for his services to Theoretical Physics, and especially for his discovery of the law of the photoelectric effect"},
}

// Handles request from RequestURI: /nobel/winners/list and
// returns list of Nobel Prize Winners as JSON Object
func getNobelWinnersList(httpRes http.ResponseWriter, httpReq *http.Request) {
	if hasAttributes(httpReq, httpRes) {
		nobelPrizeJSONResponse, error := json.Marshal(nobelPrizeWinners)
		writeHTTPResponseInWriter(httpRes, httpReq, nobelPrizeJSONResponse, error)
	}
}

// Handles request from RequestURI: /nobel/winners/fetch/{id} and
// returns respective Nobel Prize Winner as JSON Object
func getNobelWinnersByID(httpRes http.ResponseWriter, httpReq *http.Request) {
	if hasAttributes(httpReq, httpRes) {
		vars := mux.Vars(httpReq)
		mapKey := vars["id"]
		log.Printf("Request to fetch %s key from Nobel Prize Collection", mapKey)
		nobelPrizeJSONResponse, error := json.Marshal(nobelPrizeWinners[mapKey])
		writeHTTPResponseInWriter(httpRes, httpReq, nobelPrizeJSONResponse, error)
	}
}

// Validates Api-Key in request header
func hasAttributes(request *http.Request, response http.ResponseWriter) bool {

	log.Printf("Request for %s is in progress....", request.RequestURI)
	if request.Header.Get("Api-Key") != "nobelApp" {
		http.Error(response, "Request Header: Invalid Api-Key", http.StatusBadRequest)
		return false
	}
	return true
}

// Write HTTP Response and check marshaling errors (return HTTP STATUS 500 ) before it writes to ResponseWriter
func writeHTTPResponseInWriter(httpRes http.ResponseWriter, httpReq *http.Request, nobelPrizeWinnersResponse []byte, err error) {
	if err != nil {
		log.Println(err.Error())
		http.Error(httpRes, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Printf("Request %s Succesfully Completed", httpReq.RequestURI)
	httpRes.Header().Set("Content-Type", "application/json")
	httpRes.Write(nobelPrizeWinnersResponse)
}

// Handles request and routes to respective resource
func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	subRouter := router.PathPrefix("/nobel/winners").Subrouter()

	// Routes consist of a path and a handler function.
	subRouter.HandleFunc("/fetch/all", getNobelWinnersList).Methods("GET")
	subRouter.HandleFunc("/fetch/{id}", getNobelWinnersByID).Methods("GET")

	log.Print("Listening port 8081...")
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8081", router))
}

// program entry point
func main() {
	handleRequests()
}
