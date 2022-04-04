package structs

type Status struct {
	CovidCasesApi  int
	CovidPolicyApi int
	Webhooks       string
	Version        string
	Uptime         float64
}

type Cases struct {
	Data struct {
		Country struct {
			Name       string `json:"name"`
			MostRecent struct {
				Date      string `json:"date"`
				Confirmed int    `json:"confirmed"`
			} `json:"mostRecent"`
		} `json:"country"`
	} `json:"data"`
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
	WebhookID string `json:"webhookID"`
	Url       string `json:"URL"`
	Country   string `json:"country"`
	Calls     int    `json:"calls"`
}
