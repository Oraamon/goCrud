package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"ramori/internal/controllers"
	"ramori/internal/repositories"
	"ramori/internal/routes"
	"ramori/internal/usecases"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load("cmd/configs/.env")

	if err != nil {
		log.Fatalf("Err .env: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

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

	router := routes.SetupRoutes(userController)
	log.Fatal(http.ListenAndServe(":8080", router))
}
