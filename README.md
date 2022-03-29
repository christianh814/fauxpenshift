# FauxpenShift

This cli utility creates a Kubernetes cluster using [µShift](https://microshift.io/docs/home/) (Known as "Microshift". A version of OpenShift that runs in a container) and a custom version of the [OpenShift Router](https://github.com/openshift/router) on top of it. This is useful for when you want to test your applications using OpenShift, but [CRC](https://developers.redhat.com/products/codeready-containers/overview) is too heavy.

Check out the usage guide for your platfrom/runtime for specifics:

* [Linux and Podman](docs/podmanLinux.md)
* [Linux and Docker](docs/dockerLinux.md)
* [Mac and Podman](docs/podmanMac.md)
* [Mac and Docker](docs/dockerMac.md)

# Prerequisites

At a minimum

* Docker or Podman
* Access to Nip.io