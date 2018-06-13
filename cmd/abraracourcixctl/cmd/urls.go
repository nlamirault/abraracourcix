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
	"fmt"
	"io"

	"github.com/golang/glog"
	"github.com/spf13/cobra"

	"github.com/nlamirault/abraracourcix/cmd/utils"
	"github.com/nlamirault/abraracourcix/pb/v2beta"
)

var (
	key  string
	link string
)

type urlCmd struct {
	out io.Writer
}

func newLeagueCmd(out io.Writer) *cobra.Command {
	urlCmd := &urlCmd{
		out: out,
	}

	cmd := &cobra.Command{
		Use:   "url",
		Short: "Manage URLs",
		Long:  "Manage URLs. See List, Get, ... subcommands.",
		RunE:  nil,
	}
	listUrlCmd := &cobra.Command{
		Use:   "list",
		Short: "List all URLs",
		RunE: func(cmd *cobra.Command, args []string) error {
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			if err := urlCmd.listURLs(client); err != nil {
				return err
			}
			return nil
		},
	}
	getUrlCmd := &cobra.Command{
		Use:   "get",
		Short: "Retreive an URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(key) == 0 {
				return errors.New("missing URL key")
			}
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			if err := urlCmd.getUrl(client, key); err != nil {
				return err
			}
			return nil
		},
	}
	addUrlCmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new URL",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(link) == 0 {
				return errors.New("missing URL link")
			}
			client, err := utils.NewGRPCClient(cmd)
			if err != nil {
				return err
			}
			if err := urlCmd.addUrl(client, link); err != nil {
				return err
			}
			return nil
		},
	}

	getUrlCmd.PersistentFlags().StringVar(&key, "key", "", "URL's key")
	addUrlCmd.PersistentFlags().StringVar(&link, "link", "", "URL link")
	cmd.AddCommand(listUrlCmd)
	cmd.AddCommand(getUrlCmd)
	cmd.AddCommand(addUrlCmd)
	return cmd
}

func (cmd urlCmd) listURLs(gRPCClient *utils.GRPCClient) error {
	glog.V(1).Info("List all URLs")

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v2beta.NewUrlServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}
	resp, err := client.List(ctx, &v2beta.GetUrlsRequest{})
	if err != nil {
		return err
	}
	fmt.Printf(utils.GreenOut("URLs:\n"))
	for _, key := range resp.Keys {
		fmt.Printf("- %s\n", key)
	}
	return nil
}

func (cmd urlCmd) getUrl(gRPCClient *utils.GRPCClient, key string) error {
	glog.V(1).Info("Get URL %s", key)

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v2beta.NewUrlServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}
	resp, err := client.Get(ctx, &v2beta.GetUrlRequest{
		Key: key,
	})
	if err != nil {
		return err
	}
	fmt.Printf(utils.GreenOut("URL:\n"))
	fmt.Printf("- Key: %s\n", resp.Url.Key)
	fmt.Printf("- Link: %s\n", resp.Url.Link)
	fmt.Printf("- Date: %s\n", resp.Url.Creation)
	return nil
}

func (cmd urlCmd) addUrl(gRPCClient *utils.GRPCClient, link string) error {
	glog.V(1).Info("Add a new URL %s", link)

	conn, err := gRPCClient.GetConn()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := v2beta.NewUrlServiceClient(conn)
	ctx, err := gRPCClient.GetContext(cliName)
	if err != nil {
		return err
	}
	resp, err := client.Create(ctx, &v2beta.CreateUrlRequest{
		Link: link,
	})
	if err != nil {
		return err
	}
	fmt.Printf(utils.GreenOut("URL:\n"))
	fmt.Printf("- Key: %s\n", resp.Url.Key)
	fmt.Printf("- Link: %s\n", resp.Url.Link)
	fmt.Printf("- Date: %s\n", resp.Url.Creation)
	return nil
}
