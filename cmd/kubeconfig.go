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

	log "github.com/sirupsen/logrus"

	"github.com/christianh814/fauxpenshift/pkg/container"
	"github.com/spf13/cobra"
)

// kubeconfigCmd represents the kubeconfig command
var kubeconfigCmd = &cobra.Command{
	Use:   "kubeconfig",
	Short: "Extracts the kubeconfig from the Kubernetes cluster",
	Long: `This extracts the kubeconfig from the Kubernetes cluster
so that you can write it out to a different place.

Useful since the cluster was created using SUDO.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Get the Kubeconfig
		kc, err := container.DisplayMicroshiftKubeconfig(Runtime, "fauxpenshift")

		// check for errors
		if err != nil {
			log.Fatal(err)
		}

		// Display the kubeconfig
		fmt.Println(string(kc))
	},
}

func init() {
	rootCmd.AddCommand(kubeconfigCmd)
}
