// Copyright (C) 2015-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"errors"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/nlamirault/abraracourcix/server/abraracourcixd"
)

var (
	config string
)

type runCmd struct {
	out io.Writer
}

func newRunCmd(out io.Writer) *cobra.Command {
	runCmd := &runCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "run",
		Short: "Run a Abraracourcix server",
		Long:  `All software has infos.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(config) == 0 {
				return errors.New("missing configuration filename")
			}
			runServer(runCmd.out, config)
			return nil
		},
	}
	cmd.PersistentFlags().StringVar(&config, "config", "abraracourcix.toml", "Configuration filename")
	return cmd
}

func runServer(out io.Writer, config string) {
	glog.V(2).Infof("Start the server")
	abraracourcixd.StartServer(config)
}
