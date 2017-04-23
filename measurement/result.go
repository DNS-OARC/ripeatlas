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
    "encoding/json"
    "fmt"

    "github.com/DNS-OARC/ripeatlas/measurement/dns"
)

type Result struct {
    Data struct {
        Fw         int             `json:"fw"`
        Af         int             `json:"af"`        // [optional] IP version: "4" or "6" (int)
        DstAddr    string          `json:"dst_addr"`  // [optional] IP address of the destination (string)
        DstName    string          `json:"dst_name"`  // [optional] hostname of the destination (string)
        Error      json.RawMessage `json:"error"`     // [optional] error message (associative array)
        From       string          `json:"from"`      // [optional] IP address of the source (string)
        Lts        int             `json:"lts"`       // last time synchronised. How long ago (in seconds) the clock of the probe was found to be in sync with that of a controller. The value -1 is used to indicate that the probe does not know whether it is in sync (from 4650) (int)
        MsmId      int             `json:"msm_id"`    // measurement identifier (int)
        PrbId      int             `json:"prb_id"`    // source probe ID (int)
        Proto      string          `json:"proto"`     // "TCP" or "UDP" (string)
        Qbuf       string          `json:"qbuf"`      // [optional] query payload buffer which was sent to the server, UU encoded (string)
        Result     json.RawMessage `json:"result"`    // [optional] response from the DNS server (associative array)
        Resultsets json.RawMessage `json:"resultset"` // [optional] an array of objects containing all the fields of a DNS result object, except for the fields: fw, from, msm_id, prb_id, and type. Available for queries sent to each local resolver.
        Retry      int             `json:"retry"`     // [optional] retry count (int)
        Timestamp  int             `json:"timestamp"` // start time, in Unix timestamp (int)
        Type       string          `json:"type"`      // "dns" (string)
    }

    DnsError      dns.Error
    DnsResult     dns.Result
    DnsResultsets []dns.Resultset
}

func (r *Result) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.Data); err != nil {
        return err
    }

    if r.Data.Type == "dns" {
        if r.Data.Error != nil {
            if err := json.Unmarshal(r.Data.Error, &r.DnsError); err != nil {
                return fmt.Errorf("Unable to process DNS error: %s", err.Error())
            }
        }
        if r.Data.Result != nil {
            if err := json.Unmarshal(r.Data.Result, &r.DnsResult); err != nil {
                return fmt.Errorf("Unable to process DNS result: %s", err.Error())
            }
        }
        if r.Data.Resultsets != nil {
            if err := json.Unmarshal(r.Data.Resultsets, &r.DnsResultsets); err != nil {
                return fmt.Errorf("Unable to process DNS resultset: %s", err.Error())
            }
        }
    }

    return nil
}

func (r *Result) Fw() int {
    return r.Data.Fw
}

func (r *Result) Af() int {
    return r.Data.Af
}

func (r *Result) DstAddr() string {
    return r.Data.DstAddr
}

func (r *Result) DstName() string {
    return r.Data.DstName
}

func (r *Result) From() string {
    return r.Data.From
}

func (r *Result) Lts() int {
    return r.Data.Lts
}

func (r *Result) MsmId() int {
    return r.Data.MsmId
}

func (r *Result) PrbId() int {
    return r.Data.PrbId
}

func (r *Result) Proto() string {
    return r.Data.Proto
}

func (r *Result) Qbuf() string {
    return r.Data.Qbuf
}

func (r *Result) Retry() int {
    return r.Data.Retry
}

func (r *Result) Timestamp() int {
    return r.Data.Timestamp
}

func (r *Result) Type() string {
    return r.Data.Type
}
