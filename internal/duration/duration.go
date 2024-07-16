package duration

import (
	"errors"
	"time"

	go_json "github.com/goccy/go-json"
)

// A Duration is a wrapper around time.Duration that provides JSON marshalling and unmarshalling.
type Duration time.Duration

func (d Duration) MarshalJSON() ([]byte, error) {
	return go_json.Marshal(d.Value().String())
}

func (d *Duration) UnmarshalJSON(b []byte) error {
	var v interface{}
	if err := go_json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		*d = Duration(time.Duration(value))
		return nil
	case string:
		tmp, err := time.ParseDuration(value)
		if err != nil {
			return err
		}
		*d = Duration(tmp)
		return nil
	default:
		return errors.New("invalid duration")
	}
}

func (d Duration) Value() time.Duration {
	return time.Duration(d)
}
