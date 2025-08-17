package main

import (
	"database/sql"
	"log"
	"net/http"

	"gchan/handlers"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// DB
	db, err := sql.Open("sqlite3", "db/imageboard.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Rutas

	http.HandleFunc("/api/boards", func(w http.ResponseWriter, r *http.Request) {
		handlers.BoardsHandler(w, r, db)
	})

	http.HandleFunc("/api/boards", func(w http.ResponseWriter, r *http.Request) {
		handlers.BoardsThreadsHandler(w, r, db)
	})

	// Rutas crud

	http.HandleFunc("/api/posts/create", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreatePostHandler(w, r, db)
	})
	http.HandleFunc("/api/posts/update", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdatePostHandler(w, r, db)
	})
	http.HandleFunc("/api/posts/delete", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeletePostHandler(w, r, db)
	})

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
