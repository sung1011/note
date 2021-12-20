# dockerfile

## FROM

## RUN

构建新的Image Layer时执行(docker build).

```dockerfile
RUN ["executable","param1","param2"] # exec格式, 推荐
RUN command param1 param2 # shell格式
```

> 可多个RUN  
> 一般用作安装软件包

## CMD

容器启动后默认执行的命令和参数(docker run 会被覆盖).

```dockerfile
CMD ["executable","param1","param2"] # exec格式, 推荐
CMD command param1 param2 # shell格式
CMD ["param1","param2"] # 用作 ENTRYPOINT 的默认补充参数
```

> docker run 指定了命令, CMD会被忽略  
> 定义多个CMD, 只执行最后一个

## ENTRYPOINT

容器启动时运行的命令(docker run 不会被覆盖).

```dockerfile
ENTRYPOINT  ["executable", "param1", "param2"] # exec格式, 推荐.
ENTRYPOINT command param1 param2 # shell格式；忽略任何 CMD 或 docker run 提供的参数.
```

> 让容器以应用程序或者服务的形式运行  
> 不会被忽略, 一定会执行

## LABEL

## EXPOSE

## ENV

## ADD

## COPY

功能类似ADD, 但是是不会自动解压文件, 也不能访问网络资源

## VOLUME

## USER

## WORKDIR

## ARG

## ONBUILD

## STOPSIGNAL

## HEALTHCHECK

## SHELL

## 实战

### Shell && Exec 运行指令

- Shell方式: `<instruction> <command>` 默认被shell解析

```dockerfile
ENV name world
CMD /bin/echo Hello $name # hello world
```

- Exec方式: `<instruction> ["executable", "param1", "param2", ...]` 可读性好, 默认不会被shell解析.

```dockerfile
ENV name world
CMD [ "/bin/echo", "Hello, $name" # hello $name 不被shell解析
CMD [ "/bin/sh", "-c", "echo", "Hello, $name" # hello world 手写/bin/sh -c才能被shell解析
```

## ref

[docs.docker](https://docs.docker.com/engine/reference/builder/)
