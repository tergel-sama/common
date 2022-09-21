package types

import (
	"database/sql"
	"encoding/json"
)

type NullBool struct {
	sql.NullBool
}

func (x *NullBool) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.Bool)
	} else {
		return json.Marshal(nil)
	}
}

func (x *NullBool) UnMarshalJSON(data []byte) error {

	var v *bool
	err := json.Unmarshal(data, &v)
	if err != nil {
		return err
	}
	if v != nil {
		x.Valid = true
		x.Bool = *v
	} else {
		x.Valid = false
		x.Bool = false
	}
	return nil
}
