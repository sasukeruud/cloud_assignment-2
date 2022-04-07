package structs

type Status struct {
	CovidCasesApi  int
	CovidPolicyApi int
	Webhooks       int
	Version        string
	Uptime         float64
}

type Policy struct {
	PolicyActions []struct {
		PolicyTypeCode    string      `json:"policy_type_code"`
		PolicyTypeDisplay string      `json:"policy_type_display"`
		Policyvalue       string      `json:"policyvalue"`
		IsGeneral         bool        `json:"is_general"`
		Notes             interface{} `json:"notes"`
	} `json:"policyActions"`
	StringencyData struct {
		DateValue        string  `json:"date_value"`
		CountryCode      string  `json:"country_code"`
		Confirmed        int     `json:"confirmed"`
		StringencyActual float64 `json:"stringency_actual"`
		Stringency       float64 `json:"stringency"`
	} `json:"stringencyData"`
}

type Webhooks struct {
	WebhookID string `json:"webhookID,omitempty"`
	Url       string `json:"URL,omitempty"`
	Country   string `json:"country,omitempty"`
	Calls     int    `json:"calls,omitempty"`
}

type Country_calls struct {
	Country_id string `json:"country_calls"`
	Country    string `json:"country"`
	Called     int    `json:"called"`
}
