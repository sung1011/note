# git

## 架构

Workspace：工作区  
Index / Stage：暂存区  
Repository：仓库区（或本地仓库）  
Remote：远程仓库  

## 对象类型

blob

- 每个blob代表一个（版本的）文件，blob只包含文件的数据，而忽略文件的其他元数据，如名字、路径、格式等。

tree

- 每个tree代表了一个目录的信息，包含了此目录下的blobs，子目录（对应于子trees），文件名、路径等元数据。因此，对于有子目录的目录，git相当于存储了嵌套的trees。

commit

- 每个commit记录了提交一个更新的所有元数据，如指向的tree，父commit，作者、提交者、提交日期、提交日志等。每次提交都指向一个tree对象，记录了当次提交时的目录信息。一个commit可以有多个（至少一个）父commits。

tag

- tag用于给某个上述类型的对象指配一个便于开发者记忆的名字, 通常用于某次commit。

## 概念

### HEAD    头指针

### ^   父

### ^^^ 父父父

### ~3  父父父

### --  指定文件

### 分离头指针 HEAD detached

基于一个commit的操作(非分支，非tag)  
可创建commit, branch  
修改后最终若不创建branch，将被回收丢弃  

## 常用命令 cmd

### catfile 调试对象信息

- -t    查看对象类型
- -s    查看对象size
- -p    查看对象内容

### help

- -w --web

### mv  移动文件

- 一般重命名大小写时用。 # 另外可通过配置使大小写敏感`git config core.ignorecase false`

### log

- --oneline 一行
- -< num >, -n < num >
- --all
- --graph

### diff

- HEAD~3, - HEAD^^^
- < commit1 > < commit2 >
- --cached HEAD与暂存区比较

### branch

- -d, -D

### commit

- --amend
- -m, --message < msg >

### checkout

- -b
- -- < filename > 丢弃工作区该文件的修改
- < hash > < filename > 检出指定commitid 的 指定文件

### reset

- --soft    reset only HEAD
- --mixed   reset HEAD and index    *default
- --hard    reset HEAD, index and working tree

### rebase

注意: **不要在主分支操作**

- -i, --interactive 交互rebase:
  - p, pick   use commit
  - r, reword use commit, but edit the commit message
  - e, edit   use commit, but stop for amending
  - s, squash use commit, but meld into previous commit
  - f, fixup  like "squash", but discard this commit's log message
  - x, exec   run command (the rest of the line) using shell
  - d, drop   remove commit

- 变基 `git rebase master`
- 变基并修改历史 `git rebase -i master`

```git
1. [master] git commit 12 // 12
2. [master] git checkout -b feature; // 检出功能分支
3. [master] git commit 34; // 1234
4. [feature] git commit ab; // 12ab
5. [feature] git rebase master; // 变基到master HEAD之后 // 1234ab
6. [master] git merge feature; // 1234ab
```

### revert

提交一个与指定commit内容相反的commit。

> 若在主分支revert一个功能分支，则该功能分支无法重新merge到主分支，需要用cherry-pick。

### stash - Stash the changes in a dirty working directory away

- apply 弹出一个stash，并且保留记录
- pop   弹出一个stash，不保留记录
- show
- branch
- clear
- list
- drop

## gitk

git图形界面工具

## .gitignore 文件

```bash
*.a             表示忽略所有 .a 结尾的文件
!lib.a          表示但lib.a除外
/TODO           表示仅仅忽略项目根目录下的 TODO 文件，不包括 subdir/TODO
build/          表示忽略 build/目录下的所有文件，过滤整个build文件夹；
doc/*.txt       表示会忽略doc/notes.txt但不包括 doc/server/arch.txt

bin/:           表示忽略当前路径下的bin文件夹，该文件夹下的所有内容都会被忽略，不忽略 bin 文件
/bin:           表示忽略根目录下的bin文件
/*.c:           表示忽略cat.c，不忽略 build/cat.c
debug/*.obj:    表示忽略debug/io.obj，不忽略 debug/common/io.obj和tools/debug/io.obj
**/foo:         表示忽略/foo,a/foo,a/b/foo等
a/**/b:         表示忽略a/b, a/x/b,a/x/y/b等
!/bin/run.sh    表示不忽略bin目录下的run.sh文件
*.log:          表示忽略所有 .log 文件
config.php:     表示忽略当前路径的 config.php 文件

/mtk/           表示过滤整个文件夹
*.zip           表示过滤所有.zip文件
/mtk/do.c       表示过滤某个具体文件
fd1/*           忽略目录 fd1 下的全部内容；注意，不管是根目录下的 /fd1/ 目录，还是某个子目录 /child/fd1/ 目录，都会被忽略；
```

## 实战

### 迁移

1. `git clone --bare git://github.com/username/project.git` 克隆裸库(仅代码)
2. `git push --mirror git@gitcafe.com/username/newproject.git` 推送到新地址

### 回滚

回滚指定版本 git checkout; 以新建分支回滚 (临时回滚)

- `git checkout {commit_id} && git checkout -b {new_branch_name}`

回滚指定版本、n个版本 git reset --hard; 以主分支回滚 (永久回滚)

- `git reset --hard [^回退上一版本|^^回退上两个版本|~n回退上n个版本|commit_id回退到某一版本] && git push --force`
