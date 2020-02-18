package routes

import (
	//"github.com/anandagireesh/blog/middlewares"
	"github.com/gorilla/mux"
	//"github.com/urfave/negroni"
	//"net/http"

	"github.com/anandagireesh/blog/controllers"

)


func MaiRoute() *mux.Router{

	r := mux.NewRouter().StrictSlash(false)
	//home handler
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")

	//register user
	r.HandleFunc("/api/user/register", controllers.UserRegister).Methods("POST")

	// user login
	r.HandleFunc("/api/user/login", controllers.UserLogin).Methods("POST")
	r.HandleFunc("/api/user/loggedin", controllers.WelcomeUser).Methods("GET")

	//blogs
	r.HandleFunc("/api/user/addblog", controllers.AddBlog).Methods("POST")
	r.HandleFunc("/api/user/viewblogs", controllers.ViewBlogs).Methods("GET")
	r.HandleFunc("/api/user/viewuserblogs", controllers.ViewUserBlogs).Methods("GET")
	r.HandleFunc("/api/user/viewsingleblog", controllers.ViewSingleBlog).Methods("POST")
	r.HandleFunc("/api/user/updateblog", controllers.UpdateBlog).Methods("PUT")





	return r
}
