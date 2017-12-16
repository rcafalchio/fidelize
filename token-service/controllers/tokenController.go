package controllers

import (
	// Standard packages
	"net/http"
	"time"

	//Third party packages
	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/julienschmidt/httprouter"
)

type (
	// AuthController represents the controller for authentication
	AuthController struct {
		keyword []byte
	}
)

//Set up a global access object
var auth = NewAuthController()

// Set up a global key for jwt
var keyword = []byte("secret")

// NewAuthController provides a reference to a AuthController
func NewAuthController() *AuthController {
	return &AuthController{keyword}
}

// ValidateToken validate token from token request, returning if the token is valid and a msgError
func (ac AuthController) ValidateToken(r *http.Request) (bool, string) {

	//Getting the token from request
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			return ac.keyword, nil
		})
	//If there is no error, then the token is valid
	if err == nil {
		if token.Valid {
			return true, ""
		}
		return false, "Token is not valid"
	}
	return false, "Unauthorized access to this resource"

}

// Create retrieves a valid token to the client
func (ac AuthController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	// Create a new token
	token := jwt.New(jwt.SigningMethodHS256)

	// Create a map for our claim
	claims := token.Claims.(jwt.MapClaims)

	// Set Claims
	claims["admin"] = true
	claims["name"] = "Pedro Ribeiro"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	// Sign the token with keyword
	tokenString, _ := token.SignedString(ac.keyword)

	// Return the token to the client
	w.Write([]byte(tokenString))
}
