package traceroute

import (
    "encoding/json"
    "fmt"
)

// Traceroute reply.
type Reply struct {
    data struct {
        X          string          `json:"x"`
        Err        interface{}     `json:"err"`
        From       string          `json:"from"`
        Ittl       int             `json:"ittl"`
        Edst       string          `json:"edst"`
        Late       int             `json:"late"`
        Mtu        int             `json:"mtu"`
        Rtt        float64         `json:"rtt"`
        Size       int             `json:"size"`
        Ttl        int             `json:"ttl"`
        Flags      string          `json:"flags"`
        Dstoptsize int             `json:"dstoptsize"`
        Hbhoptsize int             `json:"hbhoptsize"`
        Icmpext    json.RawMessage `json:"icmpext"`
    }

    icmpext *Icmpext
}

func (r *Reply) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.data); err != nil {
        return fmt.Errorf("%s for %s", err.Error(), string(b))
    }
    switch v := r.data.Err.(type) {
    case string:
        r.data.Err = v
    case int:
        r.data.Err = fmt.Sprintf("%v", v)
    case float64:
        r.data.Err = fmt.Sprintf("%d", v)
    case nil:
    default:
        return fmt.Errorf("err type %T unexpected for %s", r.data.Err, string(b))
    }

    if r.data.Icmpext != nil {
        r.icmpext = &Icmpext{}
        if err := json.Unmarshal(r.data.Icmpext, r.icmpext); err != nil {
            return fmt.Errorf("Unable to process ICMP extensions: %s", err.Error())
        }
    }

    return nil
}

// On timeout: "*".
func (r *Reply) X() string {
    return r.data.X
}

// Error ICMP: "N" (network unreachable,), "H" (destination unreachable),
// "A" (administratively prohibited), "P" (protocol unreachable),
// "p" (port unreachable) "h" (beyond scope, from fw 4650) (optional).
func (r *Reply) Err() string {
    return r.data.Err.(string)
}

// IPv4 or IPv6 source address in reply.
func (r *Reply) From() string {
    return r.data.From
}

// Time-to-live in packet that triggered the error ICMP. Omitted if equal
// to 1 (optional).
func (r *Reply) Ittl() int {
    return r.data.Ittl
}

// Destination address in the packet that triggered the error ICMP
// if different from the target of the measurement (optional).
func (r *Reply) Edst() string {
    return r.data.Edst
}

// Number of packets a reply is late, in this case rtt is not
// present (optional).
func (r *Reply) Late() int {
    return r.data.Late
}

// Path MTU from a packet too big ICMP (optional).
func (r *Reply) Mtu() int {
    return r.data.Mtu
}

// Round-trip-time of reply, not present when the response is late.
func (r *Reply) Rtt() float64 {
    return r.data.Rtt
}

// Size of reply.
func (r *Reply) Size() int {
    return r.data.Size
}

// Time-to-live in reply.
func (r *Reply) Ttl() int {
    return r.data.Ttl
}

// TCP flags in the reply packet, for TCP traceroute, concatenated, in
// the order 'F' (FIN), 'S' (SYN), 'R' (RST), 'P' (PSH), 'A' (ACK),
// 'U' (URG) (optional).
func (r *Reply) Flags() string {
    return r.data.Flags
}

// Size of destination options header (IPv6) (optional).
func (r *Reply) Dstoptsize() int {
    return r.data.Dstoptsize
}

// Size of hop-by-hop options header (IPv6) (optional).
func (r *Reply) Hbhoptsize() int {
    return r.data.Hbhoptsize
}

// Information when icmp header is found in reply (optional).
func (r *Reply) Icmpext() *Icmpext {
    return r.icmpext
}
