package router

import (
	"context"
	"time"

	"github.com/christianh814/fauxpenshift/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
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

func PatchRouter(envars []corev1.EnvVar, client kubernetes.Interface, ns string, deployment string) error {
	// get the named deployment
	rd, err := client.AppsV1().Deployments(ns).Get(context.TODO(), deployment, v1.GetOptions{})
	if err != nil {
		return err
	}

	// Add the env vars
	envars = append(envars, rd.Spec.Template.Spec.Containers[0].Env...)
	rd.Spec.Template.Spec.Containers[0].Env = envars

	// Update the deployment
	if _, err = client.AppsV1().Deployments(ns).Update(context.TODO(), rd, v1.UpdateOptions{
		FieldManager: "fauxpenshift",
	}); err != nil {
		return err
	}

	//if we are here we're okay
	return nil
}
