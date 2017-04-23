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

package ripeatlas

import (
    "encoding/base64"
    "encoding/json"
    "github.com/miekg/dns"
)

type Result struct {
    Ancount int               `json:"ANCOUNT"`  // answer count, RFC 1035 4.1.1 (int)
    Arcount int               `json:"ARCOUNT"`  // additional record count, RFC 1035, 4.1.1 (int)
    Id      int               `json:"ID"`       // query ID, RFC 1035 4.1.1 (int)
    Nscount int               `json:"NSCOUNT"`  // name server count (int)
    Qdcount int               `json:"QDCOUNT"`  // number of queries (int)
    Abuf    string            `json:"abuf"`     // answer payload buffer from the server, UU encoded (string)
    Answers []AnswerContainer `json:"answers"`  // first two records from the response decoded by the probe, if they are TXT or SOA; other RR can be decoded from "abuf" (array)
    Rt      float64           `json:"rt"`       // [optional] response time in milli seconds (float)
    Size    int               `json:"size"`     // [optional] response size (int)
    SrcAddr string            `json:"src_addr"` // [optional] the source IP address added by the probe (string).
    Subid   int               `json:"subid"`    // [optional] sequence number of this result within a group of results, available if the resolution is done by the probe's local resolver
    Submax  int               `json:"submax"`   // [optional] total number of results within a group (int)
}

func (r *Result) UnpackAbuf() (*dns.Msg, error) {
    m := &dns.Msg{}
    if r.Abuf != "" {
        b, err := base64.StdEncoding.DecodeString(r.Abuf)
        if err != nil {
            return nil, err
        }
        if err := m.Unpack(b); err != nil {
            return nil, err
        }
    }
    return m, nil
}

type ResultContainer struct {
    Result *Result
}

func (r *ResultContainer) UnmarshalJSON(b []byte) error {
    r.Result = &Result{}
    if err := json.Unmarshal(b, &r.Result); err != nil {
        return err
    }
    return nil
}

func (r *ResultContainer) Contained() *Result {
    if r.Result == nil {
        r.Result = &Result{}
    }
    return r.Result
}

func (r *ResultContainer) Ancount() int {
    return r.Contained().Ancount
}

func (r *ResultContainer) Arcount() int {
    return r.Contained().Arcount
}

func (r *ResultContainer) Id() int {
    return r.Contained().Id
}

func (r *ResultContainer) Nscount() int {
    return r.Contained().Nscount
}

func (r *ResultContainer) Qdcount() int {
    return r.Contained().Qdcount
}

func (r *ResultContainer) Abuf() string {
    return r.Contained().Abuf
}

func (r *ResultContainer) UnpackAbuf() (*dns.Msg, error) {
    return r.Contained().UnpackAbuf()
}

func (r *ResultContainer) Answers() []AnswerContainer {
    return r.Contained().Answers
}

func (r *ResultContainer) Rt() float64 {
    return r.Contained().Rt
}

func (r *ResultContainer) Size() int {
    return r.Contained().Size
}

func (r *ResultContainer) SrcAddr() string {
    return r.Contained().SrcAddr
}

func (r *ResultContainer) Subid() int {
    return r.Contained().Subid
}

func (r *ResultContainer) Submax() int {
    return r.Contained().Submax
}
