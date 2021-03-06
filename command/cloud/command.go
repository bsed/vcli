// Copyright 2016 Sisa-Tech Pty Ltd
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
package cmdcloud

import (
	"github.com/alecthomas/kingpin"
	"github.com/sisatech/vcli/command"
	"github.com/sisatech/vcli/shared"
)

// Command ...
type Command struct {
	*kingpin.CmdClause
	arg string
}

// New ...
func New() *Command {

	return &Command{}

}

// Attach ...
func (cmd *Command) Attach(parent command.Node) {

	cmd.CmdClause = parent.Command("cloud", shared.Catenate(`The cloud
		subcommand contains all functions that interact directly with
		remote cloud infrastructure.`))

	cmd.PreAction(cmd.preaction)

	// Add new sub-commands below.
	newInfrastructureCmd().Attach(cmd)
	newUploadTemplateCmd().Attach(cmd)
	newRemoveTemplateCmd().Attach(cmd)
}

func (cmd *Command) preaction(ctx *kingpin.ParseContext) error {

	return command.NodeOnlyCheck(cmd, ctx)

}
