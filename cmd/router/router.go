package router

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/christianh814/fauxpenshift/cmd/utils"
	"k8s.io/client-go/tools/clientcmd"
)

// InstallRouter installs the OpenShift router on the given cluster
func InstallRouter(kcfg string) error {
	// Create a TMP dir
	routerTmpDir := "/tmp/fxs-router"
	err := os.MkdirAll(routerTmpDir, 0775)
	if err != nil {
		return err
	}

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
	for i, s := range files {
		// write each config into a file
		tf := routerTmpDir + "/" + fmt.Sprint(i) + ".yaml"
		f, err := os.Create(tf)
		if err != nil {
			return err
		}
		_, err = f.WriteString(s)
		if err != nil {
			return err
		}
		f.Close()

		// do server side apply to it
		kRestClient, err := clientcmd.BuildConfigFromFlags("", kcfg)
		if err != nil {
			return err
		}
		err = utils.DoSSA(context.TODO(), kRestClient, tf)
		if err != nil {
			return err
		}
		// HACK come back and fix later
		time.Sleep(time.Second)
	}

	// Remote the TMP dir
	err = os.RemoveAll(routerTmpDir)
	if err != nil {
		return err
	}

	// If we're here, we must be okay
	return nil
}
