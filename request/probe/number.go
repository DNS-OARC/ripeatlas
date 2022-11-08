package probe

import (
    "encoding/json"
    "fmt"
)

type Number struct {
    ParseError error

    data struct {
        Type        string    `json:"type"`
        Coordinates []float64 `json:"coordinates"`
    }
}

func (n *Number) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &n.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }

    return nil
}

// .
func (n *Number) Type() string {
    return n.data.Type
}

// .
func (n *Number) Coordinates() []float64 {
    return n.data.Coordinates
}
