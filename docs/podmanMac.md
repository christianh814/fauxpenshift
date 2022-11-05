# Podman on Mac

This will guide you through how to set up your workstation to use
Fauxpenshift with Podman and Mac. This is a one time setup.

# Install Podman

Use `brew` to install Podman

```shell
brew install podman
```

Make sure you're running at LEAST version 4.0.2

```shell
podman --version
```

# Set up Podman

Create (or add) a Podman machine with enough resources

> **NOTE** These values are the minimum. I recommend higher if you can spare it.

```shell
podman machine init --cpus 4 --memory 8192
```

Podman requires to run in `rootful` mode.

```shell
podman machine set --rootful
```

Start your podman machine

```shell
podman machine start
```

You need to allow containers within this machine to run containers inside the continer. First SSH into the machine.

```shell
podman machine ssh
```

Set the following SELinux Boolean

```shell
setsebool -P container_manage_cgroup true
```

You can now exit out of the Podman machine.

```shell
exit
```

# Install the CLI

Download the CLI from and put it in your path.

```shell
sudo wget -O /usr/local/bin/fauxpenshift https://github.com/christianh814/fauxpenshift/releases/download/v0.0.8/fauxpenshift-darwin-arm64
```

Make it executable 

```shell
sudo chmod +x /usr/local/bin/fauxpenshift
```

Shell completion if you wish

> **NOTE** A lot of Mac users use zsh, substitue the shell of your choice

```shell
source <(fauxpenshift completion bash)
```

# Ready to use

You can now follow the [general usage guide](generalUsage.md)
