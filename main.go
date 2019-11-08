package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	certFile = "/opt/certificates/bb-qa-pb.avaldigitallabs.com.crt"
	keyFile  = "/opt/certificates/bb-qa-pb.avaldigitallabs.com.key"
)

// TokenEvent struct
type TokenEvent struct {
	URL          string `json:"url"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

// TokenResponse struct
type TokenResponse struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	ConsentedOn int64  `json:"consented_on"`
	Scope       string `json:"scope"`
}

func main() {
	lambda.Start(handlerToken)
}

func handlerToken(event TokenEvent) (*TokenResponse, error) {
	response := &TokenResponse{}
	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		log.Fatal(err)
		return response, err
	}

	tlsConfig := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}
	tlsConfig.BuildNameToCertificate()
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	body := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=%s", event.ClientID, event.ClientSecret, event.Scope)
	fmt.Println("body", body, "\n")

	req, err := http.NewRequest("POST", event.URL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		log.Fatal(err)
		return response, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return response, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return response, err
	}

	json.Unmarshal(data, response)

	return response, nil
}
