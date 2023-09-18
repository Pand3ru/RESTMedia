package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type Post struct {
	ID      int    `json:"postID"`
	User    string `json:"user"`
	Message string `json:"message"`
}

var Posts []Post // "Database"

// creates a POST

func createPostHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	var retrievedData Post
	idString := vars["id"]

	id, errS := strconv.Atoi(idString)

	if errS != nil {
		http.Error(w, "ID has to be an Integer", http.StatusForbidden) // ID is non int.
		return
	} else {
		retrievedData.ID = id // assign struct the ID manually
	}

	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&retrievedData); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		Posts = append(Posts, retrievedData) // Instead of appending Slice, a Database implementation would be prefered. Maybe achievable through another API

		//fmt.Fprintf(w, "Post made successfully!\nPost:\n\tUser: %s\n\tMessage: %s\n\tID: %d\n", retrievedData.User, retrievedData.Message, retrievedData.ID)

		jsonData, err := json.Marshal(retrievedData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)

	} else if r.Method == "GET" {
		for i := 0; i < len(Posts); i++ {
			if Posts[i].ID == id {
				fmt.Fprintf(w, "Post:\n\tUser: %s\n\tMessage: %s\n\tID: %d\n", Posts[i].User, Posts[i].Message, Posts[i].ID)
				return
			}
		}
		http.Error(w, "Post not found", http.StatusNotFound)
	}
}

// Handle Homepage

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint hit: Homepage")
}

// Handles URI requests

func handleReq() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleHomepage)
	r.HandleFunc("/post/{id}", createPostHandler)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"}) // Allow requests from any origin

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
func main() {
	handleReq()
}
