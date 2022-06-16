package main

type ContactModel struct {
	Name  string `json:"name" form:"name" query:"name"`
	Age   int    `json:"age" form:"age" query:"age"`
	Email string `json:"email" form:"email" query:"email"`
}
