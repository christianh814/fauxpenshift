package container

import (
	"os"
	"os/exec"

	log "github.com/sirupsen/logrus"
)

// RunMicroShiftContainer starts microshift it needs a runtime and a container to run
func RunMicroShiftContainer(runtime string, container string) error {
	// Make sure that the runtime is in the path
	if _, err := exec.LookPath(runtime); err != nil {
		return err
	}

	// Try and run the microshift container. For now, this is crude. Comeback and fix later
	if err := exec.Command(
		runtime,
		"run",
		"-d",
		"--rm",
		"--name",
		"fauxpenshift",
		"--privileged",
		"-v",
		"microshift-data:/var/lib",
		"-p",
		"6443:6443",
		"-p",
		"80:80",
		"-p",
		"443:443",
		"--label",
		"fauxpenshift=instance",
		container,
	).Run(); err != nil {
		// TODO: Need to figure out why running a container with go returns 125
		//return err
		log.Warn(err.(*exec.ExitError).Error())
		return nil
	}

	// if we're here we should be okay
	return nil
}

//CopyKubeConfig copies the kubeconfig from the given container using the runtime and copies it over to a destination
func CopyKubeConfig(runtime string, instance string, dest string) error {
	// Let's copy over the
	if err := exec.Command(
		runtime,
		"cp",
		instance+":/var/lib/microshift/resources/kubeadmin/kubeconfig",
		dest,
	).Run(); err != nil {
		//return err
		// TODO: figure out why executing contianers always returns 125
		return nil
	}

	// Let's fix permissions
	if err := exec.Command(
		"chmod",
		"0600",
		dest,
	).Run(); err != nil {
		return err
	}

	// Let's try and fix ownershit
	if err := exec.Command(
		"chown",
		os.Getenv("SUDO_USER")+":"+os.Getenv("SUDO_USER"),
		dest,
	).Run(); err != nil {
		return err
	}

	// if we're here we should be okay
	return nil
}
