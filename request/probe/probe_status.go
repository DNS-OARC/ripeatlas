package probe

import (
    "encoding/json"
    "fmt"
)

type ProbeStatus struct {
    ParseError error

    data struct {
        Status string `json:"status"`
        Id     int    `json:"id"`
        Name   string `json:"name"`
    }
}

func (p *ProbeStatus) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &p.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }

    return nil
}

// .
func (p *ProbeStatus) Status() string {
    return p.data.Status
}

// .
func (p *ProbeStatus) Id() int {
    return p.data.Id
}

// .
func (p *ProbeStatus) Name() string {
    return p.data.Name
}
