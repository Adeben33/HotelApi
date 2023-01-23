package cmd

import (
	"context"
	"fmt"
	"github.com/adeben33/HotelApi/cmd/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var v1 *gin.RouterGroup

func Setup() {
	port := os.Getenv("port")
	if port != "" {
		port = ":8080"
	}
	//route
	route := gin.New()
	//route = route.Group("/api/v1")
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	routes.UserRoutes(route)
	routes.PaymentRoutes(route)
	routes.ApartmentRoutes(route)
	routes.BookingRoutes(route)
	routes.ReviewsRoutes(route)
	//routes.DiscountsRoutes(route)

	//ping
	v1.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	//	No Routes
	route.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"name":    "Not Found",
			"message": "Page not found",
			"code":    404,
			"status":  http.StatusNotFound,
		})
	})
	srvDetails := http.Server{
		Addr:        fmt.Sprintf(":8080"),
		Handler:     route,
		IdleTimeout: 120 * time.Second,
	}
	go func() {
		log.Println("Server Starting on port:8080")
		err := srvDetails.ListenAndServe()
		if err != nil {
			log.Printf("Error starting server:%v", err)
			os.Exit(1)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Printf("Closing now, We've gotten signal: %v", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srvDetails.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	log.Println("Server existing")
}
