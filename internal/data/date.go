package data

import (
	"encoding/json"
	"strconv"
	"time"
)

type Date struct{ time.Time }

func (d Date) MarshalJSON() ([]byte, error) {
	if d.IsZero() {
		return []byte(strconv.Quote("")), nil
	}

	return json.Marshal(d.Time)
}
