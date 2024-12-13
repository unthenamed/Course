package app

import (
	"encoding/json"
	"strings"
	"time"
)

// First create a type alias
type JsonBirthDate time.Time

// Implement Marshaler and Unmarshaler interface
func (j *JsonBirthDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	t, err := time.Parse("02-01-2006", s)
	if err != nil {
		return err
	}
	*j = JsonBirthDate(t)
	return nil
}

func (j JsonBirthDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

// Maybe a Format function for printing your date
func (j JsonBirthDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
