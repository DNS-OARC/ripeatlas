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

import "encoding/json"

type Answer struct {
    Mname  string   `json:"MNAME"`  // domain name, RFC 1035, 3.1.13 (string)
    Name   string   `json:"NAME"`   // domain name. (string)
    Rdata  []string `json:"RDATA"`  // [type TXT] txt value (from 4710, list of strings, before it was a single string)
    Rname  string   `json:"RNAME"`  // [if type SOA] mailbox, RFC 1035 3.3.13 (string)
    Serial int      `json:"SERIAL"` // [type SOA] zone serial number, RFC 1035 3.3.13 (string)
    Ttl    int      `json:"TTL"`    // [type SOA] time to live, RFC 1035 4.1.3 (int)
    Type   string   `json:"TYPE"`   // RR "SOA" or "TXT" (string), RFC 1035
}

type AnswerContainer struct {
    Answer *Answer
}

func (a *AnswerContainer) UnmarshalJSON(b []byte) error {
    a.Answer = &Answer{}
    if err := json.Unmarshal(b, &a.Answer); err != nil {
        return err
    }
    return nil
}

func (a *AnswerContainer) Contained() *Answer {
    if a.Answer == nil {
        a.Answer = &Answer{}
    }
    return a.Answer
}

func (a *AnswerContainer) Mname() string {
    return a.Contained().Mname
}

func (a *AnswerContainer) Name() string {
    return a.Contained().Name
}

func (a *AnswerContainer) Rdata() []string {
    return a.Contained().Rdata
}

func (a *AnswerContainer) Rname() string {
    return a.Contained().Rname
}

func (a *AnswerContainer) Serial() int {
    return a.Contained().Serial
}

func (a *AnswerContainer) Ttl() int {
    return a.Contained().Ttl
}

func (a *AnswerContainer) Type() string {
    return a.Contained().Type
}
