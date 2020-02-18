package models

type UserBlog struct {

	UserID string `json:"user_id"`
	BlogHeading string `json:"blogheading"`
	BlogText string `json:"blog"`

}

type SingleBlog struct {

	BlogId string `json:"blog_id"`

}

type Blog struct {

	Id string `json:"id"`
	BlogTitle string `json:"title"`
	UserBlogs string `json:"blog"`

}
