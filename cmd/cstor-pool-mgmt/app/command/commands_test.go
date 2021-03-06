// Copyright © 2017-2019 The OpenEBS Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package command

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

// TestNewCStorPoolMgmt is to test cstor-pool-mgmt command.
func TestNewCStorPoolMgmt(t *testing.T) {
	cases := []struct {
		use string
	}{
		{"start"},
	}
	cmd, err := NewCStorPoolMgmt()
	if err != nil {
		t.Errorf("Unable to Instantiate cstor-pool-mgmt")
	}
	cmds := cmd.Commands()
	if len(cmds) != len(cases) {
		t.Errorf("ExpectedCommands: %d ActualCommands: '%d'", len(cases), len(cmds))
	}
	for i, c := range cases {
		if c.use != cmds[i].Use {
			t.Errorf("ExpectedCommand: '%s' ActualCommand: '%s'", c.use, cmds[i].Use)
		}
	}
}

// TestRun is to test running cstor-pool-mgmt without sub-commands.
func TestRun(t *testing.T) {
	var cmd *cobra.Command
	err := Run(cmd)
	if err != nil {
		t.Errorf("Expected: '%s' Actual: '%s'", "nil", err)
	}
}

// TestNewCmdOptions is to test type of CLI command.
func TestNewCmdOptions(t *testing.T) {
	var expectedCmd *cobra.Command
	gotCmd := NewCmdOptions()
	if reflect.TypeOf(gotCmd) != reflect.TypeOf(expectedCmd) {
		t.Errorf("Expected: '%s' Actual: '%v'", reflect.TypeOf(gotCmd), reflect.TypeOf(expectedCmd))
	}
}
