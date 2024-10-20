package main

import (
	"database/sql"
	"log"
	"net/http"
	"ramori/internal/controllers"
	"ramori/internal/repositories"
	"ramori/internal/routes"
	"ramori/internal/usecases"

	_ "github.com/go-sql-driver/mysql"
	"ramori/pkg/database"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	userRepo := repositories.NewMySQLUserRepository(db)

	userUseCase := usecases.NewUserUseCase(userRepo)

	userController := &controllers.UserController{
		UserUseCase: userUseCase,
	}

	router := routes.SetupRoutes(userController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
