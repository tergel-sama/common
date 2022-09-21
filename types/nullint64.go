package types

import (
	"database/sql"
	"encoding/json"
)

type NullInt64 struct {
	sql.NullInt64
}

func (x *NullInt64) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.Int64)
	} else {
		return json.Marshal(nil)
	}
}
func (x *NullInt64) UnmarshalJSON(data []byte) error {
	var v *int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v != nil {
		x.Valid = true
		x.Int64 = *v
	} else {
		x.Valid = false
		x.Int64 = 0
	}
	return nil
}
