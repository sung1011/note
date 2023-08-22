# docker build cache

dockerfile中的每一行都会被缓存(layers), 但是当某一行发生变化时, 会导致该行以及后面的所有行都会重新执行

## order layers

将不会发生变化的层放在前面, 会发生变化的层放在后面

```docker
# 未优化, 会导致每次修改代码都要重新下载依赖
FROM golang:1.19.1 # [cached]

COPY . . # [x]
RUN cd src && go mod download # [x]

CMD ["src/bin/server"] # [x]
```

```docker
# 优化, 只有当 go.mod 或 go.sum 发生变化时才会重新下载依赖
FROM golang:1.19.1 # [cached]

COPY src/go.mod src/go.mod # [cached]
RUN cd src && go mod download # [cached]
COPY src src # [x]

CMD ["src/bin/server"] # [x]
```

## keep layers small

### Don't install unnecessary files

COPY 时, 不要包含非必要的文件. 如: log, 包管理器的缓存, 依赖的源码等

利用 `.dockerignore` 忽略不必要的文件

### Use package manager wisely

使用包管理器帮助管理依赖. 如: apt, apk, pip, npm, yarn...

### Use the dedicated RUN cache

使用更细粒度的专用缓存, 以便在更改代码时不必重新下载依赖

```docker
docker run --mount=type=cache,target=/var/cache/apt apt-get update && apt-get install -y git
```

## minimize the number of layers

### use an appropriate minimal base image

基于适当的最小化镜像

### use multi-stage builds

docker会并行执行stage, 但是会阻塞后续有依赖的stage的执行  

```docker
# stage 1
FROM alpine as git
RUN apk add git

# stage 2
FROM git as fetch
WORKDIR /repo
RUN git clone https://github.com/your/repository.git .

# stage 3
FROM nginx as site
COPY --from=fetch /repo/docs/ /usr/share/nginx/html
```

### combine commands into one layer

尽可能将命令组合在一起, 减少层数

```docker
RUN echo "the first command" && \
    echo "the second command"
```

```docker
RUN <<EOF
set -e
echo "the first command"
echo "the second command"
EOF
```

```docker
RUN deploy.sh
```

