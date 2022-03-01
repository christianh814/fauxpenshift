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
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// destroyCmd represents the destroy command
var destroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "destroys a cluster",
	Long: `This will destroy a cluster. There is no way
to "save" your cluster.

The functionality exists within KIND, but not when using 
this tool. PRs are welcome!`,
	Run: func(cmd *cobra.Command, args []string) {
		// For now, let's just use the default K8S config path. Later this can be an option
		homedir, _ := os.UserHomeDir()
		kcfg := homedir + "/.kube/config"

		// Create the Kubernetes Cluste
		log.Info("Destroying Kubernetes Cluster")
		err := kind.DeleteKindCluster("fauxpenshift", kcfg)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(destroyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// destoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// destoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
