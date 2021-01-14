package models

// Plugin :: This struct will be used for queries in mongodb
type Plugin struct {
	ID struct {
		Oid string `json:"$oid,omitempty"`
	} `json:"_id"`
	PluginName      string `json:"PluginName,omitempty"`
	Vulnerabilities []struct {
		Title      string   `json:"Title,omitempty"`
		Published  string   `json:"Published,omitempty"`
		References []string `json:"References,omitempty"`
		Version    string   `json:"Version,omitempty"`
	} `json:"Vulnerabilities,omitempty"`
	Type string `json:"Type,omitempty"`
}
