package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// AuthenticationService è l'interfaccia del servizio di autenticazione sul server principale
type AuthenticationService struct {
	URL string
}

// NewAuthenticationService crea un nuovo servizio di autenticazione e controlla se è attivo
func NewAuthenticationService(url string) (*AuthenticationService, error) {
	service := new(AuthenticationService)
	service.URL = url

	res, err := service.Get("status")

	if err != nil {
		return nil, err
	}

	status, _ := ioutil.ReadAll(res.Body)

	if string(status) != "online" {
		log.Fatalf("Authentication service isn't online, status: '%s'", status)
	}

	return service, nil
}

// Get ...
func (service *AuthenticationService) Get(url string) (*http.Response, error) {
	return http.Get(service.URL + "/" + url)
}

// Authenticate ...
func (service *AuthenticationService) Authenticate(username, password string) bool {

	json, _ := json.Marshal(struct {
		Username, Password string
	}{username, password})

	res, _ := http.Post(service.URL+"/auth", "application/json", bytes.NewReader(json))

	result, _ := ioutil.ReadAll(res.Body)

	return string(result) == "true"
}
