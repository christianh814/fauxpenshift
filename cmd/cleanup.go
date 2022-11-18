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
	"github.com/christianh814/fauxpenshift/pkg/container"
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
		log.Info("Setting runtime to ", Runtime)
		log.Info("Cleaning up any instances")

		// Stop the instance
		container.StopMicroshiftContainer(Runtime, "fauxpenshift")

		//cleanup volume
		container.CleanupMicroshiftVolume(Runtime, "microshift-data")
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
