<div style="padding-bottom: 50px;">
    <img src="./resources/cloner_light.png#gh-dark-mode-only" alt="Cloner" width="200" style="display: block; margin-left: auto; margin-right: auto; padding-top: 50px;">
    <img src="./resources/cloner_dark.png#gh-light-mode-only" alt="Cloner" width="200" style="display: block; margin-left: auto; margin-right: auto; padding-top: 50px;">
</div>

Cloner is a simple Docker container, built on Go, to easily clone a git repo into a Docker Volume on a regular basis. This lets you easily mount a git repo into a container, in a way that will auto-update, for deployments with systems like Docker Compose and Portainer.

<div style="padding-top:50px; text-align: center;">

![GitHub last commit](https://img.shields.io/github/last-commit/Homelabian/cloner?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/homelabian/cloner?style=for-the-badge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/homelabian/cloner?style=for-the-badge)
![GitHub](https://img.shields.io/github/license/homelabian/cloner?style=for-the-badge)

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/homelabian/cloner?style=for-the-badge)

![GitHub language count](https://img.shields.io/github/languages/count/homelabian/cloner?style=for-the-badge)
![GitHub top language](https://img.shields.io/github/languages/top/homelabian/cloner?style=for-the-badge)

</div>

## Pulling the container

You can pull the container like any other Docker Container:

```shell
docker pull ghcr.io/homelabian/cloner
```

## Running the container

You can run the container with very little info, the only requirement is you supply a repo URL to clone.

```shell
docker run -e CLONER_1_REPO=https://github.com/homelabian/cloner -v clonervol:/job-1 ghcr.io/homelabian/cloner
```

This will clone the Cloner repo to your volume `clonervol`, and will clone again everyday at midnight.

## Understanding Jobs

Cloner can run any number of jobs, this is done by numbering the environment variables. Each job can have the following settings:

| Variable | Use |
| --- | --- |
| Cron | The cron string to represent running the clone again. |
| RepoAuth | Enable or disable auth login when cloning |
| Repo | The repo to clone |
| RepoUser | The username to use when using Auth |
| RepoPass | The password to use when using Auth |
| Output | The directory to clone to |

### How a job runs

Each time a job will run through a few steps when cloning:

1. Delete the existing clone
2. Clone the repo to the output folder

For this reason, Cloner doesn't support cloning to the root of a volume, but we are looking for ways to acomplish this.

### Default Job

By default, a job will refresh daily at midnight to an output folder called `/job-<NUM>/output`, and will do so with not authentication. Therefore, you only NEED to provide the Repo you want to clone.

## Environment Variables

The format for a Environment Variable is quite simple:

```
CLONER_<NUM>_<OPTION>=<VALUE>
```

Where `<NUM>` is the Job Number, `<OPTION>` is the Job Variable from the above table, in all caps. Finally, `<VALUE>` is the value you want to set the environment variable to.

## Docker Compose Example

```yml
version: '3.8'

services:
    cloner:
        image: ghcr.io/homelabian/cloner
        environment:
            - CLONER_1_REPO=https://github.com/homelabian/cloner
            - CLONER_1_CRON="0 0 0 * * *"
            - CLONER_1_OUTPUT=/data/cloner
        volumes:
            - data:/data

volumes:
    data:
```

## Contributing

Cloner is open to anyone who wants to expand on it and contribute! Simeply Fork, Change, Commit, PR!

Homelabian is commited to providing a safe open-source space, as such Homelabian will endevour to moderate any contributor environments to meet everyones needs. This may mean we need to update our rules from time to time to refelect issues we are facing in our communities.

TLDR; Be nice, don't make people feel uncomfortable and be supportive in all communications.