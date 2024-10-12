package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"ramori/internal/controllers"
	"ramori/internal/repositories"
	"ramori/internal/usecases"
)

func main() {
	dsn := "user:password@tcp(localhost:3306)/dbname"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	userRepo := repositories.NewMySQLUserRepository(db)

	userUseCase := usecases.NewUserUseCase(userRepo)

	userController := &controllers.UserController{
		UserUseCase: userUseCase,
	}

	http.HandleFunc("/users", userController.CreateUser)

	log.Println("Server running in port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
