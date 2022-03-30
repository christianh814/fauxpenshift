package container

import (
	"fmt"
	"os"
	"os/exec"
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

//CopyKubeConfig cats the kubeconfig from the given container using the runtime and copies it over to a destination
func CopyKubeConfig(runtime string, instance string, dest string) error {
	// Get the kubeconfig file as a []byte. Check for error
	kcf, err := DisplayMicroshiftKubeconfig(runtime, instance)
	if err != nil {
		return err
	}

	// write that to the file and check for error
	err = os.WriteFile(dest, kcf, 0555)
	if err != nil {
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
		if err.(*exec.ExitError).ExitCode() != 125 {
			return nil, err
		}
	}

	// return the output
	return cmdOutPut, nil
}

//DisplayMicroshiftKubeconfig shows the kubeconfig file based specified instance name
func DisplayMicroshiftKubeconfig(runtime string, instance string) ([]byte, error) {
	// Get the container based on the name given
	cmdOutPut, err := exec.Command(
		runtime,
		"exec",
		"-i",
		instance,
		"cat",
		"/var/lib/microshift/resources/kubeadmin/kubeconfig",
	).Output()

	// check of errors
	if err != nil {
		if err.(*exec.ExitError).ExitCode() != 125 {
			return nil, err
		}
	}

	// return the output
	return cmdOutPut, nil
}

//StopMicroshiftKubeconfig shows the kubeconfig file based specified instance name
func StopMicroshiftContainer(runtime string, instance string) error {
	// Stop container based on the given name
	if err := exec.Command(
		runtime,
		"stop",
		instance,
	).Run(); err != nil {
		if err.(*exec.ExitError).ExitCode() != 125 {
			return nil
		}
	}

	// if we're here it's probably okay...right?
	return nil
}

func CleanupMicroshiftVolume(runtime string, volume string) error {
	// Cleanup any volumes that may have been created
	// TODO: This is ugly but good for now. It's hardcoded to microshift-data
	if err := exec.Command(
		runtime,
		"volume",
		"rm",
		volume,
	).Run(); err != nil {
		/*
			if err.(*exec.ExitError).ExitCode() != 125 {
				return nil
			}
		*/
		return err
	}

	// if we're here it's probably okay...right?
	return nil

}
