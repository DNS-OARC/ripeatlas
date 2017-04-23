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

type Error struct {
    Timeout     int    `json:"timeout"`     // query timeout (int)
    Getaddrinfo string `json:"getaddrinfo"` // error message (string)
}

type ErrorContainer struct {
    Error *Error
}

func (e *ErrorContainer) UnmarshalJSON(b []byte) error {
    e.Error = &Error{}
    if err := json.Unmarshal(b, &e.Error); err != nil {
        return err
    }
    return nil
}

func (e *ErrorContainer) Contained() *Error {
    if e.Error == nil {
        e.Error = &Error{}
    }
    return e.Error
}

func (e *ErrorContainer) Timeout() int {
    return e.Contained().Timeout
}

func (e *ErrorContainer) Getaddrinfo() string {
    return e.Contained().Getaddrinfo
}
