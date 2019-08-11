package main

import (
	//"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"flightsearch/handlers"
	"flightsearch/token"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load env variables: %s", err)
	}

	//	t = time.Now()
	token.GetToken()

}

func main() {
	r := gin.Default()
	r.GET("/flight-destinations", handlers.FlightDestinations)
	r.GET("/access-token", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Access Token": token.Token.AccessToken})
	})

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//signal.Notify(c, syscall.SIGINT)
	go token.TokenRefresh()
	//go func() {
	//	<-c
	//}()
	r.Run(":8080")

}
