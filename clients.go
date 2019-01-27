package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

/*
  List zones in UAA
*/
func ListOAuthClients(w http.ResponseWriter, r *http.Request) {
	apiQuery := "/oauth/clients"
	uaaRespBytes, flash, userNameVal := ClientRequest(w, r, apiQuery)
	uaaResp := OauthClients{}
	if uaaServErr := json.Unmarshal([]byte(uaaRespBytes), &uaaResp); uaaServErr != nil {
		fmt.Println(uaaServErr)
	}
	uaaRespdata := uaaResp.Clients
	data := ListPage{
		PageTitle: "OAuth Clients",
		UserName:  userNameVal,
		PageData:  uaaRespdata,
		Flash:     flash,
	}
	tmpl := template.Must(template.ParseFiles("templates/list/clients.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base", data)
	return
}
