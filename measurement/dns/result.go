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

package dns

import (
    "encoding/base64"
    "encoding/json"

    "github.com/miekg/dns"
)

type Result struct {
    Data struct {
        Ancount int      `json:"ANCOUNT"`  // answer count, RFC 1035 4.1.1 (int)
        Arcount int      `json:"ARCOUNT"`  // additional record count, RFC 1035, 4.1.1 (int)
        Id      int      `json:"ID"`       // query ID, RFC 1035 4.1.1 (int)
        Nscount int      `json:"NSCOUNT"`  // name server count (int)
        Qdcount int      `json:"QDCOUNT"`  // number of queries (int)
        Abuf    string   `json:"abuf"`     // answer payload buffer from the server, UU encoded (string)
        Answers []Answer `json:"answers"`  // first two records from the response decoded by the probe, if they are TXT or SOA; other RR can be decoded from "abuf" (array)
        Rt      float64  `json:"rt"`       // [optional] response time in milli seconds (float)
        Size    int      `json:"size"`     // [optional] response size (int)
        SrcAddr string   `json:"src_addr"` // [optional] the source IP address added by the probe (string).
        Subid   int      `json:"subid"`    // [optional] sequence number of this result within a group of results, available if the resolution is done by the probe's local resolver
        Submax  int      `json:"submax"`   // [optional] total number of results within a group (int)
    }
}

func (r *Result) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.Data); err != nil {
        return err
    }
    return nil
}

func (r *Result) Ancount() int {
    return r.Data.Ancount
}

func (r *Result) Arcount() int {
    return r.Data.Arcount
}

func (r *Result) Id() int {
    return r.Data.Id
}

func (r *Result) Nscount() int {
    return r.Data.Nscount
}

func (r *Result) Qdcount() int {
    return r.Data.Qdcount
}

func (r *Result) Abuf() string {
    return r.Data.Abuf
}

func (r *Result) Answers() []Answer {
    return r.Data.Answers
}

func (r *Result) Rt() float64 {
    return r.Data.Rt
}

func (r *Result) Size() int {
    return r.Data.Size
}

func (r *Result) SrcAddr() string {
    return r.Data.SrcAddr
}

func (r *Result) Subid() int {
    return r.Data.Subid
}

func (r *Result) Submax() int {
    return r.Data.Submax
}

func (r *Result) UnpackAbuf() (*dns.Msg, error) {
    m := &dns.Msg{}
    if r.Data.Abuf != "" {
        b, err := base64.StdEncoding.DecodeString(r.Data.Abuf)
        if err != nil {
            return nil, err
        }
        if err := m.Unpack(b); err != nil {
            return nil, err
        }
    }
    return m, nil
}
