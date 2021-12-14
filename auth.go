package main

// "github.com/phc-dm/auth-poisson/user"

// AuthService rappresenta un servizio di autenticazione
// type AuthService interface {
// 	GetUsers() []User

// 	GetUser(username string) User

// 	// LoginUser if successful returns the token for this user that will be stored in an HTTP cookie.
// 	LoginUser(username, password string) (string, error)
// }

// LdapService ...
type LdapService struct {
	URL string
}

// FakeService ...
type FakeService struct {
	URL string
}

// NewAuthenticationService crea un nuovo servizio di autenticazione e controlla se è attivo
// func NewAuthenticationService(url string) (*LdapService, error) {
// 	service := new(LdapService)
// 	service.URL = url

// 	res, err := service.Get("status")

// 	if err != nil {
// 		return nil, err
// 	}

// 	status, _ := ioutil.ReadAll(res.Body)

// 	if string(status) != "true" {
// 		log.Fatalf("Authentication service isn't online, status: '%s'", status)
// 	}

// 	return service, nil
// }
