package ripeatlas

type ProbeStatus struct {
	Since string `json:"status"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

type Tags struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

type Geometry struct {
	Gtype       string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

type Probe struct {
	Id              int    `json:"id"`
	Description     string `json:"description"`
	ProbeStatus     `json:"status"`
	Prefix_v4       string `json:"prefix_v4"`
	Prefix_v6       string `json:"prefix_v6"`
	Last_connected  int    `json:"last_connected"`
	Ptype           string `json:"type"`
	Address_v6      string `json:"address_v6"`
	Address_v4      string `json:"address_v4"`
	Total_uptime    int    `json:"total_uptime"`
	Country_code    string `json:"country_code"`
	Is_public       bool   `json:"is_public"`
	Asn_v4          int    `json:"asn_v4"`
	Asn_v6          int    `json:"asn_v6"`
	Status_since    int    `json:"status_since"`
	First_connected int    `json:"first_connected"`
	Is_anchor       bool   `json:"is_anchor"`
	Geometry        `json:"geometry"`
	Tags            []Tags `json:"tags"`
}
