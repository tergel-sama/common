package types

import (
	"database/sql"
	"encoding/json"
)

type NullFloat64 struct {
	sql.NullFloat64
}

func (x *NullFloat64) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.Float64)
	} else {
		return json.Marshal(nil)
	}
}

func (x *NullFloat64) UnmarshalJSON(data []byte) error {
	var v *float64
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v != nil {
		x.Valid = true
		x.Float64 = *v
	} else {
		x.Valid = false
		x.Float64 = 0
	}
	return nil
}
