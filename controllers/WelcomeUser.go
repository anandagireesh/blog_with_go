package controllers

import (
	//"fmt"
	"net/http"
	"log"
)

func WelcomeUser(w http.ResponseWriter, r *http.Request) {

	//db.GetConnection()

	//keys := r.Header.Get("authorization")



	if m := r.Context().Value("userid"); m != nil {
		if value, ok := m.(string); ok {
			usrid := value
			log.Println("user id from context variable: "+ usrid)
			//fmt.Println(usrid)
		}
	}







}
