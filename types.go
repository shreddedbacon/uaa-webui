package main

type UsersInfo struct {
	Users        []UserInfo `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type UserInfo struct {
	GUID                 string `json:"id"`
	ExternalID           string
	Username             string
	Name                 Name
	Groups               []Group
	Emails               []UserEmail
	Active               bool
	Verified             bool
	Origin               string
	ZoneID               string
	PasswordLastModified string
	Schemas              []string
}

type Name struct {
	GivenName  string
	FamilyName string
}

type UserEmail struct {
	Value   string
	Primary bool
}

type Group struct {
	Value   string
	Display string
	Type    string
}

type ExternalGroups struct {
	Groups       []ExternalGroup `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type ExternalGroup struct {
	DisplayName   string `json:"displayName"`
	ExternalGroup string `json:"externalGroup"`
	GroupID       string `json:"groupId"`
	Origin        string `json:"origin"`
}

type OauthClients struct {
	Clients      []OauthClient `json:"resources"`
	StartIndex   int
	ItemsPerPage int
	TotalResults int
	Schemas      []string
}

type OauthClient struct {
	ID   string `json:"client_id"`
	Name string
	//AutoApprove          bool - TODO - this field can be a bool or an array??
	Action                 string
	Scope                  []string
	ResourceIDs            []string `json:"resource_ids"`
	Authorities            []string
	AuthorizedGrantTypes   []string `json:"authorized_grant_types"`
	LastModified           int
	RedirectURI            []string `json:"redirect_uri"`
	SignupRedirectURL      string   `json:"signup_redirect_url"`
	ChangeEmailRedirectURL string   `json:"change_email_redirect_url"`
}

type IdentityZone struct {
	ID           string
	Subdomain    string
	Name         string
	Version      int
	Description  string
	Created      int
	LastModified int `json:"last_modified"`
}

type IdentityProviders struct {
	ID           string
	OriginKey    string
	Name         string
	Version      int
	Config       string
	Created      int
	LastModified int `json:"last_modified"`
}

type ListPage struct {
	PageTitle string
	UserName  string
	PageData  interface{}
	Flash     Flash
}
