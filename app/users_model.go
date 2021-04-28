package main

type user struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	AccessToken string `json:"accesstoken"`
}
type users struct {
	Users []user `json:"users"`
}
