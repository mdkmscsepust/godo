package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	application "godo/internal/application/user"
	domain "godo/internal/domain/user"
	infrastructure "godo/internal/infrastructure/user"
	"log"
	"net/http"
)

//import "godo/cmd"

func main() {
	//cmd.Execute()
	db, err := sql.Open("postgres", "user=masum password=12345678 dbname=go_db host=localhost port=5432")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	repo := infrastructure.NewPostgresUserRepo(db) 
	service := domain.NewService(repo)
	usecase := application.NewUseCase(service)

	http.HandleFunc("/register", func (w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		email := r.FormValue("email")

		u, err := usecase.RegisterUser(name, email)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "User registered: %s", u.Id)
	})

	

	log.Println("Server running on http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}