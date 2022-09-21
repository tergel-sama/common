package types

import (
	"database/sql"
	"encoding/json"
)

type NullString struct {
	sql.NullString
}

func (x *NullString) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.String)
	} else {
		return json.Marshal(nil)
	}
}

func (x *NullString) UnmarshalJSON(data []byte) error {
	var v *string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v != nil {
		x.Valid = true
		x.String = *v
	} else {
		x.Valid = false
		x.String = ""
	}
	return nil
}
