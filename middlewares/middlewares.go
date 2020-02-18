package middlewares

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	//"fmt"
	"strings"
	//"github.com/dgrijalva/jwt-go"
	"net/http"
)

func Auth(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	if r.URL.Path == "/api/user/login" || r.URL.Path == "/api/user/register" || r.URL.Path == "/api/user/viewblogs"{
		next(rw, r)
		return
	}


	//get the header token

	keys := r.Header.Get("Authorization")

	toknString := strings.Split(keys , " ")



	tokenString := toknString[1]

	//s

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("ironman"), nil
	})

	if token.Valid {
		log.Println("You look nice today")
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c := context.Background()

			ctx :=context.WithValue(c, "userid", claims["user"])

			req := r.WithContext(ctx)
			next.ServeHTTP(rw, req)

			//next(rw, r)

		} else {
			fmt.Println(err)
		}
		//ctx := context.WithValue(context.Background(), "userid", userid)
		//fmt.Println(ctx.Value("userid"))

		return
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			log.Println("That's not even a token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			log.Println("Timing is everything")
		} else {
			log.Println("Couldn't handle this token:", err)
		}
	} else {
		log.Println("Couldn't handle this token:", err)
	}



	next(rw, r)
	// do some stuff after
}

