# FauxpenShift

This cli utility creates a Kubernetes cluster using [KIND](kind.sigs.k8s.io) and installs the [OpenShift Router](https://github.com/openshift/router) on top of it. This is useful for when you want to test your applications using OpenShift routes, but CRC is too heavy.

# Prerequisites

At a minimum

* Docker
* Access to Nip.io

> **NOTE:** This might work with Podman, but not if you're doing "rootless" mode (i.e. Just run this as root)

While you don't need the `kind` CLI, you do need to satisfy all the prereqs for KIND. If you're having trouble see [their official docs](https://kind.sigs.k8s.io/).

# Running it

Download the CLI from HERE and put it in your path.


```shell
fauxpenshift create
```

> **NOTE** If you want to be brave and run `podman` then do: `KIND_EXPERIMENTAL_PROVIDER=podman fauxpenshift create`

You should have a Kubernetes Cluster with the router running

```shell
oc get pods -n openshift-ingress 
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

Patch things that the `oc expose` didn't 100% get you.

```shell
kubectl patch route welcome-php -n welcome-app --type=json -p='[{"op": "add", "path": "/spec/to/kind", "value":"Service"}]'
kubectl patch route welcome-php -n welcome-app --type=json -p='[{"op": "add", "path": "/spec/wildcardPolicy", "value":"Subdomain"}]'
```

Get your route

```shell
oc get route -n welcome-app
```

Curl it (or open it up in a browser)

```shell
curl -sI http://$(oc get route welcome-php -n welcome-app -o jsonpath='{.status.ingress[0].host}')
```

