package controllers

//noinspection GoUnresolvedReference
import (
	"database/sql"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"

	//"database/sql"
	"encoding/json"
	db "github.com/anandagireesh/blog/database"
	"github.com/anandagireesh/blog/models"
	"golang.org/x/crypto/bcrypt"

	//"time"

	//"github.com/anandagireesh/blog/models"
	"io/ioutil"
	"log"
	"net/http"
)
var mySigningKey = []byte("ironman")
func GenerateJWT(email string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = email
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}

	return tokenString, nil
}

type myToken struct {
	Token string
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

func UserLogin(w http.ResponseWriter, r *http.Request)  {

	db.GetConnection()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var msg models.UserLogin

	err = json.Unmarshal(b, &msg)
	if err != nil {
		log.Println(err)
	}

	log.Println(msg.Email)
	log.Println(msg.Password)

	passwordenc := msg.Password
	bpassword := []byte(passwordenc)
	hash, err := bcrypt.GenerateFromPassword(bpassword, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	encPassword := string(hash)
	log.Println(encPassword)

	var password, id string

	sqlStatement := "SELECT id, password FROM users WHERE email= '"+msg.Email+"'"

	log.Println(sqlStatement)


	var row  =db.Db.QueryRow(sqlStatement)
	switch err := row.Scan(&id, &password); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		log.Println(password)

			userid := id
		pwdMatch := comparePasswords(password, bpassword)
		//fmt.Println("Passwords Match?", password)
		if pwdMatch == false {
			//log.Println(pwdMatch)
			log.Println("User does not exist")
			return
		}

		//log.Println("Passwords Matched", password)

		// token generation

	tokenString, err := GenerateJWT(userid)
		if err != nil {

			fmt.Println("Error generating Token string")
		}
	
	fmt.Println(tokenString)



		token:= &myToken{Token: tokenString}
		tokendata, err := json.Marshal(token)
		if err != nil{
			panic(err)
		}

		//Set Content-Type header so that clients will know how to read response
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		//Write json response back to response
		w.Write(tokendata)




	default:
		panic(err)
	}



	
}
