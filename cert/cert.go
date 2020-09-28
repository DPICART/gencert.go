package cert

import (
	"fmt"
	"strings"
	"time"
)

var MaxLenCourse = 30
var maxLenName = 30

type Cert struct {
	Course             string
	Name               string
	Date               time.Time
	LabelTitle         string
	LabelCompletion    string
	LabelPresented     string
	LabelParticipation string
	LabelDate          string
}

type Saver interface {
	Save(c Cert) error
}

func New(course, name, date string) (*Cert, error) {

	c, err := validateCourse(course)
	if nil != err {
		return nil, err
	}
	n, err := validateName(name)
	if nil != err {
		return nil, err
	}

	d, err := parseDate(date)
	if nil != err {
		return nil, err
	}

	cert := &Cert{
		Course:             c,
		Name:               n,
		Date:               d,
		LabelTitle:         fmt.Sprintf("%v Certificate - %v", c, n),
		LabelCompletion:    "Certificate of completion",
		LabelPresented:     "This certificate is presented to",
		LabelParticipation: fmt.Sprintf("For participation in the %v", c),
		LabelDate:          fmt.Sprintf("Date: %v", d.Format("02/01/2006")),
	}
	return cert, nil
}

func parseDate(date string) (time.Time, error) {
	d, err := time.Parse("2006-01-02", date)
	if err != nil {
		return d, err
	}
	return d, nil
}

func validateCourse(course string) (string, error) {
	c, err := validateStr(course, MaxLenCourse)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(c, " course") {
		c = c + " course"
	}
	return strings.ToTitle(c), nil
}

func validateName(name string) (string, error) {
	c, err := validateStr(name, maxLenName)
	if nil != err {
		return "", err
	}
	return strings.ToTitle(c), nil
}

func validateStr(str string, maxLen int) (string, error) {
	c := strings.TrimSpace(str)
	if len(c) <= 0 {
		return c, fmt.Errorf("Invalid string got=%v. len=%d", c, len(c))
	}

	if len(c) > maxLen {
		return c, fmt.Errorf("Invalid string got=%v. len=%d", c, len(c))
	}
	return c, nil
}
