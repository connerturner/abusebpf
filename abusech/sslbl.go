package abusech

import (
	"encoding/csv"
	"io"
	"net/http"
	"strconv"
)

type SslEntry struct {
	FirstSeen string
	IpAddress string
	Port      uint16
}

func unpackEntry(csvRecord []string) SslEntry {
	port, err := strconv.ParseUint(csvRecord[2], 0xA, 0x10)
	if err != nil {
		return SslEntry{}
	}
	return SslEntry{
		FirstSeen: csvRecord[0],
		IpAddress: csvRecord[1],
		Port:      uint16(port),
	}
}

func unrollCsvData(reader io.Reader) ([]SslEntry, error) {
	// Create new reader from the net/http response
	rd := csv.NewReader(reader)
	// Set comment character to a # rune
	rd.Comment = '#'
	// Ensure each record set is the same length of fields
	rd.FieldsPerRecord = 0
	// Each record is being unpacked into a struct, so we can reuse the slice
	rd.ReuseRecord = true

	var s []SslEntry

	for {
		rec, err := rd.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		s = append(s, unpackEntry(rec))

	}

	return s, nil

}

func RetrieveSslBlEntries() ([]SslEntry, error) {
	req, err := http.NewRequest(http.MethodGet, sslblUrl, nil)
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

	return unrollCsvData(res.Body)

}
