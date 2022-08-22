package abusech

type Entry struct {
	IpAddress string `json:"ip_address"`
	Port      uint16
	Status    string
	Hostname  string
	AsNum     int    `json:"as_number"`
	AsName    string `json:"as_name"`
	Country   string
	FirstSeen string `json:"first_seen"`
	LastSeen  string `json:"last_seen"`
	Malware   string
}
