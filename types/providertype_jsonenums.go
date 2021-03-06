// generated by jsonenums -type=ProviderType; DO NOT EDIT

package types

import (
	"encoding/json"
	"fmt"
)

var (
	_ProviderTypeNameToValue = map[string]ProviderType{
		"ProviderTypeUnknown":    ProviderTypeUnknown,
		"ProviderTypeKubernetes": ProviderTypeKubernetes,
		"ProviderTypeHelm":       ProviderTypeHelm,
	}

	_ProviderTypeValueToName = map[ProviderType]string{
		ProviderTypeUnknown:    "ProviderTypeUnknown",
		ProviderTypeKubernetes: "ProviderTypeKubernetes",
		ProviderTypeHelm:       "ProviderTypeHelm",
	}
)

func init() {
	var v ProviderType
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_ProviderTypeNameToValue = map[string]ProviderType{
			interface{}(ProviderTypeUnknown).(fmt.Stringer).String():    ProviderTypeUnknown,
			interface{}(ProviderTypeKubernetes).(fmt.Stringer).String(): ProviderTypeKubernetes,
			interface{}(ProviderTypeHelm).(fmt.Stringer).String():       ProviderTypeHelm,
		}
	}
}

// MarshalJSON is generated so ProviderType satisfies json.Marshaler.
func (r ProviderType) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _ProviderTypeValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid ProviderType: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so ProviderType satisfies json.Unmarshaler.
func (r *ProviderType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ProviderType should be a string, got %s", data)
	}
	v, ok := _ProviderTypeNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid ProviderType %q", s)
	}
	*r = v
	return nil
}
