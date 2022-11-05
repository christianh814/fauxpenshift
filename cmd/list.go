/*
Copyright © 2022 Christian Hernandez christian@chernand.io

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
	"os"

	"github.com/christianh814/fauxpenshift/pkg/container"
	"github.com/christianh814/fauxpenshift/pkg/microshift"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// listCmd lists currently running clusters
var listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"get", "ls"},
	Short:   "Lists all your instances",
	Long:    `Shows you a simple list of your clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		// Set runtime
		// TODO: Want to probably set this centrally somehow
		var rt string = microshift.Runtime
		if os.Getenv("FAUXPENSHIFT_SET_RUNTIME") == "docker" {
			rt = "docker"
		}

		// Display current clusters
		output, err := container.DisplayMicroshiftInstance(rt, "label=fauxpenshift=instance")

		// check for errors
		if err != nil {
			log.Fatal(err)
		}

		// If we're here, we can display the output
		fmt.Println(string(output))
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
