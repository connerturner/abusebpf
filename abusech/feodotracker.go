package abusech

import "encoding/json"

type Entry struct {
	IpAddress string `json:"ip_address"`
	Port      uint16
	Status    string
	Hostname  string
	AsNum     int    `json:"as_number"`
	AsName    string `json:"as_name"`
	Country   string
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_online"`
	Malware   string
}

func UnmarshallFeodoJson(jsonBytes []byte) ([]Entry, error) {
	var entries []Entry
	err := json.Unmarshal(jsonBytes, &entries)

	if err != nil {
		return nil, err
	}

	return entries, nil
}
