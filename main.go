package main

type book struct {
	ID       string `json: "id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json: "quantity"`
}


var books = []book{
	{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 2},
	{}
}