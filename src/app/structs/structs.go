package structs

type Status struct {
	CovidCasesApi  int
	CovidPolicyApi int
	Webhooks       string
	Version        string
	Uptime         float64
}

type Cases struct {
	Country     string
	Date        string
	Confirmed   int64
	Recovers    int64
	Deaths      int64
	Growth_rate float64
}

type Policy struct {
	Country_code string
	Scope        string
	Stringency   string
	Policies     int
}
