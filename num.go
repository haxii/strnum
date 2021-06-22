package strnum

import (
	"encoding/json"
)

type StrNum struct {
	json.Number
}

func (s *StrNum) UnmarshalJSON(b []byte) error {
	if s == nil {
		return nil
	}
	return json.Unmarshal(b, &s.Number)
}

func (s StrNum) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s.Number))
}

func (s StrNum) Val() int64 {
	v, _ := s.Number.Int64()
	return v
}
