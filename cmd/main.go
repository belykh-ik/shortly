package main

import (
	"fmt"
	"net/http"
	"os"

	"api/shorturl/internal/db"
	"api/shorturl/internal/handlers"
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"api/shorturl/middleware"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error Load .env")
	}
	dbConf := models.Config{
		DSN:   os.Getenv("DSN"),
		TOKEN: os.Getenv("TOKEN"),
	}
	//Connect to Database Postgres
	db := db.ConnectDb(&dbConf)

	//Create LinkDependence
	link := service.NewLinkDeps(db)

	//Create UserRepository
	userRepository := service.NewUserRepository(db)

	mux := http.NewServeMux()

	//Register Routes
	handlers.RegisterRoutes(mux, link)
	handlers.RegisterAuthRoutes(mux, &dbConf, userRepository)

	//Create Server
	server := http.Server{
		Addr:    ":8082",
		Handler: middleware.Cors(middleware.Logging(mux)),
	}

	fmt.Println("Server is listening on port 8082")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error %s", err)
	}
}
