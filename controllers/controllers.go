package controllers

import (
	"net/http"

)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	//write response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Working"))

}

