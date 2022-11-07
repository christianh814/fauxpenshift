/*
Copyright © 2022 Christian Hernandez <christian@chernand.io>

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
	"os"

	"github.com/christianh814/fauxpenshift/pkg/container"
	"github.com/christianh814/fauxpenshift/pkg/microshift"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// cleanupCmd represents the cleanup command
var cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Stops the container immediately. And cleans up any artifacts that it may have left behind.",
	Long: `This command will stop the container immediately and clean up any artifacts that it may have left behind.
This command should only be used if you are having issues with creating an instance and you want to start over.

Example:
	fauxpenshift cleanup

This is "scorch the earth", so use with caution.`,
	Run: func(cmd *cobra.Command, args []string) {
		// set runtime
		var rt string = microshift.Runtime
		if os.Getenv("FAUXPENSHIFT_SET_RUNTIME") == "docker" {
			rt = "docker"
		}

		log.Info("Cleaning up any instances")

		// Stop the instance
		container.StopMicroshiftContainer(rt, "fauxpenshift")

		//cleanup volume
		container.CleanupMicroshiftVolume(rt, "microshift-data")
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// cleanupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// cleanupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
