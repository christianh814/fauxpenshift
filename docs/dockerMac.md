# Docker with Mac

This will guide you through how to set up your workstation to use
Fauxpenshift with Docker and Mac. This is a one time setup.

# Install Docker Desktop

Follow the instructions to install Docker Desktop. This can be found on the Docker website.

[https://docs.docker.com/desktop/mac/install/](https://docs.docker.com/desktop/mac/install/)

# Set up Docker Desktop

You will need to increse the resouces allowcated to the Docker Desktop. 

> **NOTE** I recommed setting it to 4CPUs and 8GB of RAM

![docker-pref](https://d33wubrfki0l68.cloudfront.net/23353e57f81ecdd1485b2fb6db9489d2f635fd1e/1ad25/docs/user/images/docker-pref-2.png)

# Install the CLI

Download the CLI from and put it in your path.

> *NOTE* M1 version only

```shell
sudo wget -O /usr/local/bin/fauxpenshift https://github.com/christianh814/fauxpenshift/releases/download/v0.0.10/fauxpenshift-darwin-arm64
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

When reading the usage guide, make sure you export `FAUXPENSHIFT_RUNTIME=docker` before every command (or export that env var globally). Example:

```shell
sudo FAUXPENSHIFT_RUNTIME=docker fauxpenshift create
```

> NOTE: You can also use the `--runtime` option

You can now follow the [general usage guide](generalUsage.md)
