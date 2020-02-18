package models



type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"firstname",valid:"alpha,required"`
	LastName string `json:"lastname",valid:"alphanum"`
	Email string `json:"email",valid:"email,required"`
	Password string `json:"password",valid:"matches(^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#\$%\^&\*])(?=.{8,})),required"`
	Country string `json:"country",valid:"alpha,required"`
	Status string `json:"status"`
}

type UserLogin struct {

	Email string `json:"email"`
	Password string `json:"password"`

}



