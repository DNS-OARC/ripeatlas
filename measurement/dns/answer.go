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

import "encoding/json"

type Answer struct {
    Data struct {
        Mname  string   `json:"MNAME"`  // domain name, RFC 1035, 3.1.13 (string)
        Name   string   `json:"NAME"`   // domain name. (string)
        Rdata  []string `json:"RDATA"`  // [type TXT] txt value (from 4710, list of strings, before it was a single string)
        Rname  string   `json:"RNAME"`  // [if type SOA] mailbox, RFC 1035 3.3.13 (string)
        Serial int      `json:"SERIAL"` // [type SOA] zone serial number, RFC 1035 3.3.13 (string)
        Ttl    int      `json:"TTL"`    // [type SOA] time to live, RFC 1035 4.1.3 (int)
        Type   string   `json:"TYPE"`   // RR "SOA" or "TXT" (string), RFC 1035
    }
}

func (a *Answer) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &a.Data); err != nil {
        return err
    }
    return nil
}

func (a *Answer) Mname() string {
    return a.Data.Mname
}

func (a *Answer) Name() string {
    return a.Data.Name
}

func (a *Answer) Rdata() []string {
    return a.Data.Rdata
}

func (a *Answer) Rname() string {
    return a.Data.Rname
}

func (a *Answer) Serial() int {
    return a.Data.Serial
}

func (a *Answer) Ttl() int {
    return a.Data.Ttl
}

func (a *Answer) Type() string {
    return a.Data.Type
}
