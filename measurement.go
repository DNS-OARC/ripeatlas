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

type Measurement struct {
    Fw        int                  `json:"fw"`
    Af        int                  `json:"af"`        // [optional] IP version: "4" or "6" (int)
    DstAddr   string               `json:"dst_addr"`  // [optional] IP address of the destination (string)
    DstName   string               `json:"dst_name"`  // [optional] hostname of the destination (string)
    Error     ErrorContainer       `json:"error"`     // [optional] error message (associative array)
    From      string               `json:"from"`      // [optional] IP address of the source (string)
    Lts       int                  `json:"lts"`       // last time synchronised. How long ago (in seconds) the clock of the probe was found to be in sync with that of a controller. The value -1 is used to indicate that the probe does not know whether it is in sync (from 4650) (int)
    MsmId     int                  `json:"msm_id"`    // measurement identifier (int)
    PrbId     int                  `json:"prb_id"`    // source probe ID (int)
    Proto     string               `json:"proto"`     // "TCP" or "UDP" (string)
    Qbuf      string               `json:"qbuf"`      // [optional] query payload buffer which was sent to the server, UU encoded (string)
    Result    ResultContainer      `json:"result"`    // [optional] response from the DNS server (associative array)
    Resultset []ResultsetContainer `json:"resultset"` // [optional] an array of objects containing all the fields of a DNS result object, except for the fields: fw, from, msm_id, prb_id, and type. Available for queries sent to each local resolver.
    Retry     int                  `json:"retry"`     // [optional] retry count (int)
    Timestamp int                  `json:"timestamp"` // start time, in Unix timestamp (int)
    Type      string               `json:"type"`      // "dns" (string)
}

type MeasurementContainer struct {
    Measurement *Measurement
}

func (m *MeasurementContainer) UnmarshalJSON(b []byte) error {
    m.Measurement = &Measurement{}
    if err := json.Unmarshal(b, &m.Measurement); err != nil {
        return err
    }
    return nil
}

func (m *MeasurementContainer) Contained() *Measurement {
    if m.Measurement == nil {
        m.Measurement = &Measurement{}
    }
    return m.Measurement
}

func (m *MeasurementContainer) Fw() int {
    return m.Contained().Fw
}

func (m *MeasurementContainer) Af() int {
    return m.Contained().Af
}

func (m *MeasurementContainer) DstAddr() string {
    return m.Contained().DstAddr
}

func (m *MeasurementContainer) DstName() string {
    return m.Contained().DstName
}

func (m *MeasurementContainer) Error() ErrorContainer {
    return m.Contained().Error
}

func (m *MeasurementContainer) From() string {
    return m.Contained().From
}

func (m *MeasurementContainer) Lts() int {
    return m.Contained().Lts
}

func (m *MeasurementContainer) MsmId() int {
    return m.Contained().MsmId
}

func (m *MeasurementContainer) PrbId() int {
    return m.Contained().PrbId
}

func (m *MeasurementContainer) Proto() string {
    return m.Contained().Proto
}

func (m *MeasurementContainer) Qbuf() string {
    return m.Contained().Qbuf
}

func (m *MeasurementContainer) Result() ResultContainer {
    return m.Contained().Result
}

func (m *MeasurementContainer) Resultset() []ResultsetContainer {
    return m.Contained().Resultset
}

func (m *MeasurementContainer) Retry() int {
    return m.Contained().Retry
}

func (m *MeasurementContainer) Timestamp() int {
    return m.Contained().Timestamp
}

func (m *MeasurementContainer) Type() string {
    return m.Contained().Type
}
