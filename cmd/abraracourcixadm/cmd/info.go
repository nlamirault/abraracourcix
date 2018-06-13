// Copyright (C) 2016-2018 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"io"

	"github.com/golang/glog"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/nlamirault/abraracourcix/cmd/utils"
	"github.com/nlamirault/abraracourcix/pb/info"
)

const (
	serverService = "abraracourcixd"
)

type infoCmd struct {
	out io.Writer
}

func newInfoCmd(out io.Writer) *cobra.Command {
	infoCmd := &infoCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "info",
		Short: "Print info about the Abraracourcix server",
		Long:  `All software has infos.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			return printInfo(client, infoCmd.out)
		},
	}
	return cmd
}

func printInfo(gRPCClient *utils.GRPCClient, out io.Writer) error {
	glog.V(1).Infof("Retrieve informations")

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := info.NewInfoServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, &info.GetInfoRequest{})
	if err != nil {
		return err
	}
	return printServerInformations(out, resp)
}

func printServerInformations(out io.Writer, resp *info.GetInfoResponse) error {

	table := tablewriter.NewWriter(out)
	table.SetRowLine(true)
	table.SetAutoWrapText(false)
	table.SetHeader([]string{"Service", "URI", "Version", "Status"})

	table.Append([]string{
		serverService,
		utils.ServerAddress,
		resp.Version,
		"OK",
	})
	for _, ws := range resp.Webservices {
		table.Append([]string{
			ws.Name,
			ws.Endpoint,
			ws.Text,
			ws.Status,
		})
	}
	table.Render()
	return nil
}
