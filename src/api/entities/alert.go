package entities

import "time"

type KV map[string]string

type Alert struct {
	Status       string
	Labels       KV
	Annotations  KV
	StartsAt     time.Time
	EndsAt       time.Time
	GeneratorURL string
	Fingerprint  string
}

type Data struct {
	Receiver          string
	Status            string
	Alerts            []Alert
	GroupLabels       KV
	CommonLabels      KV
	CommonAnnotations KV
	ExternalURL       string
}
