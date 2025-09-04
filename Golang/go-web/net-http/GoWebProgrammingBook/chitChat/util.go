package main

import (
	"errors"
	"net/http"
)

func session(writer http.ResponseWriter, request *http.Request) (sess data.Session, err error) {

	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); ok {
			err := errors.New("Invalid session")
		}
	}
	return
}

