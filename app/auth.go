package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

type Claims struct {
	ID uint64 `json:"id,omitempty"`
	jwt.StandardClaims
}

func CreateToken(userid uint64) (string, error) {
	var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "ashssfksd") //this should be in an env file
	atClaims := &Claims{
		ID: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().UTC().Unix() + 15000,
		},
	}
	// atClaims := jwt.MapClaims{"authorized": true, "user_id": userid, "exp": time.Now().Add(time.Minute * 30).Unix()}
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func AuthMiddleware(h http.HandlerFunc) http.HandlerFunc {
	fmt.Println("auth")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var access_token string = mux.Vars(r)["access_token"]

		tkn, _ := jwt.ParseWithClaims(access_token, &Claims{}, func(token *jwt.Token) (interface{}, error) {

			return []byte(access_token), nil
		})
		if claimsTkn, ok := tkn.Claims.(*Claims); ok {
			log.Printf("id: %d", claimsTkn.ID)
			if claimsTkn.ExpiresAt < time.Now().UTC().Unix() {
				log.Printf("não validado %d", time.Now().UTC().Unix())

			}
			// log.Println(ok)
		} else {
			log.Println("não validado")
		}
		log.Println("middleware", r.URL)
		// h.ServeHTTP(w, r)
	})
}
