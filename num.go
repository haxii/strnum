package strnum

import (
	"encoding/json"
	"strconv"
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

func NewStr(s string) StrNum {
	return StrNum{json.Number(s)}
}

func NewInt64(i int64) StrNum {
	return StrNum{json.Number(strconv.FormatInt(i, 10))}
}

func NewInt(i int) StrNum {
	return StrNum{json.Number(strconv.Itoa(i))}
}
