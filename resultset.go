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

type Resultset struct {
    Af        int             `json:"af"`        // [optional] IP version: "4" or "6" (int)
    DstAddr   string          `json:"dst_addr"`  // [optional] IP address of the destination (string)
    DstName   string          `json:"dst_name"`  // [optional] hostname of the destination (string)
    Error     ErrorContainer  `json:"error"`     // [optional] error message (associative array)
    Lts       int             `json:"lts"`       // last time synchronised. How long ago (in seconds) the clock of the probe was found to be in sync with that of a controller. The value -1 is used to indicate that the probe does not know whether it is in sync (from 4650) (int)
    Proto     string          `json:"proto"`     // "TCP" or "UDP" (string)
    Qbuf      string          `json:"qbuf"`      // [optional] query payload buffer which was sent to the server, UU encoded (string)
    Result    ResultContainer `json:"result"`    // [optional] response from the DNS server (associative array)
    Retry     int             `json:"retry"`     // [optional] retry count (int)
    Timestamp int             `json:"timestamp"` // start time, in Unix timestamp (int)
}

type ResultsetContainer struct {
    Resultset *Resultset
}

func (r *ResultsetContainer) UnmarshalJSON(b []byte) error {
    r.Resultset = &Resultset{}
    if err := json.Unmarshal(b, &r.Resultset); err != nil {
        return err
    }
    return nil
}

func (r *ResultsetContainer) Contained() *Resultset {
    if r.Resultset == nil {
        r.Resultset = &Resultset{}
    }
    return r.Resultset
}

func (r *ResultsetContainer) Af() int {
    return r.Contained().Af
}

func (r *ResultsetContainer) DstAddr() string {
    return r.Contained().DstAddr
}

func (r *ResultsetContainer) DstName() string {
    return r.Contained().DstName
}

func (r *ResultsetContainer) Error() ErrorContainer {
    return r.Contained().Error
}

func (r *ResultsetContainer) Lts() int {
    return r.Contained().Lts
}

func (r *ResultsetContainer) Proto() string {
    return r.Contained().Proto
}

func (r *ResultsetContainer) Qbuf() string {
    return r.Contained().Qbuf
}

func (r *ResultsetContainer) Result() ResultContainer {
    return r.Contained().Result
}

func (r *ResultsetContainer) Retry() int {
    return r.Contained().Retry
}

func (r *ResultsetContainer) Timestamp() int {
    return r.Contained().Timestamp
}
