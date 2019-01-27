package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

/*
  List users in UAA
*/
func ListUsers(w http.ResponseWriter, r *http.Request) {
	apiQuery := "/Users"
	uaaRespBytes, flash, userNameVal := ClientRequest(w, r, apiQuery)
	uaaResp := UsersInfo{}
	if uaaServErr := json.Unmarshal([]byte(uaaRespBytes), &uaaResp); uaaServErr != nil {
		fmt.Println(uaaServErr)
	}
	uaaRespdata := uaaResp.Users
	data := ListPage{
		PageTitle: "Users",
		UserName:  userNameVal,
		PageData:  uaaRespdata,
		Flash:     flash,
	}
	tmpl := template.Must(template.ParseFiles("templates/list/users.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base", data)
	return
}
