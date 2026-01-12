package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/abdoulousseini2028-droid/taskmaster-api/internal/config"
	"github.com/abdoulousseini2028-droid/taskmaster-api/internal/handlers"
	"github.com/abdoulousseini2028-droid/taskmaster-api/internal/repository"
)

func main() {
	cfg := config.Load()
	pool, err := pgxpool.New(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	defer pool.Close()
	if err := pool.Ping(context.Background()); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}
	log.Println("Successfully connected to database!")
	
	taskRepo := repository.NewTaskRepository(pool)
	taskHandler := handlers.NewTaskHandler(taskRepo)
	
	router := setupRouter(taskHandler, cfg)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}
	go func() {
		log.Printf("Server starting on port %s\n", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v\n", err)
		}
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server exited")
}

func setupRouter(taskHandler *handlers.TaskHandler, cfg *config.Config) *gin.Engine {
	if cfg.Environment == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "taskmaster-api",
			"time":    time.Now().Format(time.RFC3339),
		})
	})
	
	v1 := router.Group("/api/v1")
	{
		tasks := v1.Group("/tasks")
		{
			tasks.POST("", taskHandler.CreateTask)
			tasks.GET("", taskHandler.ListTasks)
			tasks.GET("/:id", taskHandler.GetTask)
			tasks.PUT("/:id", taskHandler.UpdateTask)
			tasks.DELETE("/:id", taskHandler.DeleteTask)
		}
	}
	return router
}
