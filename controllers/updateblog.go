package controllers

import (
	"encoding/json"
	db "github.com/anandagireesh/blog/database"
	"github.com/anandagireesh/blog/models"
	"io/ioutil"
	"net/http"
	"log"
)

//type UpdateBlogs struct {
//
//	Id string `json:"id"`
//	BlogTitle string `json:"title"`
//	UserBlogs string `json:"blog"`
//}

func UpdateBlog(w http.ResponseWriter, r *http.Request)  {
	db.GetConnection()

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var blogid models.Blog

	err = json.Unmarshal(b, &blogid)
	if err != nil {
		log.Println(err)
	}

	insForm, err := db.Db.Prepare("UPDATE user_blog SET blog_heading=?, blog=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(blogid.BlogTitle, blogid.UserBlogs, blogid.Id)
	log.Println("UPDATE: Title: " + blogid.BlogTitle + " | Blog: " + blogid.UserBlogs)


}
