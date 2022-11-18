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
	"github.com/christianh814/fauxpenshift/pkg/container"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:     "destroy",
	Aliases: []string{"rm"},
	Short:   "Destroys a cluster",
	Long: `This will destroy a cluster. There is no way
to "save" your cluster.

The functionality is possible, but not when using 
this tool. PRs are welcome!`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Setting runtime to ", Runtime)

		//Stop the instance
		log.Info("Destroying Microshift instance")
		if err := container.StopMicroshiftContainer(Runtime, "fauxpenshift"); err != nil {
			log.Fatal(err)
		}

		//cleanup volume
		if err := container.CleanupMicroshiftVolume(Runtime, "microshift-data"); err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)
}
