package cert

import (
	"testing"
)

func TestValidCertData(t *testing.T) {
	c, err := New("Golang", "Bob", "2020-09-28")
	if nil != err {
		t.Errorf("Cert data should be valid. err=%v", err)
	}
	if c == nil {
		t.Errorf("Cert should be a valid reference. got=nil")
	}

	if c.Course != "GOLANG COURSE" {
		t.Errorf("Course name is not valid. Expected=GOLANG COURSE got=%v", c.Course)
	}
}

func TestCourseEmptyValue(t *testing.T) {
	_, err := New("", "Bob", "2020-09-28")
	if err == nil {
		t.Error("Error should be returned on an empty course")
	}
}

func TestCourseTooLong(t *testing.T) {
	_, err := New("azertyuiopqsdfghjklmwxcvbnazertyuiopqsdfghjklmwxcvbn", "Bob", "2020-09-28")
	if err == nil {
		t.Error("Error should be returned with a long course name")
	}
}

func TestNameTooLong(t *testing.T) {
	_, err := New("Golang", "azertyuiopqsdfghjazertyuiopqsdfghj", "2020-09-28")
	if err == nil {
		t.Error("Error should be returned with a long name")
	}
}

func TestNameEmptyValue(t *testing.T) {
	_, err := New("Golang", "", "2020-09-28")
	if err == nil {
		t.Error("Error should be returned on an empty name")
	}
}
