package main

type Todo struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Completed int `json:"completed"`
}