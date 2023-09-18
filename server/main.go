package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Post struct {
	ID      int    `json:"postID"`
	User    string `json:"user"`
	Message string `json:"message"`
}

var Posts []Post

func postHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idString := vars["id"]

	var retrievedData Post
	if err := json.NewDecoder(r.Body).Decode(&retrievedData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, errS := strconv.Atoi(idString)

	if errS != nil {
		return
	} else {
		retrievedData.ID = id
	}

	Posts = append(Posts, retrievedData)

	fmt.Fprintf(w, "Post made successfully!\nPost:\n\tUser: %s\n\tMessage: %s\n\tID: %d\n", retrievedData.User, retrievedData.Message, retrievedData.ID)

}

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint hit: Homepage")
}

func handleReq() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHomepage)
	r.HandleFunc("/post/{id}", postHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
func main() {
	handleReq()
}
