package dns

import (
    "encoding/json"
    "fmt"
)

// Error message.
type Error struct {
    data struct {
        Timeout     int    `json:"timeout"`
        Getaddrinfo string `json:"getaddrinfo"`
    }
}

func (e *Error) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &e.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }
    return nil
}

// Query timeout.
func (e *Error) Timeout() int {
    return e.data.Timeout
}

// Error message.
func (e *Error) Getaddrinfo() string {
    return e.data.Getaddrinfo
}
