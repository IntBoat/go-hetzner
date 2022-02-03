package hetzner

import (
	"bytes"
	"fmt"
	"time"
)

// JSONTime For ordering API, MarketProduct.NextReduceDate
type JSONTime struct {
	time.Time
}

func (t *JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02 15:04:05"))), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, []byte(`"null"`)) {
		return nil
	}

	t.Time, err = time.Parse(`"2006-01-02 15:04:05"`, string(data))
	return
}

// JSONDate For server API, server Cancellation.EarliestCancellationDate and Cancellation.CancellationDate
type JSONDate struct {
	time.Time
}

func (t *JSONDate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, t.Format("2006-01-02"))), nil
}

func (t *JSONDate) UnmarshalJSON(data []byte) (err error) {
	if bytes.Equal(data, []byte(`"null"`)) {
		return nil
	}

	t.Time, err = time.Parse(`"2006-01-02"`, string(data))
	return
}
