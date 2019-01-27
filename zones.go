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
func ListZones(w http.ResponseWriter, r *http.Request) {
	apiQuery := "/identity-zones"
	uaaRespBytes, flash, userNameVal := ClientRequest(w, r, apiQuery)
	uaaResp := []IdentityZone{}
	if uaaServErr := json.Unmarshal([]byte(uaaRespBytes), &uaaResp); uaaServErr != nil {
		fmt.Println(uaaServErr)
	}
	uaaRespdata := uaaResp
	data := ListPage{
		PageTitle: "Identity Zones",
		UserName:  userNameVal,
		PageData:  uaaRespdata,
		Flash:     flash,
	}
	tmpl := template.Must(template.ParseFiles("templates/list/zones.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base", data)
	return
}
