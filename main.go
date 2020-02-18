package main

import (
	"github.com/anandagireesh/blog/database"
	"github.com/anandagireesh/blog/middlewares"
	"github.com/anandagireesh/blog/routes"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
	"net/http"
)

func  main()  {

	database.DbConnection()
	defer database.GetConnection().Close()

	//init
	//log.SetFormatter(&log.JSONFormatter{})
	//log.SetLevel(log.ErrorLevel)

	//set route



	n := negroni.New()
	n.UseFunc(middlewares.Auth)
	n.Use(negroni.NewRecovery())
	n.UseHandler(routes.MaiRoute())

	server := &http.Server{
		Addr:    "0.0.0.0:8006",
		Handler: n,
	}

	log.Info("Server Running")

	server.ListenAndServe()


}




