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
	"os"

	"github.com/christianh814/fauxpenshift/cmd/kind"
	"github.com/christianh814/fauxpenshift/cmd/router"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a cluster",
	Long: `Create a local Kubernetes cluster and installs the
OpenShift router on this cluster.

Since the router binds to 80/443, you must run this as root.
Rootless won't work because of the aforementioned binding.
PRs are welcome!`,
	Run: func(cmd *cobra.Command, args []string) {
		// This program should be run as root as "rootless" containers probably won't work
		if os.Getuid() != 0 {
			log.Warn("Running as non-root can produce some problems. YMMV.")
		}

		// For now, let's just use the default K8S config path. Later this can be an option
		homedir, _ := os.UserHomeDir()
		kcfg := homedir + "/.kube/config"

		// Create the Kubernetes Cluster
		log.Info("Creating Kubernetes Cluster")
		err := kind.CreateKindCluster("fauxpenshift", kcfg)
		if err != nil {
			log.Fatal(err)
		}

		// Install router
		log.Info("Installing OpenShift Router")
		err = router.InstallRouter(kcfg)
		if err != nil {
			log.Fatal(err)
		}

	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
