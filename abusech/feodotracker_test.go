package abusech

import (
	"fmt"
	"testing"
)

func TestUnmarhsall(t *testing.T) {
	testData := []byte(`[
		{
			"ip_address": "51.178.161.32",
			"port": 4643,
			"status": "online",
			"hostname": "srv-web.ffconsulting.com",
			"as_number": 16276,
			"as_name": "OVH",
			"country": "FR",
			"first_seen": "2021-01-17 07:44:46",
			"last_online": "2022-08-21",
			"malware": "Dridex"
		},
		{
			"ip_address": "66.175.217.172",
			"port": 13786,
			"status": "online",
			"hostname": "li512-172.members.linode.com",
			"as_number": 63949,
			"as_name": "LINODE-AP Linode, LLC",
			"country": "US",
			"first_seen": "2021-07-14 15:30:51",
			"last_online": "2022-08-21",
			"malware": "Dridex"
		}
	]`)
	outcome, err := UnmarshallFeodoJson(testData)
	fmt.Printf("%+v\n", outcome)
	if err != nil {
		t.Errorf("Encountered Error: %s", err)
	}

	if len(outcome) != 2 {
		t.Errorf("Given JSON has 2 entries, expecting 2 length but got: %d", len(outcome))
	}

	if outcome[0].IpAddress != "51.178.161.32" {
		t.Errorf("Expecting IP address of 51.178.161.32, but got: %s", outcome[1].IpAddress)
	}

}
