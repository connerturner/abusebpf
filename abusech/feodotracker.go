package abusech

import (
	"encoding/json"
	"io"
	"net/http"
)

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

func unmarshallFeodoJson(jsonBytes []byte) ([]Entry, error) {
	var entries []Entry
	err := json.Unmarshal(jsonBytes, &entries)

	if err != nil {
		return nil, err
	}

	return entries, nil
}

func RetrieveFeodoEntries() ([]Entry, error) {
	req, err := http.NewRequest(http.MethodGet, feodoUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, err
	}

	responseData, err := io.ReadAll(res.Body)
	entries, err := unmarshallFeodoJson(responseData)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
