# FauxpenShift

This cli utility creates a Kubernetes cluster using [µShift](https://microshift.io/docs/home/) (Known as "Microshift". A version of OpenShift that runs in a container) and a custom version of the [OpenShift Router](https://github.com/openshift/router) on top of it. This is useful for when you want to test your applications using OpenShift, but [CRC](https://developers.redhat.com/products/codeready-containers/overview) is too heavy.

# Prerequisites

At a minimum

* Docker or Podman (Docker is experimental, but should work)
* Access to Nip.io

You will also have to satisfy all the prerequisites of µShift. You will need to pay attention to the [Appliation Development](https://microshift.io/docs/getting-started/#using-microshift-for-application-development) SELinux setup and the [Firewall Rules](https://microshift.io/docs/getting-started/#deploying-microshift) setup as well.

> **NOTE** Don't actually run Microshift, just do the prereqs

# Running it

Download the CLI from and put it in your path.

## Linux

```shell
sudo wget -O /usr/local/bin/fauxpenshift https://github.com/christianh814/fauxpenshift/releases/download/v0.0.3/fauxpenshift-amd64-linux
```

## Mac OS (Intel)

```shell
sudo wget -O /usr/local/bin/fauxpenshift https://github.com/christianh814/fauxpenshift/releases/download/v0.0.3/fauxpenshift-amd64-darwin
```

Make it executable 

```shell
sudo chmod +x /usr/local/bin/fauxpenshift
```

Bash completion if you wish

```shell
source <(fauxpenshift completion bash)
```

Create a Kubernetes cluster with an OpenShift Router:

```shell
fauxpenshift create
```

> **NOTE** To use Docekr run: `sudo  FAUXPENSHIFT_SET_RUNTIME=docker fauxpenshift create`

You should have a Kubernetes Cluster with the router running

```shell
oc get pods -A
```

# Testing It

Now let's create an app and expose a route. First create a namespace

```shell
oc create ns welcome-app
```

Create a deployment in this namespace

```shell
oc create deployment welcome-php \
--image=quay.io/redhatworkshops/welcome-php:latest -n welcome-app
```

Create a service for this deployment

```shell
oc expose deployment welcome-php --port=8080 --target-port=8080 -n welcome-app
```

Now create a route

```shell
oc expose svc/welcome-php -n welcome-app
```

Get your route

```shell
oc get route -n welcome-app
```

Curl it (or open it up in a browser)

```shell
curl -sI http://$(oc get route welcome-php -n welcome-app -o jsonpath='{.status.ingress[0].host}')
```

# Clean Up

Delete your cluster

```shell
fauxpenshift destroy
```

> **NOTE** If using Docekr, run:  `sudo FAUXPENSHIFT_SET_RUNTIME=docker fauxpenshift destroy`
