package measurement

import (
    "encoding/json"
    "fmt"
)

type MeasurementStatus struct {
    ParseError error

    data struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
    }
}

func (m *MeasurementStatus) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &m.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }

    return nil
}

// Numeric ID of this status.
func (m *MeasurementStatus) Id() int {
    return m.data.Id
}

// Human-readable description of this status.
func (m *MeasurementStatus) Name() string {
    return m.data.Name
}
