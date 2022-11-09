# General Usage

This is a general usage guide. If you are using the `docker` runtime, prefix every command with `FAUXPENSHIFT_SET_RUNTIME` 

Example:

```shell
sudo FAUXPENSHIFT_SET_RUNTIME=docker fauxpenshift create
```

You can also set this env variable globally if you wish.

# Create a Cluster

Create a Kubernetes cluster with an OpenShift Router:

```shell
sudo fauxpenshift create
```

Extract your `kubeconfig` file.

```shell
sudo fauxpenshift kubeconfig > fauxpenshift.kubeconfig
```

Export the `KUBECONFIG` env var

```shell
export KUBECONFIG=fauxpenshift.kubeconfig
```

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
oc expose deployment welcome-php --port=8078 --target-port=8080 -n welcome-app
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
sudo fauxpenshift destroy
```

# Troubleshooting

If you find that µShift hasn't started or that it's having trouble starting, you should probably clean up the volume.

> **NOTE** Same commands work for `docker`

First force shutdown the µShift instance

```shell
sudo podman stop fauxpenshift
```

Then delete the data

```shell
sudo podman volume rm microshift-data
```

Then try and start the [Fauxpenshift instance](#general-usage).
