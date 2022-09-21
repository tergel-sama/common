package types

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullTime struct {
	sql.NullTime
}

func (x *NullTime) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (x *NullTime) UnmarshalJSON(data []byte) error {
	var v *time.Time
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v != nil {
		x.Valid = true
		x.Time = *v
	} else {
		x.Valid = false
		x.Time = time.Now()
	}
	return nil
}
