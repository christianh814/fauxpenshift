package kind

import (
	"strings"
	"time"

	"github.com/christianh814/fauxpenshift/pkg/utils"
	"sigs.k8s.io/kind/pkg/cluster"
)

var KindConfig string = `kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
networking:
  podSubnet: "10.254.0.0/16"
  serviceSubnet: "172.30.0.0/16"
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    listenAddress: 0.0.0.0
  - containerPort: 443
    hostPort: 443
    listenAddress: 0.0.0.0
`

// Set the default Kind Image version
var KindImageVersion string = "kindest/node:v1.23.3"

// We are using the same kind of provider for this whole package
var Provider *cluster.Provider = cluster.NewProvider(
	utils.GetDefaultRuntime(),
)

// CreateKindCluster creates KIND cluster
func CreateKindCluster(name string, cfg string) error {
	// Create a KIND instance and write out the kubeconfig in the specified location
	err := Provider.Create(
		name,
		cluster.CreateWithKubeconfigPath(cfg),
		cluster.CreateWithRawConfig([]byte(KindConfig)),
		cluster.CreateWithDisplayUsage(false),
		cluster.CreateWithDisplaySalutation(false),
		cluster.CreateWithWaitForReady(30*time.Second),
		cluster.CreateWithNodeImage(KindImageVersion),
	)

	if err != nil {
		return err
	}

	return nil
}

// DeleteKindCluster deletes KIND cluster based on the name given
func DeleteKindCluster(name string, cfg string) error {
	err := Provider.Delete(name, cfg)

	if err != nil {
		return err
	}

	return nil

}

// GetKindKubeconfig returns the Kubeconfig of the named KIND cluster
func GetKindKubeconfig(name string, internal bool) (string, error) {
	return Provider.KubeConfig(name, internal)
}

// GetKindClusters returns returns clusters matching a given name
func GetKindClusters(match string) ([]string, error) {
	// Create an empty []string and get all Kind clusters
	l := []string{}
	clusters, err := Provider.List()

	// Return error if there is any
	if err != nil {
		return nil, err
	}
	// no clusters is actually valid
	if len(clusters) == 0 {
		return nil, nil
	}

	// Loop through each cluster and just filter out what isn't being looked for. Append what is found
	for _, cluster := range clusters {
		if strings.Contains(cluster, match) {
			l = append(l, cluster)
		}
	}
	return l, nil
}
