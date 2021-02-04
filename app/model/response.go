package model

type Response struct {
	Code   int    `json:code`   // Response Code
	Status string `json:status` // Response Status
}
