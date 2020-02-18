package controllers

import (
	"encoding/json"
	db "github.com/anandagireesh/blog/database"
	"github.com/anandagireesh/blog/models"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"

	//"fmt"
	"log"
	"net/http"

	"github.com/asaskevich/govalidator"
)

func UserRegister(w http.ResponseWriter, r *http.Request)  {

db.GetConnection()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Unmarshal
	var msg models.User
	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	result, err := govalidator.ValidateStruct(models.User{})
	if err != nil {
		println("error: " + err.Error())
		return
	}
	println(result)

	password := msg.Password
	bpassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bpassword, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	Fname := msg.FirstName
	Lname := msg.LastName
	email := msg.Email
	encPassword := string(hash)
	country :=msg.Country



	insert, err := db.Db.Prepare("INSERT INTO users(firstname, lastname, email, password, country) VALUES ( ? , ? , ? , ? , ?)")

	//if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	insert.Exec(Fname,Lname,email, encPassword,country)
	// be careful deferring Queries if you are using transactions
	
	log.Println("INSERT: Name: " + msg.FirstName +" " + msg.LastName + " | Email: " + msg.Email + " | Hashed Password: " + encPassword + " | Password: " + msg.Password+ " | Country: " + msg.Country)


}
