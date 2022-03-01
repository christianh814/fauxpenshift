package router

import (
	"context"
	"time"

	"github.com/christianh814/fauxpenshift/cmd/utils"
	"k8s.io/client-go/tools/clientcmd"
)

// InstallRouter installs the OpenShift router on the given cluster
func InstallRouter(kcfg string) error {
	// We're doing the same for all configurations
	files := []string{
		ClusterRole,
		ClusterRoleBindingSA,
		ClusterRoleBindingDelegator,
		Namespace,
		ServiceAccount,
		RouterCrd,
		RouterDeploy,
	}

	// Create files out of these and apply it to the kubernetes cluster that was created
	for _, s := range files {
		// do server side apply to it
		kRestClient, err := clientcmd.BuildConfigFromFlags("", kcfg)
		if err != nil {
			return err
		}
		err = utils.DoSSA(context.TODO(), kRestClient, []byte(s))
		if err != nil {
			return err
		}
		// HACK come back and fix later
		time.Sleep(time.Second)
	}

	// If we're here, we must be okay
	return nil
}
