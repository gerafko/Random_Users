package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v4/stdlib"
	log "github.com/sirupsen/logrus"
	"net/http"
	"test/internal/handler"
	"test/internal/storage"
	"test/internal/storage/postgres"
	"test/pkg/randomuser"
	"time"
)

func main() {

	db, err := storage.NewPostgres(
		context.TODO(),
		"host=localhost port=5432 dbname=postgres user=postgres password=postgres",
		7*time.Second,
		3,
	)
	if err != nil {
		log.Fatalf("db connect error: %v ", err)
	}
	defer func() {
		dbErr := db.Close()
		if dbErr != nil {
			log.Errorf("db close error: %v ", dbErr)
		}
	}()

	// Общие зависимости для проекта
	randomuserClient := randomuser.NewClient("https://randomuser.me/api", &http.Client{})
	userRepository := postgres.NewUserRepository(db)

	// Initialize Gin router
	router := gin.Default()

	v1 := router.Group("api/v1")
	v1.Handle(http.MethodPatch, "/random-users/sync", handler.NewPatchRandomUsersSync(randomuserClient, userRepository).Handle)

	// Start HTTP server
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
