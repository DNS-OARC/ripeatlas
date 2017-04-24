// Author Jerry Lundström <jerry@dns-oarc.net>
// Copyright (c) 2017, OARC, Inc.
// All rights reserved.
//
// This file is part of ripeatlas.
//
// ripeatlas is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ripeatlas is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with ripeatlas.  If not, see <http://www.gnu.org/licenses/>.

package measurement

import (
    "encoding/base64"
    "encoding/json"
    "fmt"

    "github.com/DNS-OARC/ripeatlas/measurement/dns"
    mdns "github.com/miekg/dns"
)

// A measurement result object.
type Result struct {
    data struct {
        Fw         int             `json:"fw"`
        Af         int             `json:"af"`
        DstAddr    string          `json:"dst_addr"`
        DstName    string          `json:"dst_name"`
        Error      json.RawMessage `json:"error"`
        From       string          `json:"from"`
        Lts        int             `json:"lts"`
        MsmId      int             `json:"msm_id"`
        PrbId      int             `json:"prb_id"`
        Proto      string          `json:"proto"`
        Qbuf       string          `json:"qbuf"`
        Result     json.RawMessage `json:"result"`
        Resultsets json.RawMessage `json:"resultset"`
        Retry      int             `json:"retry"`
        Timestamp  int             `json:"timestamp"`
        Type       string          `json:"type"`
    }

    dnsError      *dns.Error
    dnsResult     *dns.Result
    dnsResultsets []*dns.Resultset
}

func (r *Result) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.data); err != nil {
        fmt.Printf("%s\n", string(b))
        return err
    }

    if r.data.Type == "dns" {
        if r.data.Error != nil {
            r.dnsError = &dns.Error{}
            if err := json.Unmarshal(r.data.Error, r.dnsError); err != nil {
                return fmt.Errorf("Unable to process DNS error (fw %d): %s", r.data.Fw, err.Error())
            }
        }
        if r.data.Result != nil {
            r.dnsResult = &dns.Result{}
            if err := json.Unmarshal(r.data.Result, r.dnsResult); err != nil {
                return fmt.Errorf("Unable to process DNS result (fw %d): %s", r.data.Fw, err.Error())
            }
        }
        if r.data.Resultsets != nil {
            if err := json.Unmarshal(r.data.Resultsets, &r.dnsResultsets); err != nil {
                return fmt.Errorf("Unable to process DNS resultset (fw %d): %s", r.data.Fw, err.Error())
            }
        }
    }

    return nil
}

// The firmware version used by the probe that generated this result.
func (r *Result) Fw() int {
    return r.data.Fw
}

// IP version: "4" or "6" (optional).
func (r *Result) Af() int {
    return r.data.Af
}

// IP address of the destination (optional).
func (r *Result) DstAddr() string {
    return r.data.DstAddr
}

// Hostname of the destination (optional).
func (r *Result) DstName() string {
    return r.data.DstName
}

// IP address of the source (optional).
func (r *Result) From() string {
    return r.data.From
}

// Last time synchronised. How long ago (in seconds) the clock of the probe
// was found to be in sync with that of a controller. The value -1 is used
// to indicate that the probe does not know whether it is in sync.
func (r *Result) Lts() int {
    return r.data.Lts
}

// Measurement identifier.
func (r *Result) MsmId() int {
    return r.data.MsmId
}

// Source probe ID.
func (r *Result) PrbId() int {
    return r.data.PrbId
}

// Protocol, "TCP" or "UDP".
func (r *Result) Proto() string {
    return r.data.Proto
}

// Query payload buffer which was sent to the server, UU encoded (optional).
func (r *Result) Qbuf() string {
    return r.data.Qbuf
}

// Retry count (optional).
func (r *Result) Retry() int {
    return r.data.Retry
}

// Start time, in Unix timestamp.
func (r *Result) Timestamp() int {
    return r.data.Timestamp
}

// The type of measurement within this result.
func (r *Result) Type() string {
    return r.data.Type
}

// DNS error message, nil if the type of measurement is not "dns" (optional).
func (r *Result) DnsError() *dns.Error {
    return r.dnsError
}

// DNS response from the DNS server, nil if the type of measurement is
// not "dns" (optional).
func (r *Result) DnsResult() *dns.Result {
    return r.dnsResult
}

// An array of objects containing the DNS results when querying multiple
// local resolvers, empty if the type of measurement is not "dns" (optional).
func (r *Result) DnsResultsets() []*dns.Resultset {
    return r.dnsResultsets
}

// Decode the Qbuf() as a DNS message, returns a *Msg from the
// github.com/miekg/dns package.
func (r *Result) DnsUnpackQbuf() (*mdns.Msg, error) {
    if r.data.Type != "dns" {
        return nil, fmt.Errorf("Result type is not DNS")
    }

    m := &mdns.Msg{}
    if r.data.Qbuf != "" {
        b, err := base64.StdEncoding.DecodeString(r.data.Qbuf)
        if err != nil {
            return nil, err
        }
        if err := m.Unpack(b); err != nil {
            return nil, err
        }
    }
    return m, nil
}
