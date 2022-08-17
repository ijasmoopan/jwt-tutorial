package main

import (
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var SigningKey = []byte("mysupersecret")

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Secret Information")
}

func isAuthorized(endpoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// if r.Header["Token"] != nil {
		// 	token, err := jwt.Parse(r.Header["Token"][0], func (token *jwt.Token)(interface{}, error){
		// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		// 			return nil, fmt.Errorf("There was an error")
		// 		}
		// 		return SigningKey, nil
		// 	})
		// 	if err != nil {
		// 		fmt.Fprintln(w, err.Error())
		// 	}
		// 	if token.Valid {
		// 		endpoint(w, r)
		// 	}
		// } else {
		// 	fmt.Fprintln(w, "Not Authorized")
		// }
		var reqToken string
		if r.Header["Authorization"] != nil {
			reqToken = r.Header.Get("Authorization")
			fmt.Println(reqToken)
			splitToken := strings.Split(reqToken, "Bearer ")
			reqToken = splitToken[1]
		}
		token, err := jwt.Parse(reqToken, func(token *jwt.Token)(interface{}, error){
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error")
			}
			return SigningKey, nil
		})
		if err != nil {
			fmt.Fprintln(w, err.Error())
		}
		if token.Valid {
			endpoint(w, r)
		} else {
			fmt.Fprintln(w, "Not Authorized")
		}
	})
}

func handleRequests() {
	http.Handle("/", isAuthorized(homePage))

	http.ListenAndServe(":9090", nil)
}

func main() {
	fmt.Println("My server")
	handleRequests()
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NjA3MTUxMjMsInVzZXIiOiJJamFzIE1vb3BhbiJ9.rexduXX_Igc7kFn5uzPUBXmlvGJhDQcz-EhUH5UGEws
// eyPhanfgoiiejvllz.cv.psodfihvjbapodsvja[sdodivj/aodivjqadovnjalscvj
