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
	port int
	bind string
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a websocket server",
	Long: `
Start a websocket server`,
	Run: func(cmd *cobra.Command, args []string) {
		addr := fmt.Sprintf("%s:%d", bind, port)
		logger.Infof("wsgo server listen at: %s", addr)
		logger.Fatal(http.ListenAndServe(addr, &server.Server{
			ServerType: server.Echo,
		}))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVarP(&port, "port", "p", 8080, "Listening Port")
	serverCmd.Flags().StringVarP(&bind, "bind", "b", "0.0.0.0", "Bind Address")
}
