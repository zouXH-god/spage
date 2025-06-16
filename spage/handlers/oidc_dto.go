package handlers

type OidcDto struct {
	Name             string   `json:"name"`
	DisplayName      string   `json:"display_name"`
	ClientID         string   `json:"client_id"`
	ClientSecret     string   `json:"client_secret"`
	OidcDiscoveryUrl string   `json:"oidc_discovery_url"`
	Icon             string   `json:"icon"`
	GroupClaims      string   `json:"group_claims"`
	AdminGroups      []string `json:"admin_groups"`
	AllowedGroups    []string `json:"allowed_groups"`
}
