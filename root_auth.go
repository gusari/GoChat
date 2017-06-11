package main

import (
	"fmt"
	"github.com/gusari/practice/chitChat/data"
	"net/http"
)

//POST /authenticae
//Authenticate the user given the email & password
//ユーザーを認証してクッキーを返すハンドラ
func authenfication(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm() //urlが渡すオプションを解析。返り値Error型。POSTに関してはレスポンスパケットのボディを解析する
	//もしParseFormがコールされなければ以下でフォームのデータ取得はできない
	fmt.Println(r.Form)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.Scheme)

	//ParseForm populates r.Form and r.PostForm.
	user, err := data.UserByEmail(r.PostFormValue("email"))
	if err != nil {
		danger(err, "cannnot find user")
	}

	if user.Password == data.Encrypt(r.PostFormValue("password")) {
		session, err := user.CreateSession()
		if err != nil {
			danger(err, "cannnot create session")
		}

		cookie := http.Cookie{
			Name:     "_cookies",
			Value:    session.Uuid, //ユーザのユニークid
			HttpOnly: true,         //非HTTPのAPIアクセス禁止
		}
		http.SetCookie(w, &cookie)
		http.Redirect(w, r, "/", 302) //300台はリダイレクト
	} else {
		http.Redirect(w, r, "/login", 302)
	}
}
