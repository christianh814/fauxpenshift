# Docker with Linux

This will guide you through how to set up your workstation to use
Fauxpenshift with Docker and Linux. This is a one time setup.

# Install Docker

Use [this guide](https://fedoramagazine.org/docker-and-fedora-32/) to install docker.

# Networking/Firewall

The requirements are the [same as µShift](https://microshift.io/docs/getting-started/). First you have to trust the container networking range.

```shell
sudo firewall-cmd --zone=trusted --add-source=10.42.0.0/16 --permanent
```

Next, you will need to allow 80, 443, 6443, and 5353 through your firewall

```shell
sudo firewall-cmd --zone=public --add-port=80/tcp --permanent
sudo firewall-cmd --zone=public --add-port=443/tcp --permanent
sudo firewall-cmd --zone=public --add-port=6443/tcp --permanent
sudo firewall-cmd --zone=public --add-port=5353/udp --permanent
sudo firewall-cmd --reload
```

# SELinux

In order for µShift to run containers, you'll have to allow it via SELinux

```shell
sudo setsebool -P container_manage_cgroup true
```

# Install CLI

Download the CLI from and put it in your path.

```shell
sudo wget -O /usr/local/bin/fauxpenshift https://github.com/christianh814/fauxpenshift/releases/download/v0.0.4/fauxpenshift-amd64-linux
```

Make it executable 

```shell
sudo chmod +x /usr/local/bin/fauxpenshift
```

Shell completion is available, if you wish

```shell
source <(fauxpenshift completion bash)
```

# Ready to use

When reading the usage guide, make sure you export `FAUXPENSHIFT_SET_RUNTIME=docker` before every command (or export that env var globally). Example:

```shell
sudo FAUXPENSHIFT_SET_RUNTIME=docker fauxpenshift create
```

You can now follow the [general usage guide](generalUsage.md)