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
    "encoding/json"
    "fmt"
    "io/ioutil"
    "os"
)

type File struct {
    Name         string
    measurements []MeasurementContainer
}

func NewFile() *File {
    return &File{}
}

func (f *File) Measurements() []MeasurementContainer {
    return f.measurements
}

func (f *File) Read() error {
    if f.Name == "" {
        return fmt.Errorf("Name must be set")
    }

    file, err := os.Open(f.Name)
    defer file.Close()
    if err != nil {
        return fmt.Errorf("os.Open(%s): %s", f.Name, err.Error())
    }

    content, err := ioutil.ReadAll(file)
    if err != nil {
        return fmt.Errorf("ioutil.ReadAll(%s): %s", f.Name, err.Error())
    }

    if err := json.Unmarshal(content, &f.measurements); err != nil {
        return fmt.Errorf("json.Unmarshal(%s): %s", f.Name, err.Error())
    }

    return nil
}
