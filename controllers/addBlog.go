package controllers

import (
	"encoding/json"
	db "github.com/anandagireesh/blog/database"
	"github.com/anandagireesh/blog/models"
	"io/ioutil"
	"log"
	"net/http"

)

func AddBlog(w http.ResponseWriter, r *http.Request)  {

	db.GetConnection()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var blog models.UserBlog

	err = json.Unmarshal(b, &blog)
	if err != nil {
		log.Println(err)
	}

	if m := r.Context().Value("userid"); m != nil {
		if value, ok := m.(string); ok {
			usrid := value

			insert, err := db.Db.Prepare("INSERT INTO user_blog(user_id, blog_heading, blog) VALUES ( ? , ?, ?)")

			//if there is an error inserting, handle it
			if err != nil {
				panic(err.Error())
			}
			insert.Exec(usrid, blog.BlogHeading, blog.BlogText)

			//log.Println(qry)
			//be careful deferring Queries if you are using transactions

			log.Println("INSERT: Userid : " + usrid + " | blog Heading: " + blog.BlogHeading + " | blog: " + blog.BlogText)

		}
	}









}
