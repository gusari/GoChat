package main

import ()
import (
	"github.com/gusari/practice/chitChat/data"
	"github.com/pkg/errors"
	"net/http"
)

//check if the user is logged in and has a session, if not err is not nil
func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	cookie, err := r.Cookie("_cookie") //ログインしていなければないよね
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
