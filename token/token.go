package token

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

type AccessToken struct {
	Type            string `json:"type"`
	Username        string `json:"username"`
	ApplicationName string `json:"application_name"`
	ClientId        string `json:"client_id"`
	TokenType       string `json:"token_type"`
	AccessToken     string `json:"access_token"`
	ExpiresIn       int    `json:"expires_in"`
	State           string `json:"state"`
	Scope           string `json:"scope"`
}

var Token *AccessToken

func GetToken() {

	var (
		apiKey    string = os.Getenv("AMADEUS_CLIENT_ID")
		apiSecret string = os.Getenv("AMADEUS_CLIENT_SECRET")
		baseUrl   string = os.Getenv("AMADEUS_CLIENT_BASE_URL")
	)

	u := url.URL{}
	u.Host = baseUrl
	u.Path = "/v1/security/oauth2/token"
	u.Scheme = "https"
	uri := baseUrl//u.String()

	v := url.Values{}
	v.Add("grant_type", "client_credentials")
	v.Add("client_id", apiKey)
	v.Add("client_secret", apiSecret)

	client := http.Client{}
	req, err := http.NewRequest("POST", uri, strings.NewReader(v.Encode()))
	if err != nil {
		log.Fatalf("Error in getting access token req: %s", uri)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Unable to get access token: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Error reading response from Token Response Body: %s", err)
	}

	err = ioutil.WriteFile("token.json", body, 0664)
	if err != nil {
		log.Println("Unable to write to token.json file")
	}

	//token := &token.AccessToken{}
	log.Println(string(body))
	err = json.Unmarshal(body, &Token)
	if err != nil {
		log.Printf("Error in accessing Access Token body in response: %s", err)
		return
	}

	return
}

func checkTokenFileExist() bool {

	_, err := os.Stat("token.json")

	if os.IsExist(err) {
		return false
	}

	return true
}

func checkTokenExist() bool {
	f, err := ioutil.ReadFile("token.json")
	if err != nil {
		log.Println("Error in reading token.json file!")
		return false
	}

	err = json.Unmarshal(f, &Token)

	if Token.AccessToken == "" {
		return false
	}

	return true
}

func checkTokenExpired(at *AccessToken) bool {
	if at.State == "expired" {
		return true
	}

	return false
}

func TokenRefresh() {
	for {
		time.Sleep(time.Duration(Token.ExpiresIn-100) * time.Second)
		GetToken()
	}
}
