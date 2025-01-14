/*
Copyright 2018-2019 The Doctl Authors All rights reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package doctl

import "fmt"

// MissingArgsErr is returned when there are too few arguments for a command.
type MissingArgsErr struct {
	Command string
}

var _ error = &MissingArgsErr{}

// NewMissingArgsErr creates a MissingArgsErr instance.
func NewMissingArgsErr(cmd string) *MissingArgsErr {
	return &MissingArgsErr{Command: cmd}
}

func (e *MissingArgsErr) Error() string {
	return fmt.Sprintf("(%s) command is missing required arguments", e.Command)
}

// InvalidURNErr is returned when the URN format is not valid.
type InvalidURNErr struct {
	URN string
}

var _ error = &InvalidURNErr{}

// NewInvalidURNErr creates a InvalidURNErr instance.
func NewInvalidURNErr(urn string) *InvalidURNErr {
	return &InvalidURNErr{URN: urn}
}

func (e *InvalidURNErr) Error() string {
	return fmt.Sprintf("URN must be in the format \"do:<resource_type>:<resource_id>\"")
}
