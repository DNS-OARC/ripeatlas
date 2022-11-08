package traceroute

import (
    "encoding/json"
    "fmt"
)

// Traceroute result.
type Result struct {
    data struct {
        Hop    int             `json:"hop"`
        Error  string          `json:"error"`
        Result json.RawMessage `json:"result"`
    }

    replies []*Reply
}

func (r *Result) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }

    if r.data.Result != nil {
        if err := json.Unmarshal(r.data.Result, &r.replies); err != nil {
            return fmt.Errorf("Unable to process Replies: %s", err.Error())
        }
    }

    return nil
}

// Hop number.
func (r *Result) Hop() int {
    return r.data.Hop
}

// When an error occurs trying to send a packet. In that case there will
// not be a result structure (optional).
func (r *Result) Error() string {
    return r.data.Error
}

// Traceroute replies (called "result" in RIPE Atlas API documentation).
func (r *Result) Replies() []*Reply {
    return r.replies
}
