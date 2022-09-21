package types

import (
	"database/sql"
	"encoding/json"
)

type NullInt16 struct {
	sql.NullInt16
}

func (x *NullInt16) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.Int16)
	} else {
		return json.Marshal(nil)
	}
}

func (x *NullInt16) UnmarshalJSON(data []byte) error {
	var v *int16
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v != nil {
		x.Valid = true
		x.Int16 = *v
	} else {
		x.Valid = false
		x.Int16 = 0
	}
	return nil
}
