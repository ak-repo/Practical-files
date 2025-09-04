package main

import "net/http"




func authenticate(writer http.ResponseWriter, request *http.Request){

	request.ParseForm()
	user,_ := data.UserByEmail(request.PostFormValue("email"))
	if user.Password:= data.Encrypt(request.PostFormValue("passwprd")){
		session:= user.CreateSession()
		cookie:= http.Cookie{
			Name: "_cookie",
			Value: session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)


	}else{
		http.Redirect(writer,request,"/login",302)
	}
}