package kind

import (
	"os"
	"time"

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

var KindImageVersion string = "kindest/node:v1.23.3"

// CreateKindCluster creates KIND cluster to use as the temp cluster manager
func CreateKindCluster(name string, cfg string) error {
	//create a new KIND provider
	provider := cluster.NewProvider()
	// This lists clusters ~>
	// https://github.com/kubernetes-sigs/kind/blob/v0.11.1/pkg/cmd/kind/get/clusters/clusters.go#L48
	// provider.List()

	// write the config file out until I find a better way of doing this
	kindconfig := "/tmp/fxskc.yaml"
	f, err := os.Create(kindconfig)
	if err != nil {
		return err
	}
	_, err = f.WriteString(KindConfig)
	if err != nil {
		return err
	}
	f.Close()

	// Create a KIND instance and write out the kubeconfig in the specified location
	err = provider.Create(
		name,
		cluster.CreateWithKubeconfigPath(cfg),
		cluster.CreateWithConfigFile(kindconfig),
		cluster.CreateWithDisplayUsage(false),
		cluster.CreateWithDisplaySalutation(false),
		cluster.CreateWithWaitForReady(30*time.Second),
		cluster.CreateWithNodeImage(KindImageVersion),
	)

	if err != nil {
		return err
	}

	// Remove the file
	os.RemoveAll(kindconfig)

	return nil
}

// DeleteKindCluster deletes KIND cluster based on the name given
func DeleteKindCluster(name string, cfg string) error {
	/* Testing quieting down deleting too*/
	/*
		klogger := kindcmd.NewLogger()
		provider := cluster.NewProvider(
			cluster.ProviderWithLogger(klogger),
		)
	*/

	provider := cluster.NewProvider()

	err := provider.Delete(name, cfg)

	if err != nil {
		return err
	}

	return nil

}

// GetKindKubeconfig returns the Kubeconfig of the named KIND cluster
func GetKindKubeconfig(name string, internal bool) (string, error) {
	// Create a provider and return the named kubeconfig file as a string
	provider := cluster.NewProvider()
	return provider.KubeConfig(name, internal)
}
