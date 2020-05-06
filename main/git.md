# git

## 架构

Workspace：工作区  
Index / Stage：暂存区  
Repository：仓库区（或本地仓库）  
Remote：远程仓库  

## 对象类型

blob

> 每个blob代表一个（版本的）文件，blob只包含文件的数据，而忽略文件的其他元数据，如名字、路径、格式等。

tree

> 每个tree代表了一个目录的信息，包含了此目录下的blobs，子目录（对应于子trees），文件名、路径等元数据。因此，对于有子目录的目录，git相当于存储了嵌套的trees。

commit

> 每个commit记录了提交一个更新的所有元数据，如指向的tree，父commit，作者、提交者、提交日期、提交日志等。每次提交都指向一个tree对象，记录了当次提交时的目录信息。一个commit可以有多个（至少一个）父commits。

tag

> tag用于给某个上述类型的对象指配一个便于开发者记忆的名字, 通常用于某次commit。

## 标记

```bash
HEAD 头指针
^   父
^^^ 父父父
~3  父父父
--  指定文件
```

## 常用命令 cmd

### catfile 调试对象信息

```bash
git catfile -t  # 查看对象类型
git catfile -s  # 查看对象size
git catfile -p  # 查看对象内容
```

### help

```bash
git help -w --web #
```

### clone

```bash
git clone --depth 10 # 深度。保留最新的n个commit，更前的commit嫁接(grafted)成一个整体
```

### blame

```bash
git blame -b -w # 显示全文blame。 -b show commitID; -w ignore whitespace
```

### mv  移动文件

```bash
git mv a b # 一般重命名大小写时用。 另外可通过配置使大小写敏感`git config core.ignorecase false`
```

### log

```bash
git log --oneline
git log -< num >, -n < num >
git log --all
git log --graph
```

### diff

```bash
git diff HEAD~3, - HEAD^^^
git diff < commit1 > < commit2 >
git diff --cached # HEAD与暂存区比较
```

### branch

```bash
git branch -d -D
git branch -v 展示HEAD，分支，commitID，message
git branch --no-merged 获取未合入当前分支的分支
git branch --merged 获取已合入当前分支的分支
```

### commit

```bash
git commit --amend 替换上一次提交的msg
git commit -m, --message < msg >
```

### checkout

```bash
git checkout -b # 基于当前分支新建分支
git checkout -- < filename > # 丢弃工作区指定文件的修改
git checkout . # 丢弃工作区当前文件夹的修改
git checkout < commitID > # 检出某次commit。修改后新建分支来保存修改内容（分离头指针）。
git checkout < commitID > < filename > # 检出指定commitid 的 指定文件
```

### reset

```bash
git reset --soft  # reset only HEAD
git reset --mixed # reset HEAD and index    *default
git reset --hard  # reset HEAD, index and working tree
```

### merge

```bash
       A---B---C topic
      /
 D---E---F---G master

       A---B---C topic
      /         \
 D---E---F---G---H master
```

> 如何回滚?

### rebase

```bash
# 1. [master] git commit 12 // 12
# 2. [master] git checkout -b topic; // 检出功能分支
# 3. [master] git commit 34; // 1234
# 4. [topic] git commit ab; // 12ab
# 5. [topic] git rebase master; // 变基到master HEAD之后 // 1234ab
# 6. [master] git merge topic; // 1234ab

# [topic] git rebase master
# [topic] git rebase master topic
# [topic] git rebase -i master // 变基并改变被移动的commit

      A---B---C *topic
     /
D---E---F---G master

              'A'--'B'--'C' *topic
             /
D---E---F---G master
```

> 不要在master上变基已经push到远端的commit, 因为需要push -f

### pull

```bash
# git pull = git fetch && git merge
       A---B---C master on origin
      /
 D---E---F---G *master

> E is origin/master in your repository

       A---B---C origin/master
      /         \
 D---E---F---G---H *master

> H is origin/master in your repository && commit message is "Merge branch 'master' of < rep >"
```

### pull -r --rebase

```bash
# git pull --rebase = git fetch && git rebase FETCH_HEAD

# [topic] get pull --rebase origin master
      A---B---C *topic
     /
D---E---F---G master

              'A'--'B'--'C' *topic
             /
D---E---F---G master

# [master] get pull --rebase origin topic
      A---B---C topic
     /
D---E---F---G *master

D---E---A---B---C---'F'---'G' *master
```

> 如何回滚?

### revert

```bash
git revert < commitID > # 提交一个与指定commit内容相反的commit。
```

> 若在主分支revert一个功能分支，则该功能分支无法重新merge到主分支，需要用cherry-pick。

### stash

```bash
git stash apply 弹出一个stash，并且保留记录
git stash pop   弹出一个stash，不保留记录
git stash push  暂存一个stash
git stash show
git stash branch
git stash clear
git stash list
git stash drop
```

## gitk

## .gitignore 文件

```gitignore
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

```bash
git clone --bare git://github.com/username/project.git # 克隆裸库(仅代码)
git push --mirror git@gitcafe.com/username/newproject.git # 推送到新地址
```

### 回滚

回滚指定版本 git checkout; 以新建分支回滚 (临时回滚)

- `git checkout {commit_id} && git checkout -b {new_branch_name}`

回滚指定版本、n个版本 git reset --hard; 以主分支回滚 (永久回滚)

- `git reset --hard [^回退上一版本|^^回退上两个版本|~n回退上n个版本|commit_id回退到某一版本] && git push --force`

### 当前分支

current_branch=`git rev-parse --abbrev-ref HEAD 2> /dev/null`

### git push -f 找回

TODO
