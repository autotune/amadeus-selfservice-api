package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"

	"flightsearch/token"
)

var client http.Client

func init() {
	client = http.Client{}
}

func FlightDestinations(c *gin.Context) {
	origin := c.Query("origin")

	u := url.URL{}
	u.Host = os.Getenv("BASE_URL")
	u.Path = "/v1/shopping/flight-destinations"
	u.Scheme = "https"
	q := u.Query()
	q.Set("origin", origin)
	u.RawQuery = q.Encode()

	uri := u.String()
	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.Token.AccessToken))
	//log.Printf("Request Header: %v", req.Header)
	if err != nil {
		log.Printf("Unable to build request for flight destinations: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Unable to get response for flight destinations: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	_ = ioutil.WriteFile("flight-destinations.json", data, 0664)
	c.JSON(http.StatusOK, gin.H{"data": string(data)})
}
