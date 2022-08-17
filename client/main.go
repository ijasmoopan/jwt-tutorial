package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("mysupersecret")

func generateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Ijas Moopan"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		fmt.Println("Something went wrong", err)
		return "", err
	}
	return tokenString, err
}

func homePage(w http.ResponseWriter, r *http.Request) {
	validToken, err := generateJWT()
	if err != nil {
		fmt.Fprintln(w, "Error:", err)
		return
	}
	// client := &http.Client{}
	// req, _ := http.NewRequest("GET", "http://localhost:9090/", nil)
	// req.Header.Set("Token", validToken)
	// res, err := client.Do(req)
	// if err != nil {
	// 	fmt.Fprintln(w, err.Error())
	// }
	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Fprintln(w, err.Error())
	// }
	// fmt.Fprintln(w, string(body), validToken)
	fmt.Fprintln(w, validToken)
}

func HandleRequests() {
	http.HandleFunc("/", homePage)

	http.ListenAndServe(":8080", nil)
}

func main() {

	fmt.Println("My Client")

	HandleRequests()
}
