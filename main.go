package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Board struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
type Thread struct {
	ID        int    `json:"id"`
	Subject   string `json:"subject"`
	BoardID   int    `json:"board_id"`
	CreatedAt string `json:"created_at"`
}

func main() {
	// Abrir
	db, err := sql.Open("sqlite3", "db/imageboard.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Handler boards
	http.HandleFunc("/api/boards", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		rows, err := db.Query("SELECT id, name, description FROM boards")
		if err != nil {
			http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var boards []Board
		for rows.Next() {
			var b Board
			if err := rows.Scan(&b.ID, &b.Name, &b.Description); err != nil {
				http.Error(w, "Error al leer los resultados", http.StatusInternalServerError)
				return
			}
			boards = append(boards, b)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(boards)
	})
	// Handler threads
	http.HandleFunc("/api/boards/", func(w http.ResponseWriter, r *http.Request) {
		if !strings.HasSuffix(r.URL.Path, "/threads") {
			http.NotFound(w, r)
			return
		}

		// Ruta esperada: /api/boards/{id}/threads
		parts := strings.Split(r.URL.Path, "/")
		if len(parts) != 5 {
			http.NotFound(w, r)
			return
		}

		boardID, err := strconv.Atoi(parts[3])
		if err != nil {
			http.Error(w, "ID inválido", http.StatusBadRequest)
			return
		}

		// Obtener los threads
		rows, err := db.Query(`
		SELECT id, subject, board_id, created_at
		FROM threads
		WHERE board_id = ?
	`, boardID)
		if err != nil {
			log.Println("Error consultando threads:", err)
			http.Error(w, "Error al consultar threads", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var threads []Thread
		for rows.Next() {
			var t Thread
			if err := rows.Scan(&t.ID, &t.Subject, &t.BoardID, &t.CreatedAt); err != nil {
				http.Error(w, "Error leyendo los datos", http.StatusInternalServerError)
				return
			}
			threads = append(threads, t)
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(threads)
	})

	log.Println("Servidor corriendo en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
