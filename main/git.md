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
git cat-file -t  # 查看对象类型
git cat-file -s  # 查看对象size
git cat-file -p  # 查看对象内容
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
git blame -b -w < file > # 显示全文blame。 -b show oid; -w ignore whitespace
git blame -L 10,20 < file > # 按行范围进行blame
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
git diff --cached # 暂存区 与 HEAD 比较
```

### branch

```bash
git branch -vv # 展示HEAD，分支，oid，message
git branch -d
git branch --no-merged # 获取未合入当前分支的分支
git branch --merged # 获取已合入当前分支的分支
git branch --merged master # 获取已合入master的分支
git branch -m < new-branch-name > # 分支重命名
```

### commit

```bash
git commit --amend # 替换上一次提交的msg
git commit -m, --message < msg >
```

### push

```bash
git push -u origin < branch > # 关联分支。 当前 与 <branch>
git push origin --delete < branch > # 删除远端分支
```

### checkout

```bash
git checkout -b # 基于当前分支新建分支
git checkout -- < filename > # 丢弃工作区指定文件的修改
git checkout . # 丢弃工作区当前文件夹的 modified
git checkout < oid > # 检出某次commit。修改后新建分支来保存修改内容（分离头指针）。
git checkout < oid > -- < filename > # 检出指定oid 的 指定文件
git checkout --orphan < new branch > # 新建0提交的分支，当前内容全部转为committed状态
```

### reset

```bash
git reset --soft  # reset only HEAD
git reset --mixed # reset HEAD and index    *default
git reset --hard  # reset HEAD, index and working tree
```

### merge

```bash
git merge < branch >
git merge --squash < branch > 创建一个单独的提交而不是做一次合并

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

git rebase < 上游主分支 > < 分支 >

[topic] git rebase master // 变基并自动改变被移动的commit
[topic] git rebase -i HEAD~20
[topic] git rebase -i master // 变基并交互式改变被移动的commit

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

[topic] get pull --rebase origin master
      A---B---C *topic
     /
D---E---F---G master

              'A'--'B'--'C' *topic
             /
D---E---F---G master

```

> 如何回滚?  
> 不要在master进行rebase操作，即以topic为基点变基master的commit，由于master不接受push -f，master变基前的commit不会消失。

### revert

```bash
git revert < oid > # 提交一个与指定commit内容相反的commit。
git revert -n < oid > # 内容相反的，但不提交
```

> 若在主分支revert一个功能分支，则该功能分支无法重新merge到主分支，需要用cherry-pick。

### stash

```bash
git stash -u    # 保存一个stash 包含untracked文件
git stash save  # 保存一个stash
git stash apply < stash@{n} > # 弹出一个stash，并且保留记录
git stash pop   # 弹出一个stash，不保留记录
git stash push  # 暂存一个stash
git stash show
git stash branch
git stash clear # 删除所有stash
git stash list
git stash drop
```

### ls-files

```bash
git ls-files -m # 列出modified文件
git ls-files -o # 列出Untracked文件
git ls-files -d # 列出删除的文件
```

### clean

```bash
git clean -id # 交互询问删不删Untracked -d 和目录
git clean -nd # -n 试图删除Untracked -d 和目录
git clean -df # -f 直接删除Untracked文件; -d 和目录
```

### bundle

TODO

### remote

```bash
git remote add origin < remote-url > # 创建远程仓库
git remote set-url origin < remote-url > # 修改远程仓库
git remote show origin # 远端与本地分支的关系
```

### tag

```bash
git tag # 查看标签
git tag -ln # 标签详情
git tag < tag-name > # 创建标签
git tag -d # 删除本地标签
```

### describe

TODO

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

- `git reset --hard [^回退上一版本|^^回退上两个版本|~n回退上n个版本|commit_id回退到某一版本] && git push -f`

### 当前分支

current_branch=`git rev-parse --abbrev-ref HEAD 2> /dev/null`

### git push -f 找回

TODO

### 错误的分支merge

1. 错误的将feature合入master, 未push

   1. 版本回退 `[master] git reset --hard origin/master`

2. 错误的将feature合入master, 并push

   1. 找到merge产生的commitID
   2. 撤销提交`[master] git revert <merge commit> -m 1` // 产生revert的commitID
   3. 若需要取消上述撤销 `[master] git revert <revert commit> -m 1` 或 将feature的内容逐个cherry-pick到master `[master] git cherry-pick <feature commit>`

3. 错误的将带有feature的dev合入到master, 并push

   1. 撤销合并 `[master] git revert <merge commit> -m 1`
   2. 检出feature修改(master要保留)的文件`[master] git checkout <feature> -- <X files>; git add .;git commit` -- master去除X, master已正常
   3. master合入dev `[dev] git merge master` -- 将revert带回dev，此时dev中的a被撤销了
   4. 检出dev被撤销的文件 `[dev] git checkout <merge commit> -- <X files>; git add .; git commit` -- 还原出a, dev已正常

```bash
# 正常情况下 dev 和 master 为平行关系，feature合入dev进行测试，合入master进行上线
 D---E--a-F--X-G---H---I---J dev # 比master多一些脏提交a
            /                 \
           X feature           \  # 错误的将带有feature(X)的dev合入master
          /                     \
 D---E---F---G---H---------------abXY master # feature错误的合入master 并且 dev的a错误的合进了master
```

## ref

[git book](https://git-scm.com/book/zh/v2/)
[git tips](https://github.com/521xueweihan/git-tips)
