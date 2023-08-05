package entities

type Message struct {
	AlertName   string
	Severity    string
	Instance    string `json:"instance,omitempty"`
	StartedAt   string
	EndedAt     string `json:"endedat,omitempty"`
	Summary     string `json:"summary,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string
}

type BankMessage struct {
	PhoneNumber string `json:"PhoneNumber"`
	Message     string `json:"Message"`
}
