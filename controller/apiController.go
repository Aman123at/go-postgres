package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	database "github.com/Aman123at/go-postgres/db"
	"github.com/Aman123at/go-postgres/model"
	"github.com/gorilla/mux"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	posts := database.GetAllPostsFromDB()

	if len(posts) == 0 {
		json.NewEncoder(w).Encode(model.ResponseWithData{Success: true, Data: []any{}})
		return
	}

	json.NewEncoder(w).Encode(model.ResponseWithData{Success: true, Data: posts})
}
func GetPostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, parseErr := strconv.Atoi(params["id"])

	if parseErr != nil {
		log.Fatalf("Error: Unable to parse post id %v", parseErr.Error())
	}

	post := database.GetOnePost(id)

	json.NewEncoder(w).Encode(model.ResponseWithData{Success: true, Data: post})
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post model.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		log.Fatalf("Error : Unable to decode user post request %v", err.Error())
	}

	if post.IsEmpty() {
		json.NewEncoder(w).Encode(model.Response{Success: false, Message: "Title, Content and Author is required!!"})
		return
	}

	count := database.InsertNewPost(&post)

	message := fmt.Sprintf("Successfully Inserted recordsin posts table with id %v", count)

	json.NewEncoder(w).Encode(model.Response{Success: true, Message: message})
}

func DeletePostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Error: Unable to parse postId to int %v", err.Error())
	}

	database.DeleteOnePost(id)

	msgStr := fmt.Sprintf("Post with id %v , deleted successfully.", id)

	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: msgStr,
	})
}
func DeleteAllPostsInBulk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	database.DeleteAllPosts()

	msgStr := fmt.Sprintln("All Posts Deleted Successfully.")

	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: msgStr,
	})
}

func UpdatePostById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, convErr := strconv.Atoi(params["id"])

	if convErr != nil {
		log.Fatalf("Error: Unable to convert string to int %v", convErr.Error())
	}

	var post model.Post

	decodeErr := json.NewDecoder(r.Body).Decode(&post)

	if decodeErr != nil {
		log.Fatalf("Error: Unable to decode post %v", decodeErr.Error())
	}

	database.UpdatePost(id, post)

	json.NewEncoder(w).Encode(model.Response{
		Success: true,
		Message: "Post updated Successfully.",
	})
}
