package http

import (
    "encoding/json"
    "fmt"
)

// HTTP timing result.
type Readtiming struct {
    data struct {
        T float64 `json:"t"`
        O int     `json:"o,string"` // TODO: Not documented as string
    }
}

func (r *Readtiming) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }
    return nil
}

// Time since starting to connect when data is received (in milli seconds).
func (r *Readtiming) T() float64 {
    return r.data.T
}

// Offset in stream of reply data.
func (r *Readtiming) O() int {
    return r.data.O
}
