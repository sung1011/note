# git reference 参考

## Setup and Config

### git

### config

### help

```bash
git help -w --web #
```

### bugreport

## Getting and Creating Projects

### init

### clone

```bash
git clone --depth 10 < repo > # 深度.保留最新的10个commit, 更前的commit嫁接(grafted)成一个整体
```

## Basic Snapshotting

### add

### status

### diff

```bash
git diff HEAD~3 # HEAD^^^
git diff --cached # 暂存区 与 HEAD 比较
git diff origin/< branch > # 与远端比较 (大多同 --cached)
git diff < commit1 > < commit2 >
git diff < branch1 > < branch2 > -- < file >
git diff --stat ':!<file1>' ':!<file2>' # diff, 但排除 file1 和 file2 ...
git diff --word-diff=plain # 一行内显示diff
```

### commit

```bash
git commit --amend # 修改最近一次提交的msg
git commit -m # --message < msg >
```

### notes

### restore

### reset

```bash
git reset --soft  # reset only HEAD
git reset --mixed # reset HEAD and index    *default
git reset --hard  # reset HEAD, index and working tree
```

### rm

### mv

```bash
git mv a b # 一般重命名大小写时用. 另外可通过配置使大小写敏感`git config core.ignorecase false`
```

## Branching and Merging

### branch

```bash
git branch -vv # 展示HEAD, 分支, oid, message
git branch -d # 删除分支
git branch -a --no-merged # 未合入当前分支的(远端)分支
git branch --merged # 获取已合入当前分支的分支
git branch --merged < branch > # 已合入<branch>的分支
git branch -m < new-branch-name > # 分支重命名
git branch --contains < commit-id > # 列出包含指定commit的分支
```

### checkout

```bash
git checkout -b < new branch > < start_point > # 基于当前分支or某commit 来新建分支
git checkout -- < filename > # 丢弃工作区指定文件的修改
git checkout . # 丢弃工作区当前文件夹的 modified
git checkout < oid > # 检出某次commit; 新建分支(gco -b)来保存修改后的内容(分离头指针 detached HEAD).
git checkout < oid > -- < filename > # 检出指定oid 的 指定文件
git checkout --orphan < new branch > # new unparented branch; 新建无parented分支
git checkout stash@{0} # 检出stash0的快照
```

### switch

### merge

```bash
git merge < branch >

       A---B---C *topic
      /
 D---E---F---G *master

       A---B---C *topic
      /         \
 D---E---F---G---H *master


git merge --squash < branch > # 创建一个单独的提交而不是做一次合并

       A---B---C *topic
      /
 D---E---F---G master

 D---E---F---G---H(ABC) *master
 ```

### mergetool

### log

```bash
git log --oneline
git log --oneline --decorate # 一行 id+msg
git log -< num > # -n< num > 最近n条
git log --all # 所有分支
git log --graph
git log feature ^master # feature里有, master里没有的commmit
git log -m -p < commit-id > # 显示merge的内容; -p可替换为--name-only / --name-status
```

### stash

```bash
git stash -u    # 保存一个stash 包含untracked文件
git stash save  # 保存一个stash
git stash push  # 暂存一个stash
git stash pop < n >  # 弹出一个stash, 不保留记录; n 为 stash@{n}的值; 如 git stash pop 2
git stash apply < n > # 弹出一个stash, 并且保留记录; n 为 stash@{n}的值
git stash show
git stash branch
git stash clear # 删除所有stash
git stash list
git stash list -p
git stash drop < n >

# 工作区有modified的文件时进行pull, 同文件会报错; 可以先stash-push + pull + stash-pop, 此时相同line会冲突
```

### tag

```bash
git tag # 查看标签
git tag -ln # 标签详情
git tag < tag-name > # 创建标签
git tag -d # 删除本地标签
```

### worktree

## Sharing and Updating Projects

### fetch

### pull

```bash
git pull = git fetch && git merge

       A---B---C *feature
      /
 D---E---F---G *master

# E is origin/master in your repository

       A---B---C *feature
      /         \
 D---E---F---G---H *master

# H commit message is "Merge branch 'feature' of < rep >"

[topic] get pull --rebase origin master # fetch && rebase FETCH_HEAD

      A---B---C *topic
     /
D---E---F---G *master

              'A'--'B'--'C' *topic
             /
D---E---F---G *master
```


### push

```bash
git push -u origin < branch > # 关联分支. 当前与远端
git push origin --delete < branch > # 删除远端分支
git push -f # 强制推送 执行前需保证本地是最新(别人没再新的提交)
```


### remote

```bash
git remote add origin < remote-url > # 创建远程仓库
git remote set-url origin < remote-url > # 修改远程仓库
git remote show origin # 远端与本地分支的关系; 远端分支列表 tracked已追踪的 / stale陈旧3month以上
```

### submodule

## Inspection and Comparison

### show

### [log](#log)

### [diff](#diff)

### difftool

### range-diff

### shortlog

### describe

## Patching

### apply

### cherry-pick

### diff

### rebase

```bash
[topic] git rebase < 上游主分支 > < 指定分支 >

[topic] git rebase master # 变基并改变(移动)topic的commit, 到master HEAD的后面
[topic] git rebase -i HEAD~20
[topic] git rebase -i master # 变基并交互式改变被移动的commit

# [master] git checkout -b topic; // 检出功能分支
# [master] git commit t1; // t1
# [topic] git commit t2; // t1 t2
# [master] git commit m1 // t1 t2 m1
# [topic] git rebase master; // 变基到 master HEAD
# [master] git merge topic; // m1 t1 t2 (m1的tree, parent(顺序)变了, 但Date不会变)

      A---B---C *topic
     /
D---E---F---G master

              'A'--'B'--'C' *topic
             /
D---E---F---G master

# 只能变基未push到远端的commit (topic), 否则需要push -f

# rebase后, A'B'C'是时间连续的多个commit, 便于查看 (顺序在G后面, 时间保持原状)

# 不要在master进行rebase操作, 即以topic为基点变基master的commit, 如master不接受push -f, master变基前的commit不会消失.

```

### revert

```bash
git revert < oid > # 提交一个与指定commit内容相反的commit.
git revert -n < oid > # 内容相反的, 但不提交

# 若在主分支revert一个功能分支(revert merge commit id), 则该功能分支无法重新merge到主分支, 需要用cherry-pick.
```


## Debugging

### bisect

### blame

```bash
git blame -L 10,20 < file > # 按行范围进行blame
git blame -b -w < file > # 显示全文blame. -b show oid; -w ignore whitespace
g blame -L 14,14 < file >  | awk '{print $1}' | xargs git show # 显示某行的提交log
```

### grep

## Guides

### gitattributes

### Command-line interface conventions

### Everyday Git

### Frequently Asked Questions (FAQ)

### Glossary

### githooks

### gitignore

### gitmodules

### Revisions

### Submodules

### Tutorial

### Workflows

## Email

### am

### apply :Email

### format-patch

### send-email

### request-pull

## External Systems

### svn

### fast-import

### Administration

### clean

```bash
git clean -id # 交互询问删不删Untracked; -d 和目录
git clean -nd # -n 试图删除Untracked; -d 和目录
git clean -df # -f 直接删除Untracked文件; -d 和目录
```

### gc

### fsck

### reflog

### filter-branch 重写分支

```bash
git filter-branch --force --prune-empty --index-filter 'git rm -rf --cached --ignore-unmatch < file >' --tag-name-filter cat -- --all # 彻底删除某文件
```

### instaweb

### archive

### bundle

## Server Admin

### daemon

### update-server-info

## Plumbing Commands

### cat-file 调试对象信息

```bash
git cat-file -t  # 查看对象类型
git cat-file -s  # 查看对象size
git cat-file -p  # 查看对象内容
```

> [对象类型](git-internals.md#对象类型)


### check-ignore

### commit-tree

### diff-index

### for-each-ref

### hash-object

### ls-files

```bash
git ls-files -m # 列出Modified文件
git ls-files -o # 列出Untracked文件
git ls-files -d # 列出删除的文件
```

### merge-base

### read-tree

### rev-list

```bash
git rev-list --objects --all # 获取所有对象(commit, tree, blob) 及blob对应的文件, tree对应的目录 (commit 和 tree快照对应的数据第二列显示null);
git rev-list --objects < oid(tree) > # 获取快照中所有内容(oid, file)
git rev-list < oid1 >...< oid2 > # 两次提交之间的所有提交
```

### rev-parse

```bash
git rev-parse HEAD^ # 获取上一个commit-id
git rev-parse --short HEAD^ # 获取上一个commit-id (short)
```

### show-ref

### symbolic-ref

### update-index

### update-ref

### verify-pack 读取归档文件(idx)

```bash
git verify-pack -v .git/objects/pack/pack-*.idx # 获取所有pack中的对象详细信息; commit对应的基础tree不会显示
git verify-pack -v .git/objects/pack/pack-*.idx | sort -k 3 -g -r | head -n5 # 获取最大的5个对象

# -v 返回: SHA-1 | type | size | size-in-packfile | offset-in-packfile  

# -v 未分类的对象返回: SHA-1 | type | size | size-in-packfile | offset-in-packfile | depth | base-SHA-1
```

### write-tree

## Others

### gitk

### 标记

```bash
HEAD 头指针
detached HEAD 分离头指针 # git checkout <commit-id>, 即 直接检出obj tree, 而非分支时
^   父
^^^ 父父父
~3  父父父
--  指定文件
```

## ref

<https://www.atlassian.com/git/tutorials/saving-changes/gitignore>