package container

import (
	"fmt"
	"os"
	"os/exec"
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
		"--net",
		"host",
		"--label",
		"fauxpenshift=instance",
		container,
	).Run(); err != nil {
		// TODO: Need to figure out why trying to run a container returns a 125 even though it's running
		if err.(*exec.ExitError).ExitCode() != 125 {
			return err
		}
		fmt.Println(err)
		fmt.Println(err.(*exec.ExitError).ExitCode())
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
		return err
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
		os.Getenv("SUDO_USER")+"."+os.Getenv("SUDO_USER"),
		dest,
	).Run(); err != nil {
		return err
	}

	// if we're here we should be okay
	return nil
}
