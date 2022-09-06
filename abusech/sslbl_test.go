package abusech

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_unpackEntryValid(t *testing.T) {
	exp := SslEntry{
		FirstSeen: "2022-01-01 00:00:01",
		IpAddress: "198.51.100.254",
		Port:      8080,
	}
	res := unpackEntry([]string{
		"2022-01-01 00:00:01", "198.51.100.254", "8080",
	})

	if dis := cmp.Diff(exp, res); dis != "" {
		t.Errorf(dis)
	}
}

func Test_unpackEntryOverflow(t *testing.T) {
	res := unpackEntry([]string{
		"2022-01-01 00:00:01", "198.51.100.254", "65536",
	})

	if res.Port != 0 {
		t.Errorf("Expecting uint16 conversion error (0) but got %d", res.Port)
	}
}
