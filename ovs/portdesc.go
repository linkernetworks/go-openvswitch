// Copyright 2017 DigitalOcean.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ovs

import (
	"errors"
	"strconv"
	"strings"
)

var (
	// ErrInvalidPortDesc is returned when port statistics from 'ovs-ofctl
	// dump-ports' do not match the expected output format.
	ErrInvalidPortDesc = errors.New("invalid port description")
)

// PortDesc contains a variety of statistics about an Open vSwitch port,
// including its port ID and numbers about packet receive and transmit
// operations.
type PortDesc struct {
	// PortID specifies the OVS port ID which this PortDesc refers to.
	ID   int32
	Name string

	MACAddress string
}

// UnmarshalText unmarshals a PortDesc from textual form as output by
// 'ovs-ofctl dump-ports-desc':
// LOCAL(mybr0): addr:1a:e4:7d:9c:72:45
//   config:     PORT_DOWN
//   state:      LINK_DOWN
//   current: 10GB-FD COPPER
//   speed: 0 Mbps now, 0 Mbps max
func (p *PortDesc) UnmarshalText(b []byte) error {
	// Make a copy per documentation for encoding.TextUnmarshaler.
	s := string(b)

	// Constants only needed within this method, to avoid polluting the
	// package namespace with generic names
	ss := strings.Fields(s)

	addr := ss[1]
	if len(addr) != 22 || addr[0:4] != "addr" {
		return nil
	}

	//We only parse the ID/Name/MAC now.
	left := strings.IndexByte(ss[0], '(')
	right := strings.IndexByte(ss[0], ')')
	p.Name = ss[0][left+1 : right]

	//Name
	p.MACAddress = ss[1][5:]

	//ID
	id := ss[0][0:left]
	if id == "LOCAL" {
		p.ID = -1
	} else {
		tmp, _ := strconv.Atoi(id)
		p.ID = (int32)(tmp)
	}
	return nil
}
