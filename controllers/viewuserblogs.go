package controllers

import (
	"encoding/json"
	db "github.com/anandagireesh/blog/database"
	"net/http"
	"log"
)

type AllUserBlog struct {

	Id string `json:"id"`
	BlogTitle string `json:"title"`
	UserBlogs string `json:"blog"`
}

type UserResponse struct {
	Code   int         `json:"code"`   // 200 , 400
	Status string      `json:"status"` // "Ok" "Error"
	Error  string      `json:"error"`
	Data   interface{} `json:"data"`
}

func ViewUserBlogs(w http.ResponseWriter, r *http.Request)  {

	db.GetConnection()


	if m := r.Context().Value("userid"); m != nil {
		if value, ok := m.(string); ok {
			usrid := value
			log.Println("user id from context variable: "+ usrid)

			var Bloguser []AllUserBlog


			rows, err := db.Db.Query("SELECT id blog_heading,blog FROM user_blog WHERE user_id = '"+usrid+"'")
			if err != nil {
				// handle this error better than this
				panic(err)
			}
			defer rows.Close()
			for rows.Next() {
				var id  string
				var blog_heading string
				var user_blog string
				err = rows.Scan(&id, &blog_heading,&user_blog)
				if err != nil {
					// handle this error
					panic(err)
				}
				//fmt.Println(id, firstName)
				Bloguser = append(Bloguser, AllUserBlog{Id:id, BlogTitle: blog_heading, UserBlogs: user_blog})

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
			//fmt.Println(usrid)
		}
	}



}
