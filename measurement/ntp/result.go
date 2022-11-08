package ntp

import (
    "encoding/json"
    "fmt"
)

// NTP result.
type Result struct {
    data struct {
        FinalTs    float64 `json:"final-ts"`
        Offset     float64 `json:"offset"`
        OriginTs   float64 `json:"origin-ts"`
        ReceiveTs  float64 `json:"receive-ts"`
        Rtt        float64 `json:"rtt"`
        TransmitTs float64 `json:"transmit-ts"`
    }
}

func (r *Result) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }
    return nil
}

// NTP time the response from the server is received.
func (r *Result) FinalTs() float64 {
    return r.data.FinalTs
}

// Clock offset between client and server in seconds.
func (r *Result) Offset() float64 {
    return r.data.Offset
}

// NTP time the request was sent.
func (r *Result) OriginTs() float64 {
    return r.data.OriginTs
}

// NTP time the server received the request.
func (r *Result) ReceiveTs() float64 {
    return r.data.ReceiveTs
}

// Round trip time between client and server in seconds.
func (r *Result) Rtt() float64 {
    return r.data.Rtt
}

// NTP time the server sent the response.
func (r *Result) TransmitTs() float64 {
    return r.data.TransmitTs
}
