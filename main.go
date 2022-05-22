package main

import (
	// "fmt"
	// "net/http"

	// "github.com/gin-gonic/gin"

	// "flightsearch/handlers"
	"flightsearch/token"
)

func init() {
	//	t = time.Now()
	token.GetToken()

}

func main() {
	// r := gin.Default()
	// r.GET("/flight-destinations", handlers.FlightDestinations)
	// r.GET("/access-token", func(c *gin.Context) {
        token.GetToken() 
	// })

	//c := make(chan os.Signal, 1)
	//signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	//signal.Notify(c, syscall.SIGINT)
	go token.TokenRefresh()
	//go func() {
	//	<-c
	//}()
	// r.Run(":8080")

}
