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
	Country_code string
	Scope        string
	Stringency   string
	Policies     int
}
