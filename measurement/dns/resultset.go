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

type Resultset struct {
    Data struct {
        Af        int    `json:"af"`        // [optional] IP version: "4" or "6" (int)
        DstAddr   string `json:"dst_addr"`  // [optional] IP address of the destination (string)
        DstName   string `json:"dst_name"`  // [optional] hostname of the destination (string)
        Error     Error  `json:"error"`     // [optional] error message (associative array)
        Lts       int    `json:"lts"`       // last time synchronised. How long ago (in seconds) the clock of the probe was found to be in sync with that of a controller. The value -1 is used to indicate that the probe does not know whether it is in sync (from 4650) (int)
        Proto     string `json:"proto"`     // "TCP" or "UDP" (string)
        Qbuf      string `json:"qbuf"`      // [optional] query payload buffer which was sent to the server, UU encoded (string)
        Result    Result `json:"result"`    // [optional] response from the DNS server (associative array)
        Retry     int    `json:"retry"`     // [optional] retry count (int)
        Timestamp int    `json:"timestamp"` // start time, in Unix timestamp (int)
    }
}

func (r *Resultset) UnmarshalJSON(b []byte) error {
    if err := json.Unmarshal(b, &r.Data); err != nil {
        return err
    }
    return nil
}

func (r *Resultset) Af() int {
    return r.Data.Af
}

func (r *Resultset) DstAddr() string {
    return r.Data.DstAddr
}

func (r *Resultset) DstName() string {
    return r.Data.DstName
}

func (r *Resultset) Error() *Error {
    return &r.Data.Error
}

func (r *Resultset) Lts() int {
    return r.Data.Lts
}

func (r *Resultset) Proto() string {
    return r.Data.Proto
}

func (r *Resultset) Qbuf() string {
    return r.Data.Qbuf
}

func (r *Resultset) Result() *Result {
    return &r.Data.Result
}

func (r *Resultset) Retry() int {
    return r.Data.Retry
}

func (r *Resultset) Timestamp() int {
    return r.Data.Timestamp
}
