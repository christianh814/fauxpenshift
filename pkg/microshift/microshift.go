package microshift

import (
	"time"

	"github.com/christianh814/fauxpenshift/pkg/container"
	"github.com/christianh814/fauxpenshift/pkg/olm"
	"github.com/christianh814/fauxpenshift/pkg/router"
	"github.com/christianh814/fauxpenshift/pkg/utils"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/clientcmd"
)

// for now set these defaults
var Runtime string
var MicroShiftImage string = "quay.io/microshift/microshift-aio:latest"
var WaitTime time.Duration = 600 * time.Second

func StartMicroshift(kubeconfigfile string, runtime string) error {
	// Set the variables for the OpenShift Router
	ns := "openshift-ingress"
	depl := "router-default"
	Runtime = runtime

	// try and run the microshift container, return if there's an error
	log.Info("Running Microshift")
	if err := container.RunMicroShiftContainer(Runtime, MicroShiftImage); err != nil {
		return err
	}

	// Copy the Kubeconfig file because we'll need it
	// TODO: It takes a while for µShift to come up. Need a better method of waiting
	log.Info("Setting up Kubeconfig File")
	time.Sleep(10 * time.Second)
	if err := container.CopyKubeConfig(Runtime, "fauxpenshift", kubeconfigfile); err != nil {
		return err
	}

	// Wait until router deployment is there. First create a client so you can pass that over
	client, err := utils.NewClient(kubeconfigfile)
	if err != nil {
		return err
	}

	// Wait until the deployment appears and is ready
	log.Info("Waiting for Router Deployment to appear")
	if err = utils.WaitForDeployment(client, ns, depl, WaitTime); err != nil {
		return err
	}

	// patch router with our custom settings to make it work locally
	ev := []corev1.EnvVar{
		{
			Name:  "ROUTER_SUBDOMAIN",
			Value: "${name}-${namespace}.apps.127.0.0.1.nip.io",
		},
		{
			Name:  "ROUTER_ALLOW_WILDCARD_ROUTES",
			Value: "true",
		},
		{
			Name:  "ROUTER_OVERRIDE_HOSTNAME",
			Value: "true",
		},
	}
	log.Info("Updating Router Configuration")
	if err := router.PatchRouter(ev, client, ns, depl); err != nil {
		return err
	}

	// Wait until the deployment is ready
	log.Info("Waiting for updated Router to rollout")
	if err = utils.WaitForDeployment(client, ns, depl, WaitTime); err != nil {
		return err
	}

	// Install OLM we need the restconfig for "dossa" funciton
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigfile)
	if err != nil {
		return err
	}
	log.Info("Installing OLM from https://operatorhub.io/")
	if err = olm.InstallOLM(restConfig); err != nil {
		return err
	}

	// Wait until the olm-operator deployment is ready
	log.Info("Waiting the OLM Operator to rollout")
	if err = utils.WaitForDeployment(client, "olm", "olm-operator", WaitTime); err != nil {
		return err
	}

	// Wait until the catalog-operator deployment is ready
	log.Info("Waiting the Catalog Operator to rollout")
	if err = utils.WaitForDeployment(client, "olm", "catalog-operator", WaitTime); err != nil {
		return err
	}

	// if we're here we should be okay
	return nil
}
