package cmd

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RouteSetup() {
	r := gin.New()
	v1 := r.Group("/api/v1")

	// Middlewares
	v1.Use(gin.Logger())
	v1.Use(gin.Recovery())
	v1.Use(gzip.Gzip(gzip.DefaultCompression))
	v1.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
	}))

	// routes

	//ping route
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	//welcome message route
	v1.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to Hotel Management Server Backend Server - V1.0.0")
	})
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found.",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})

	srvDetails := http.Server{
		Addr:        fmt.Sprintf(":8080"),
		Handler:     r,
		IdleTimeout: 120 * time.Second,
	}

	go func() {
		log.Println("SERVER STARTING ON PORT:8080")
		err := srvDetails.ListenAndServe()
		if err != nil {
			log.Printf("ERROR STARTING SERVER: %v", err)
			os.Exit(1)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Printf("Closing now, We've gotten signal: %v", sig)

	ctx := context.Background()
	srvDetails.Shutdown(ctx)
}
