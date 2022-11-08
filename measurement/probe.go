package measurement

import (
    "encoding/json"
    "fmt"
)

type Probe struct {
    ParseError error

    data struct {
        Id  int    `json:"id"`
        Url string `json:"url"`
    }
}

func (p *Probe) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &p.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }

    return nil
}

// ID of this probe.
func (p *Probe) Id() int {
    return p.data.Id
}

// The URL that contains the details of this probe.
func (p *Probe) Url() string {
    return p.data.Url
}
