# git internals 内部原理

## 分区

- `Workspace 工作区` 项目文件夹

- `Index / Stage 暂存区` git add 后; 将内容存入版本库中的暂存区

- `Repository 版本库(本地仓库)` git commit 后; .git文件夹

- `Remote 远程仓库` 云端

## 状态

![img](res/git-status-flow.png)

- `未跟踪 Untracked` 新建文件; 对Unmodified进行rm

- `已修改 Modified` 对 Unmodified 文件 git add 后

- `已暂存 Staged` 对新建文件 git add 后

- `未修改 Unmodified` commit后

- `已提交 Committed` 同Unmodified

## 对象类型

### commit (or tag)


```js

    每个commit记录了提交一个更新的所有元数据, 如指向的tree, 父commit, 作者、提交者、提交日期、提交日志等.每次提交都指向一个tree对象, 记录了当次提交时的目录信息.一个commit可以有多个(至少一个)父commits.


    # git cat-file < oid or tag >
    tree 8c14e08655a16ff642e8bc340aed6abcc24118d9  
    parent 7c7f1be90f15c921fbc1de65431a5245285cfd88  
    author sunji <sung1011@gmail.com> 1600173852 +0800  
    committer sunji <sung1011@gmail.com> 1600173852 +0800  

    "add phpinfo" # message
```

### tree


```js
    每个tree代表了一个目录的信息, 包含了此目录下的blobs, 子目录(对应于子trees), 文件名、路径等元数据.因此, 对于有子目录的目录, git相当于存储了嵌套的trees.


    # git cat-file -p 8c14e08655a16ff642e8bc340aed6abcc24118d9
    040000 tree 4e28030f8c8691fd473e70c2df510df53640b733   .vscode
    040000 tree 8b12d6694b2da38afdd6cdb1507e12b77f972729   logs
    100644 blob ad24f56f441535daa342b4be1e36b6a510c2b000   README.md
    100644 blob 33b4f711f4fce3f6547b6db3919d98273cb8e692   phpinfo.php
```

### blob

  每个blob代表一个(版本的)文件, blob只包含文件的数据, 而忽略文件的其他元数据, 如名字、路径、格式等.

```php
# git cat-file -p 33b4f711f4fce3f6547b6db3919d98273cb8e692
<?php
echo phpinfo();
```

## ref

<https://www.cnblogs.com/qdhxhz/p/9757390.html>