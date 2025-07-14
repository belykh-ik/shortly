package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"api/shorturl/broker"
	"api/shorturl/broker/handleMessage"
	"api/shorturl/internal/consts"
	"api/shorturl/internal/db"
	"api/shorturl/internal/handlers"
	"api/shorturl/internal/models"
	"api/shorturl/internal/service"
	"api/shorturl/internal/statistics"
	"api/shorturl/middleware"

	"github.com/joho/godotenv"
)

func main() {
	servers := []string{"kafka0:9092", "kafka1:9092"}
	err := godotenv.Load()
	if err != nil {
		panic("Error Load .env")
	}
	dbConf := models.Config{
		DSN:    os.Getenv("DSN"),
		Secret: os.Getenv("SECRET"),
	}
	wg := sync.WaitGroup{}
	//Connect to Database Postgres
	db := db.ConnectDb(&dbConf)

	//Create LinkDependence
	link := service.NewLinkDeps(db)

	//Create UserRepository
	userRepository := service.NewUserRepository(db)

	//Create StatisticRepository
	stat := statistics.NewStatisticsRepository(db)

	mux := http.NewServeMux()

	// Зависимость для обработки сообщений
	messageDeps := handleMessage.NewHandleMessageDeps(mux)

	//Register Routes
	handlers.RegisterRoutes(mux, &dbConf, link, stat)
	handlers.RegisterAuthRoutes(mux, &dbConf, userRepository)

	//Create Server
	server := http.Server{
		Addr:    ":8082",
		Handler: middleware.Cors(middleware.Logging(mux)),
	}

	fmt.Println("Server is listening on port 8082")
	wg.Add(1)
	go func() {
		defer wg.Done()
		err = server.ListenAndServe()
		if err != nil {
			fmt.Printf("Error %s", err)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumeer, err := broker.NewConsumer(servers, consts.GroupId, consts.ConsumerTopic, messageDeps)
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		log.Println("Start Consumer...")
		consumeer.Start()
	}()

	wg.Wait()
}
