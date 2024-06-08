package date

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type ContactDate time.Time

const DefaultDateFormat = "2006-01-02 15:04:05"

func (j *ContactDate) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(DefaultDateFormat, s)
	if err != nil {
		return err
	}
	*j = ContactDate(t)
	return nil
}

func (j ContactDate) MarshalJSON() ([]byte, error) {
	dateStr := time.Time(j).Format(DefaultDateFormat)

	return json.Marshal(dateStr)
}

func (j ContactDate) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}

func (t *ContactDate) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return t.UnmarshalText(string(v))
	case string:
		return t.UnmarshalText(v)
	case time.Time:
		*t = ContactDate(v)
	case nil:
		*t = ContactDate{}
	default:
		return fmt.Errorf("cannot sql.Scan() ContactDate from: %#v", v)
	}
	return nil
}

func (t ContactDate) Value() (driver.Value, error) {
	return driver.Value(time.Time(t).Format(DefaultDateFormat)), nil
}

func (t *ContactDate) UnmarshalText(value string) error {
	dd, err := time.Parse(DefaultDateFormat, value)
	if err != nil {
		return err
	}
	*t = ContactDate(dd)
	return nil
}

func (ContactDate) GormDataType() string {
	return "TIME"
}
