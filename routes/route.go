package routes

import (
	"github.com/Aman123at/go-postgres/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/api/posts", controller.GetAllPosts).Methods("GET")

	r.HandleFunc("/api/post/{id}", controller.GetPostById).Methods("GET")

	r.HandleFunc("/api/post/{id}", controller.UpdatePostById).Methods("PUT")

	r.HandleFunc("/api/post", controller.CreatePost).Methods("POST")

	r.HandleFunc("/api/post/{id}", controller.DeletePostById).Methods("DELETE")

	r.HandleFunc("/api/posts/deleteAll", controller.DeleteAllPostsInBulk).Methods("DELETE")

	return r
}
