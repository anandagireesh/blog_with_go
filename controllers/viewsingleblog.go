package controllers

import (
	"encoding/json"
	db "github.com/anandagireesh/blog/database"
	"github.com/anandagireesh/blog/models"
	"io/ioutil"
	"net/http"
	"log"
)

type SingleBlog struct {

	Id string `json:"id"`
	BlogTitle string `json:"title"`
	UserBlogs string `json:"blog"`
}

func ViewSingleBlog(w http.ResponseWriter, r *http.Request)  {

	db.GetConnection()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var blogid models.SingleBlog

	err = json.Unmarshal(b, &blogid)
	if err != nil {
		log.Println(err)
	}

	var Bloguser []SingleBlog

	id := blogid.BlogId

	rows, err := db.Db.Query("SELECT blog_heading,blog FROM user_blog WHERE id = '"+id+"'")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var blog_heading string
		var user_blog string
		err = rows.Scan( &blog_heading,&user_blog)
		if err != nil {
			// handle this error
			panic(err)
		}
		//fmt.Println(id, firstName)
		Bloguser = append(Bloguser, SingleBlog{Id:id , BlogTitle: blog_heading, UserBlogs: user_blog})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	response := &UserResponse{
		Code : http.StatusAccepted,
		Status: "ok",
		Data : Bloguser,
	}

	urlsJson, _ := json.Marshal(response)
	log.Println(urlsJson)
	//log.Println(Bloguser)

	w.Write(urlsJson)

}
