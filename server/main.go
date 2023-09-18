package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Post struct {
	ID      int    `json:"postID"`
	User    string `json:"user"`
	Message string `json:"message"`
}

var Posts []Post

func postHandler(w http.ResponseWriter, r *http.Request) {
	splitURL := strings.Split(r.URL.Path, "/")
	if len(splitURL) < 3 {
		http.Error(w, "Invalid URL", http.StatusBadRequest)
		return
	}

	idString := splitURL[2]

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

	fmt.Fprint(w, "Post made succsessfully!\nPost:\nUser: %s\nMessage: %s", retrievedData.User, retrievedData.Message)

}

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint hit: Homepage")
}

func handleReq() {
	http.HandleFunc("/", handleHomepage)
	http.HandleFunc("/post", postHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func main() {
	handleReq()
}
