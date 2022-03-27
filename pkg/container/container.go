package container

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/christianh814/fauxpenshift/pkg/utils"
	//log "github.com/sirupsen/logrus"
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
		if err.(*exec.ExitError).ExitCode() != 125 {
			return err
		}
		//log.Warn(err.(*exec.ExitError).ExitCode())
		//return nil
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
		if err.(*exec.ExitError).ExitCode() != 125 {
			return err
		}
		//log.Warn(err.(*exec.ExitError).ExitCode())
		//return nil
	}

	// Let's fix permissions
	if err := exec.Command(
		"chmod",
		"0600",
		dest,
	).Run(); err != nil {
		return err
	}

	// TODO: For Mac, the user doesn't get it's own group so I have to guess based on the destination
	var owner string = utils.User + ":" + utils.User
	if strings.Contains(dest, "Users") {
		owner = utils.User
	}

	// Let's try and fix ownership
	if err := exec.Command(
		"chown",
		owner,
		dest,
	).Run(); err != nil {
		return err
	}

	// if we're here we should be okay
	return nil
}

//DisplayMicroshiftInstance lists containers based on the specified label
func DisplayMicroshiftInstance(runtime string, label string) ([]byte, error) {
	// Get the container based on the label given
	cmdOutPut, err := exec.Command(
		runtime,
		"ps",
		"--filter",
		label,
		"--format",
		fmt.Sprintf(`table {{.Names}}\t{{.Image}}\t{{.Status}}`),
	).Output()

	// check of errors
	if err != nil {
		return nil, err
	}

	// return the output
	return cmdOutPut, nil
}
