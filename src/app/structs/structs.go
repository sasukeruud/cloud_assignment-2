package structs

import "time"

type status struct {
	covidCasesApi string
	covidPolicy   string
	webhooks      []string
	version       string
	uptime        time.Time
}
