package controllers

import (
	"encoding/json"
	//"fmt"
	db "github.com/anandagireesh/blog/database"
	"net/http"
	"log"
)

type AllBlog struct {
	FirstName string `json:"firstname"`
	LastName string  `json:"lastname"`
	BlogTitle string `json:"title"`
	UserBlogs string `json:"blog"`
}

type Response struct {
	Code   int         `json:"code"`   // 200 , 400
	Status string      `json:"status"` // "Ok" "Error"
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func ViewBlogs(w http.ResponseWriter, r *http.Request)  {

	db.GetConnection()



	var Bloguser []AllBlog


	rows, err := db.Db.Query("SELECT users.firstname,users.lastname,user_blog.blog_heading,user_blog.blog FROM users INNER JOIN user_blog ON users.id=user_blog.user_id")
	if err != nil {
		// handle this error better than this
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		var firstname string
		var lastname string
		var blog_heading string
		var user_blog string
		err = rows.Scan(&firstname, &lastname,&blog_heading,&user_blog)
		if err != nil {
			// handle this error
			panic(err)
		}
		//fmt.Println(id, firstName)
		Bloguser = append(Bloguser, AllBlog{FirstName: firstname, LastName: lastname, BlogTitle: blog_heading, UserBlogs: user_blog})

	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	response := &Response{
		Code : http.StatusAccepted,
		Status: "ok",
		Data : Bloguser,
	}

	urlsJson, _ := json.Marshal(response)
	log.Println(urlsJson)
	//log.Println(Bloguser)

	w.Write(urlsJson)
	
}
