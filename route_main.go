package main

import (
	"github.com/gusari/practice/chitChat/data"
	"net/http"
)

func err(w http.ResponseWriter, r *http.Request){
	vals := r.URL.Query()
	_, err := session(w, r)//ログインしているかをチェック
	if err != nil {
		generateHTML()
	} else {

	}
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(w, r, "cannnot get threads")
	} else {
		_, err := session(w, r)
		if err != nil {

		} else {

		}
	}
}
