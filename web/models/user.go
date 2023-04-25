package models

type User struct {
	UserName string
	Age      int
}

type TForm struct {
	Name   string  `form:"username"`
	Age    int     `form:"userage"`
	Price  float64 `form:"price"`
	Single bool    `form:"single"`
}

type Tajax struct {
	User string `form:"username"`
	Age  int    `form:"age"`
}
