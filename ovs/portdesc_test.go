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
	"reflect"
	"testing"
)

func TestPortDescUnmarshalText(t *testing.T) {
	var tests = []struct {
		desc string
		s    string
		p    *PortDesc
		err  error
	}{
		{
			desc: "empty string",
			err:  ErrInvalidPortDesc,
		},
		{
			desc: "Unused config",
			s:    "config:     PORT_DOWN",
			err:  ErrIgnoreUnusedDesc,
		},
		{
			desc: "OK",
			s:    "LOCAL(mybr0): addr:1a:e4:7d:9c:72:45",
			p: &PortDesc{
				ID:         PortLOCAL,
				Name:       "mybr0",
				MACAddress: "1a:e4:7d:9c:72:45",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			p := new(PortDesc)
			err := p.UnmarshalText([]byte(tt.s))

			if want, got := errStr(tt.err), errStr(err); want != got {
				t.Fatalf("unexpected error:\n- want: %v\n-  got: %v",
					want, got)
			}
			if err != nil {
				return
			}

			if want, got := tt.p, p; !reflect.DeepEqual(want, got) {
				t.Fatalf("unexpected PortDesc:\n- want: %#v\n-  got: %#v",
					want, got)
			}
		})
	}
}
