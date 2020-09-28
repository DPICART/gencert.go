package cert

import (
	"encoding/csv"
	"os"
)

func ParseCSV(filename string) ([]*Cert, error) {
	certs := make([]*Cert, 0)
	f, err := os.Open(filename)

	if err != nil {
		return certs, err
	}
	defer f.Close()

	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if nil != err {
		return certs, err
	}

	for _, rec := range records {
		course := rec[0]
		name := rec[1]
		date := rec[2]
		c, err := New(course, name, date)
		if nil != err {
			return certs, err
		}
		certs = append(certs, c)

	}

	return certs, nil

}
