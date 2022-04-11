package olm

import (
	"context"
	"time"

	"github.com/christianh814/fauxpenshift/pkg/utils"
	"k8s.io/client-go/rest"
)

var OLMVer string = "v0.20.0"
var CRDURL string = "https://github.com/operator-framework/operator-lifecycle-manager/releases/download/" + OLMVer + "/crds.yaml"
var OLMURL string = "https://github.com/operator-framework/operator-lifecycle-manager/releases/download/" + OLMVer + "/olm.yaml"

// InstallOLM tries to install OLM and returns the error if it can't
func InstallOLM(k *rest.Config) error {
	// We're doing the same to both CRD and OLM urls
	urls := []string{CRDURL, OLMURL}

	for _, url := range urls {
		// First let's try and download the file
		file, err := utils.DownloadFileString(url)
		if err != nil {
			return err
		}

		// Let's turn the YAML string into a [][]byte
		yaml, err := utils.SplitYAML([]byte(file))
		if err != nil {
			return err
		}

		// Let's loop over these and try to apply it to our cluster
		for _, y := range yaml {
			if err := utils.DoSSA(context.TODO(), k, y); err != nil {
				return err
			}
			// TODO: figure out what I'm loading and wait until it's available.
			time.Sleep(3 * time.Second)
		}

	}

	// if we're here we should be okay
	return nil
}
