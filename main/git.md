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
### 分离头指针 HEAD detached
基于一个commit的操作(非分支，非tag)  
可创建commit, branch  
修改后最终若不创建branch，将被回收丢弃  

### --
指定文件

## cmd
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
- HEAD~3
- HEAD^^^
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

### reset
- --soft    reset only HEAD
- --mixed   reset HEAD and index    *default
- --hard    reset HEAD, index and working tree

### rebase
- -i, --interactive 交互rebase:
  - p, pick   use commit
  - r, reword use commit, but edit the commit message
  - e, edit   use commit, but stop for amending
  - s, squash use commit, but meld into previous commit
  - f, fixup  like "squash", but discard this commit's log message
  - x, exec   run command (the rest of the line) using shell
  - d, drop   remove commit

## gitk
git图形界面工具

## 版本撤销
revert
- 提交一个与指定commit内容相反的commit
- 若在主分支revert一个功能分支， 则该功能分支无法重新merge到主分支， 可以用cherry-pick。

---

## 版本回溯
回溯指定版本 git checkout
- `git checkout {commit_id} && git checkout -b {new_branch_name}`

回溯指定版本、回溯n个版本 git reset --hard
- `git reset --hard [^回退上一版本|^^回退上两个版本|~n回退上n个版本|commit_id回退到某一版本] && git push --force`

---

## 版本整合
rebase
- 在未合并的分支, 合并分支中的多个commit为一个commit。