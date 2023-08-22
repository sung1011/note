# docker CLI

## Subcommands

### `attach`

Attach local standard input, output, and error streams to a running container

```sh
docker attach <container>
```

### `build`

Build an image from a Dockerfile

```sh
```

### `builder`

Manage builds

```sh
docker builder build
docker builder prune -af # 删除构建时产生的缓存
```

### `checkpoint`

Manage checkpoints

```sh
```

### `commit`

Create a new image from a container’s changes

```sh
```

> not recommended

### `config`

Manage Swarm configs

```sh
```

### `container`

Manage containers

```sh
```

### `context`

Manage contexts

```sh
```

### `cp`

Copy files/folders between a container and the local filesystem

```sh
docker cp <container>:<path> <path> # 从容器拷贝到本地
docker cp <path> <container>:<path> # 从本地拷贝到容器
```

### `create`

Create a new container

```sh
```

### `diff`

Inspect changes to files or directories on a container’s filesystem

```sh
```

### `events`

Get real time events from the server

```sh
```

### `exec`

Execute a command in a running container

```sh
```

### `export`

Export a container’s filesystem as a tar archive

```sh
```

### `history`

Show the history of an image

```sh
```

### `image`

Manage images

```sh
```

### `images`

List images

```sh
```

### `import`

Import the contents from a tarball to create a filesystem image

```sh
```

### `info`

Display system-wide information

```sh
```

### `inspect`

Return low-level information on Docker objects

```sh
```

### `kill`

Kill one or more running containers

```sh
```

### `load`

Load an image from a tar archive or STDIN

```sh
docker load -i ngx.tar
```

> save

### `login`

Log in to a registry

```sh
```

### `logout`

Log out from a registry

```sh
```

### `logs`

Fetch the logs of a container

```sh
```

### `manifest`

Manage Docker image manifests and manifest lists

```sh
```

### `network`

Manage networks

```sh
```

### `node`

Manage Swarm nodes

```sh
```

### `pause`

Pause all processes within one or more containers

```sh
```

### `plugin`

Manage plugins

```sh
```

### `port`

List port mappings or a specific mapping for the container

```sh
```

### `ps`

List containers

```sh
```

### `pull`

Download an image from a registry

```sh
```

### `push`

Upload an image to a registry

```sh
```

### `rename`

Rename a container

```sh
```

### `restart`

Restart one or more containers

```sh
```

### `rm`

Remove one or more containers

```sh
```

### `rmi`

Remove one or more images

```sh
```

### `run`

Create and run a new container from an image

```sh
# docker run -it --rm --name <name> -p <port>:<port> -v <volume>:<volume> <image> <command>
docker run -it --rm --name my-nginx -p 8080:80 -v /Users/zhengkai/Downloads:/usr/share/nginx/html nginx # it: 交互模式, name: 容器起名, p: 端口映射, v: 挂载卷, rm: 退出后删除容器
```

### `save`

Save one or more images to a tar archive (streamed to STDOUT by default)

```sh
docker save ngx-app:latest -o ngx.tar
```

> load

### `search`

Search Docker Hub for images

```sh
```

### `secret`

Manage Swarm secrets

```sh
```

### `service`

Manage Swarm services

```sh
```

### `stack`

Manage Swarm stacks

```sh
```

### `start`

Start one or more stopped containers

```sh
```

### `stats`

Display a live stream of container(s) resource usage statistics

```sh
```

### `stop`

Stop one or more running containers

```sh
```

### `swarm`

Manage Swarm

```sh
```

### `system`

Manage Docker

```sh
```

### `tag`

Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE

```sh
```

### `top`

Display the running processes of a container

```sh
```

### `trust`

Manage trust on Docker images

```sh
```

### `unpause`

Unpause all processes within one or more containers

```sh
```

### `update`

Update configuration of one or more containers

```sh
```

### `version`

Show the Docker version information

```sh
```

### `volume`

Manage volumes

```sh
```

### `wait`

Block until one or more containers stop, then print their exit codes

```sh
```

## Environment variables

### `DOCKER_API_VERSION`

Override the negotiated API version to use for debugging (e.g. 1.19)

```sh
```

### `DOCKER_CERT_PATH`

Location of your authentication keys. This variable is used both by the docker CLI and the dockerd daemon

```sh
```

### `DOCKER_CONFIG`

The location of your client configuration files

```sh
```

### `DOCKER_CONTENT_TRUST_SERVER`

The URL of the Notary server to use. Defaults to the same URL as the registry

```sh
```

### `DOCKER_CONTENT_TRUST`

When set Docker uses notary to sign and verify images. Equates to --disable-content-trust=false for build, create, pull, push, run

```sh
```

### `DOCKER_CONTEXT`

Name of the docker context to use (overrides DOCKER_HOST env var and default context set with docker context use)

```sh
```

### `DOCKER_DEFAULT_PLATFORM`

Default platform for commands that take the --platform flag

```sh
```

### `DOCKER_HIDE_LEGACY_COMMANDS`

When set, Docker hides “legacy” top-level commands (such as docker rm, and docker pull) in docker help output, and only Management commands per object-type (e.g., docker container) are printed. This may become the default in a future release

```sh
```

### `DOCKER_HOST`

Daemon socket to connect to

```sh
```

### `DOCKER_TLS_VERIFY`

When set Docker uses TLS and verifies the remote. This variable is used both by the docker CLI and the dockerd daemon

```sh
```

### `BUILDKIT_PROGRESS`

Set type of progress output (auto, plain, tty) when building with BuildKit backend. Use plain to show container output (default auto)

```sh
```
