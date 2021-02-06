package model

type Customers struct {
	Customers []Customer `json:"customers"`
}

type Customer struct {
	CustomerForCreate
	ID   int    `json:"id" example:"1"` // Customer ID
	Hash string `swaggerignore:"true"`
}

type CustomerForCreate struct {
	CustomerForUpdate
	Username string `json:"username" binding:"required" example:"choo" maxLength:"255"` // Customer Username
}

type CustomerForUpdate struct {
	Firstname string `json:"firstname" binding:"required" example:"Choopong" maxLength:"255"`   // Customer Firstname
	Lastname  string `json:"lastname" binding:"required" example:"Choosamer" maxLength:"255"`   // Customer Lastname
	Email     string `json:"email" binding:"required" example:"choo@gmail.com" maxLength:"255"` // Customer E-mail
	Gender    string `json:"gender" example:"male" enums:"male,female"`                         // Customer Gender
}
