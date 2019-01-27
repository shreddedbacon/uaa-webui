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
func ListExternalGroups(w http.ResponseWriter, r *http.Request) {
	apiQuery := "/Groups/External"
	uaaRespBytes, flash, userNameVal := ClientRequest(w, r, apiQuery)
	uaaResp := ExternalGroups{}
	if uaaServErr := json.Unmarshal([]byte(uaaRespBytes), &uaaResp); uaaServErr != nil {
		fmt.Println(uaaServErr)
	}
	uaaRespdata := uaaResp.Groups
	data := ListPage{
		PageTitle: "External Groups",
		UserName:  userNameVal,
		PageData:  uaaRespdata,
		Flash:     flash,
	}
	tmpl := template.Must(template.ParseFiles("templates/list/groups-external.html", "templates/base.html"))
	tmpl.ExecuteTemplate(w, "base", data)
	return
}
