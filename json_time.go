package hetzner

import (
	"fmt"
	"time"
)

type JSONTime struct {
	time.Time
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	var err error
	date := string(data)
	if date == `"null"` {
		return nil
	}

	t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, date)
	if err != nil {
		return err
	}

	return nil
}
