package types

import (
	"database/sql"
	"encoding/json"
)

// Implements the Marshaller and Unmarshaller interfaces.
// See more on: https://pkg.go.dev/encoding/json@go1.19
// And composed of (embeds) the sql.Null* types to implement the Scanner interface.
type NullInt32 struct {
	sql.NullInt32
}

// Executed from outer json.Marshal or encoder.Encode methods.
// Marshal traverses the value v recursively.
// If an encountered value implements the Marshaler interface and is not a nil pointer,
// Marshal calls its MarshalJSON method to produce JSON.
func (x *NullInt32) MarshalJSON() ([]byte, error) {
	if x.Valid {
		return json.Marshal(x.Int32)
	} else {
		return json.Marshal(nil)
	}
}

// Executed from outer json.Unmarshal or decoder.Decode methods.
// To unmarshal JSON into a value implementing the Unmarshaler interface,
// Unmarshal calls that value's UnmarshalJSON method, including when the input is a JSON null.
// Otherwise, if the value implements encoding.TextUnmarshaler and the input is a JSON quoted string,
// Unmarshal calls that value's UnmarshalText method with the unquoted form of the string.
func (x *NullInt32) UnmarshalJSON(data []byte) error {
	var v *int32
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	if v != nil {
		x.Valid = true
		x.Int32 = *v
	} else {
		x.Valid = false
		x.Int32 = 0
	}
	return nil
}
