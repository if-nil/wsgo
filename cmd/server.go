/*
Copyright © 2023 ifNil ifnil.git@gmail.com

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
package cmd

import (
	"fmt"
	"github.com/if-nil/wsgo/logger"
	"github.com/if-nil/wsgo/server"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	portArg int
	bindArg string
	fileArg string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a websocket server",
	Long: `
Start a websocket server`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", bindArg, portArg)
		logger.Infof("wsgo server listen at: %s", addr)
		serverType := server.Echo
		if fileArg != "" {
			serverType = server.LuaServer
		}
		logger.Fatal(http.ListenAndServe(addr, &server.Server{
			ServerType: serverType,
			LuaPool:    server.New(fileArg),
		}))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVarP(&portArg, "port", "p", 8080, "Listening Port")
	serverCmd.Flags().StringVarP(&bindArg, "bind", "b", "0.0.0.0", "Bind Address")
	serverCmd.Flags().StringVarP(&fileArg, "file", "f", "", "lua plugin file")
}
