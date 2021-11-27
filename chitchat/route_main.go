package main

import (
	"net/http"
	"github.com/1021ky/read_and_write_introduction_to_practice_go_programming/chitchat/data"
)

// HTTPハンドラ関数
func index(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"templates/layout.html",
		"templates/navbar.html",
		"templates/index.html",
	}
	templates := template.Must(template.ParseFiles(files...))
	threads,  err := data.Threads(); if err == nil{
		templates.ExecuteTemplate(w, "layout", threads)
	}
}
