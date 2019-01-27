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
func ListProviders(w http.ResponseWriter, r *http.Request) {
	apiQuery := "/identity-providers?rawConfig=false"
	uaaRespBytes, flash, userNameVal := ClientRequest(w, r, apiQuery)
	uaaResp := []IdentityProviders{}
	if uaaServErr := json.Unmarshal([]byte(uaaRespBytes), &uaaResp); uaaServErr != nil {
		fmt.Println(uaaServErr)
	}
	uaaRespdata := uaaResp
	data := ListPage{
		PageTitle: "Identity Providers",
		UserName:  userNameVal,
		PageData:  uaaRespdata,
		Flash:     flash,
	}
	tmpl := template.Must(template.ParseFiles("templates/list/providers.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base", data)
	return
}
