package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"ramori/internal/controllers"
)

func main() {
	dsn := "horacio:12345678@tcp(mysql:3306)/ramori"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	userController := &controllers.UserController{DB: db}

	http.HandleFunc("/users", userController.CreateUser)

	log.Println("Server running in port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
