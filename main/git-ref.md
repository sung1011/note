# 1. git reference 参考

- [1. git reference 参考](#1-git-reference-参考)
  - [1.1. Setup and Config](#11-setup-and-config)
    - [1.1.1. git](#111-git)
    - [1.1.2. config](#112-config)
    - [1.1.3. help](#113-help)
    - [1.1.4. bugreport](#114-bugreport)
  - [1.2. Getting and Creating Projects](#12-getting-and-creating-projects)
    - [1.2.1. init](#121-init)
    - [1.2.2. clone](#122-clone)
  - [1.3. Basic Snapshotting](#13-basic-snapshotting)
    - [1.3.1. add](#131-add)
    - [1.3.2. status](#132-status)
    - [1.3.3. diff](#133-diff)
    - [1.3.4. commit](#134-commit)
    - [1.3.5. notes](#135-notes)
    - [1.3.6. restore](#136-restore)
    - [1.3.7. reset](#137-reset)
    - [1.3.8. rm](#138-rm)
    - [1.3.9. mv](#139-mv)
  - [1.4. Branching and Merging](#14-branching-and-merging)
    - [1.4.1. branch](#141-branch)
    - [1.4.2. checkout](#142-checkout)
    - [1.4.3. switch](#143-switch)
    - [1.4.4. merge](#144-merge)
    - [1.4.5. mergetool](#145-mergetool)
    - [1.4.6. log](#146-log)
    - [1.4.7. stash](#147-stash)
    - [1.4.8. tag](#148-tag)
    - [1.4.9. worktree](#149-worktree)
  - [1.5. Sharing and Updating Projects](#15-sharing-and-updating-projects)
    - [1.5.1. fetch](#151-fetch)
    - [1.5.2. pull](#152-pull)
    - [1.5.3. push](#153-push)
    - [1.5.4. remote](#154-remote)
    - [1.5.5. submodule](#155-submodule)
  - [1.6. Inspection and Comparison](#16-inspection-and-comparison)
    - [1.6.1. show](#161-show)
    - [1.6.2. log](#162-log)
    - [1.6.3. diff](#163-diff)
    - [1.6.4. difftool](#164-difftool)
    - [1.6.5. range-diff](#165-range-diff)
    - [1.6.6. shortlog](#166-shortlog)
    - [1.6.7. describe](#167-describe)
  - [1.7. Patching](#17-patching)
    - [1.7.1. apply](#171-apply)
    - [1.7.2. cherry-pick](#172-cherry-pick)
    - [1.7.3. diff](#173-diff)
    - [1.7.4. rebase](#174-rebase)
    - [1.7.5. revert](#175-revert)
  - [1.8. Debugging](#18-debugging)
    - [1.8.1. bisect](#181-bisect)
    - [1.8.2. blame](#182-blame)
    - [1.8.3. grep](#183-grep)
  - [1.9. Guides](#19-guides)
    - [1.9.1. gitattributes](#191-gitattributes)
    - [1.9.2. Command-line interface conventions](#192-command-line-interface-conventions)
    - [1.9.3. Everyday Git](#193-everyday-git)
    - [1.9.4. Frequently Asked Questions (FAQ)](#194-frequently-asked-questions-faq)
    - [1.9.5. Glossary](#195-glossary)
    - [1.9.6. githooks](#196-githooks)
    - [1.9.7. gitignore](#197-gitignore)
    - [1.9.8. gitmodules](#198-gitmodules)
    - [1.9.9. Revisions](#199-revisions)
    - [1.9.10. Submodules](#1910-submodules)
    - [1.9.11. Tutorial](#1911-tutorial)
    - [1.9.12. Workflows](#1912-workflows)
  - [1.10. Email](#110-email)
    - [1.10.1. am](#1101-am)
    - [1.10.2. apply :Email](#1102-apply-email)
    - [1.10.3. format-patch](#1103-format-patch)
    - [1.10.4. send-email](#1104-send-email)
    - [1.10.5. request-pull](#1105-request-pull)
  - [1.11. External Systems](#111-external-systems)
    - [1.11.1. svn](#1111-svn)
    - [1.11.2. fast-import](#1112-fast-import)
    - [1.11.3. Administration](#1113-administration)
    - [1.11.4. clean](#1114-clean)
    - [1.11.5. gc](#1115-gc)
    - [1.11.6. fsck](#1116-fsck)
    - [1.11.7. reflog](#1117-reflog)
    - [1.11.8. filter-branch 重写分支](#1118-filter-branch-重写分支)
    - [1.11.9. instaweb](#1119-instaweb)
    - [1.11.10. archive](#11110-archive)
    - [1.11.11. bundle](#11111-bundle)
  - [1.12. Server Admin](#112-server-admin)
    - [1.12.1. daemon](#1121-daemon)
    - [1.12.2. update-server-info](#1122-update-server-info)
  - [1.13. Plumbing Commands](#113-plumbing-commands)
    - [1.13.1. cat-file 调试对象信息](#1131-cat-file-调试对象信息)
    - [1.13.2. check-ignore](#1132-check-ignore)
    - [1.13.3. commit-tree](#1133-commit-tree)
    - [1.13.4. diff-index](#1134-diff-index)
    - [1.13.5. for-each-ref](#1135-for-each-ref)
    - [1.13.6. hash-object](#1136-hash-object)
    - [1.13.7. ls-files](#1137-ls-files)
    - [1.13.8. merge-base](#1138-merge-base)
    - [1.13.9. read-tree](#1139-read-tree)
    - [1.13.10. rev-list](#11310-rev-list)
    - [1.13.11. rev-parse](#11311-rev-parse)
    - [1.13.12. show-ref](#11312-show-ref)
    - [1.13.13. symbolic-ref](#11313-symbolic-ref)
    - [1.13.14. update-index](#11314-update-index)
    - [1.13.15. update-ref](#11315-update-ref)
    - [1.13.16. verify-pack 读取归档文件(idx)](#11316-verify-pack-读取归档文件idx)
    - [1.13.17. write-tree](#11317-write-tree)
  - [1.14. Others](#114-others)
    - [1.14.1. gitk](#1141-gitk)
    - [1.14.2. 标记](#1142-标记)
  - [1.15. ref](#115-ref)

## 1.1. Setup and Config

### 1.1.1. git

### 1.1.2. config

### 1.1.3. help

```js
git help -w --web #
```

### 1.1.4. bugreport

## 1.2. Getting and Creating Projects

### 1.2.1. init

### 1.2.2. clone

```js
git clone --depth 10 <repo> // 深度.保留最新的10个commit, 更前的commit嫁接(grafted)成一个整体
git clone --bare <repo> // 克隆裸仓库
```

## 1.3. Basic Snapshotting

### 1.3.1. add

### 1.3.2. status

### 1.3.3. diff

```js
git diff HEAD~3 // HEAD^^^
git diff --cached // 暂存区 与 HEAD 比较
git diff origin/<branch> // 与远端比较 (大多同 --cached)
git diff <commit1> <commit2>
git diff <branch1> <branch2> -- <file>
git diff --stat ':!<file1>' ':!<file2>' // diff, 但排除 file1 和 file2 ...
git diff --word-diff=plain // 一行内显示diff
```

### 1.3.4. commit

```js
git commit --amend // 修改最近一次提交的msg
git commit -m // --message <msg>
```

### 1.3.5. notes

### 1.3.6. restore

### 1.3.7. reset

```js
git reset --soft  // reset only HEAD
git reset --mixed // reset HEAD and index    *default
git reset --hard  // reset HEAD, index and working tree
```

### 1.3.8. rm

### 1.3.9. mv

```js
git mv a b // 一般重命名大小写时用. 另外可通过配置使大小写敏感`git config core.ignorecase false`
```

## 1.4. Branching and Merging

### 1.4.1. branch

```js
git branch -vv // 展示HEAD, 分支, oid, message
git branch -d // 删除分支
git branch -a --no-merged // 未合入当前分支的(远端)分支
git branch --merged // 获取已合入当前分支的分支
git branch --merged <branch> // 已合入<branch>的分支
git branch -m <new-branch-name> // 分支重命名
git branch --contains <commit-id> // 列出包含指定commit的分支
```

### 1.4.2. checkout

```js
git checkout -b <new branch> <start_point> // 基于当前分支or某commit 来新建分支
git checkout -- <filename> // 丢弃工作区指定文件的修改
git checkout . // 丢弃工作区当前文件夹的 modified
git checkout <oid> // 检出某次commit; 新建分支(gco -b)来保存修改后的内容(分离头指针 detached HEAD).
git checkout <oid> -- <filename> // 检出指定oid 的 指定文件
git checkout --orphan <new branch> // new unparented branch; 新建无parented分支
git checkout stash@{0} // 检出stash0的快照
```

### 1.4.3. switch

### 1.4.4. merge

```js
// 原始状态
       A---B---C *topic
      /
 D---E---F---G *master

       A---B---C *topic
      /         \
 D---E---F---G---H *master

git merge --no-ff <branch> // 创建一个merge的commit

git merge --ff <branch> // 合并, 但不创建commit, 仿佛没有topic分支

git merge --ff-only <branch> // 合并, 但不创建commit, 仿佛没有topic分支; 必须当前是远端的最新

 D---E---A---F---B---G---C *master

git merge --squash <branch> // 创建一个单独的提交而不是做一次合并

 D---E---F---G---H(ABC) *master
 ```

### 1.4.5. mergetool

### 1.4.6. log

```js
git log --oneline
git log --oneline --decorate // 一行 id+msg
git log -<num> // -n<num> 最近n条
git log --all // 所有分支
git log --graph
git log feature ^master // feature里有, master里没有的commmit
git log -m -p <commit-id> // 显示merge的内容; -p可替换为--name-only / --name-status
```

### 1.4.7. stash

```js
git stash -u    // 保存一个stash 包含untracked文件
git stash save  // 保存一个stash
git stash push  // 暂存一个stash
git stash pop <n>  // 弹出一个stash, 不保留记录; n 为 stash@<n>的值; 如 git stash pop 2
git stash apply <n> // 弹出一个stash, 并且保留记录; n 为 stash@<n>的值
git stash show
git stash branch
git stash clear // 删除所有stash
git stash list
git stash list -p
git stash drop <n>

# 工作区有modified的文件时进行pull, 同文件会报错; 可以先stash-push + pull + stash-pop, 此时相同line会冲突
```

### 1.4.8. tag

```js
git tag // 查看标签列表
git tag -a <tag-name> -m <msg> // 创建带msg的标签
git tag <tag-name> // 创建标签(无msg)
git tag -d // 删除本地标签
git tag -ln // 标签详情

git describe --tags `git rev-list --tags --max-count=1` // 查看最新一个tag
git describe --abbrev=0 --tags // 查看当前分支最新一个tag
git ls-remote --tags // 查看远端标签
git push origin --delete <tag-name> // 删除远端标签
git push origin <tag-name> // 推送指定标签到远端
git push origin --tags // 推送所有标签到远端
git fetch --prune --prune-tags // 同步远端tag到本地, 删除本地不存在于远端tag; 简写 git fetch -p -P
```

### 1.4.9. worktree

```js
git worktree add ../xxx.worktrees // 上级目录新增xxx.worktrees文件夹, 内容为当前分支的克隆, 新增xxx.worktrees分支; 与当前共享git
git worktree add ../xxx.worktrees -b xyz // 上级目录新增xxx.worktrees文件夹, 内容为当前分支的克隆, 新增xyz分支; 与当前共享git
git worktree list
git worktree lock
git worktree move
git worktree prune
git worktree remove
git worktree repair
git worktree unlock
```

## 1.5. Sharing and Updating Projects

### 1.5.1. fetch

```js
git fetch --prune --prune-tags // 同步远端tag到本地, 删除本地不存在于远端tag; 简写 git fetch -p -P
```

### 1.5.2. pull

```js
git pull = git fetch && git merge

       A---B---C *feature
      /
 D---E---F---G *master

# E is origin/master in your repository

       A---B---C *feature
      /         \
 D---E---F---G---H *master

# H commit message is "Merge branch 'feature' of <rep>"

[topic] get pull --rebase origin master // fetch && rebase FETCH_HEAD

      A---B---C *topic
     /
D---E---F---G *master

              'A'--'B'--'C' *topic
             /
D---E---F---G *master
```

### 1.5.3. push

```js
git push -u origin <branch> // 关联分支. 当前与远端
git push origin --delete <branch> // 删除远端分支
git push -f // 强制推送 执行前需保证本地是最新(别人没再新的提交)
git push --tags // 推送附带本地所有tag
git push origin master // 不论当前是何分支, 推送本地的master分支到远端
git push --mirror <remote-url> // 推送到新的远端仓库; 迁移
```

### 1.5.4. remote

```js
git remote add origin <remote-url> // 创建远程仓库
git remote set-url origin <remote-url> // 修改远程仓库
git remote show origin // 远端与本地分支的关系; 远端分支列表 tracked已追踪的 / stale陈旧3month以上
```

### 1.5.5. submodule

## 1.6. Inspection and Comparison

### 1.6.1. show

### 1.6.2. log

### 1.6.3. diff

### 1.6.4. difftool

### 1.6.5. range-diff

### 1.6.6. shortlog

### 1.6.7. describe

## 1.7. Patching

### 1.7.1. apply

### 1.7.2. cherry-pick

```js
git cherry-pick <commit-id> // 将指定commit的内容应用到当前分支; commit-id会变
```

### 1.7.3. diff

### 1.7.4. rebase

```js
[topic] git rebase <上游主分支> <指定分支>

[topic] git rebase master // 变基并改变(移动)topic的commit, 到master HEAD的后面
[topic] git rebase -i HEAD~20
[topic] git rebase -i master // 变基并交互式改变被移动的commit

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

### 1.7.5. revert

```js
git revert -m 1 <oid> // 提交一个与指定commit内容相反的commit; -m 1 保持主分支 还原合入的分支; -m 2 还原主分支 保持合入的分支
git revert -n <oid> // 内容相反的, 但不提交
git revert -m 1 <last_1> <last_2> <last_3> // 从后往前回滚, 避免冲突

// 若在主分支revert一个功能分支(revert merge commit id), 则该功能分支无法重新merge到主分支, 需要用cherry-pick.
```


## 1.8. Debugging

### 1.8.1. bisect

```js
git bisect start [终点] [起点] // 开始二分查找
git bisect good // 指定commit是好的, 继续往新前进寻找bug
git bisect bad // 当前commit是坏的, 继续往旧后退寻找bug
git bisect reset // 恢复退出bisect
git bisect log > bisect-log.txt // 查看二分查找的log, 经历哪些commit 这些commit的状态(good bad skip)
git bisect replay bisect-log.txt // 重放二分查找的log, 重新进行二分查找
git bisect run <cmd> // 二分查找时, 自动执行cmd, 根据cmd的返回值判断当前commit是good还是bad
```

### 1.8.2. blame

```js
git blame -L 10,20 <file> // 按行范围进行blame
git blame -b -w <file> // 显示全文blame. -b show oid; -w ignore whitespace
g blame -L 14,14 <file>  | awk '{print $1}' | xargs git show // 显示某行的提交log
```

### 1.8.3. grep

## 1.9. Guides

### 1.9.1. gitattributes

### 1.9.2. Command-line interface conventions

### 1.9.3. Everyday Git

### 1.9.4. Frequently Asked Questions (FAQ)

### 1.9.5. Glossary

### 1.9.6. githooks

### 1.9.7. gitignore

### 1.9.8. gitmodules

### 1.9.9. Revisions

### 1.9.10. Submodules

### 1.9.11. Tutorial

### 1.9.12. Workflows

## 1.10. Email

### 1.10.1. am

### 1.10.2. apply :Email

### 1.10.3. format-patch

### 1.10.4. send-email

### 1.10.5. request-pull

## 1.11. External Systems

### 1.11.1. svn

### 1.11.2. fast-import

### 1.11.3. Administration

### 1.11.4. clean

```js
git clean -id // 交互询问删不删Untracked; -d 和目录
git clean -nd // -n 试图删除Untracked; -d 和目录
git clean -df // -f 直接删除Untracked文件; -d 和目录
```

### 1.11.5. gc

### 1.11.6. fsck

### 1.11.7. reflog

```js
git reflog show --date=iso < branch > // 分支信息, 创建 / 合并 / 变基 / 重置 / 检出 / 提交 / 恢复 / cherry-pick / rebase / pull / push / stash
```

### 1.11.8. filter-branch 重写分支

```js
git filter-branch --force --prune-empty --index-filter 'git rm -rf --cached --ignore-unmatch <file>' --tag-name-filter cat -- --all // 彻底删除某文件
```

### 1.11.9. instaweb

### 1.11.10. archive

### 1.11.11. bundle

## 1.12. Server Admin

### 1.12.1. daemon

### 1.12.2. update-server-info

## 1.13. Plumbing Commands

### 1.13.1. cat-file 调试对象信息

```js
git cat-file -t  // 查看对象类型
git cat-file -s  // 查看对象size
git cat-file -p  // 查看对象内容
```

> [对象类型](git-internals.md#对象类型)


### 1.13.2. check-ignore

### 1.13.3. commit-tree

### 1.13.4. diff-index

```js
git add -A; git diff-index -q HEAD || (git commit -m 'xx' && git push) // 暂存区有变化就提交
```

### 1.13.5. for-each-ref

### 1.13.6. hash-object

### 1.13.7. ls-files

```js
git ls-files -m // 列出Modified文件
git ls-files -o // 列出Untracked文件
git ls-files -d // 列出删除的文件
```

### 1.13.8. merge-base

### 1.13.9. read-tree

### 1.13.10. rev-list

```js
git rev-list --objects --all // 获取所有对象(commit, tree, blob) 及blob对应的文件, tree对应的目录 (commit 和 tree快照对应的数据第二列显示null);
git rev-list --objects <oid(tree)> // 获取快照中所有内容(oid, file)
git rev-list <oid1>...<oid2> // 两次提交之间的所有提交
```

### 1.13.11. rev-parse

```js
git rev-parse HEAD^ // 获取上一个commit-id
git rev-parse --short HEAD^ // 获取上一个commit-id (short)
```

### 1.13.12. show-ref

### 1.13.13. symbolic-ref

### 1.13.14. update-index

### 1.13.15. update-ref

### 1.13.16. verify-pack 读取归档文件(idx)

```js
git verify-pack -v .git/objects/pack/pack-*.idx // 获取所有pack中的对象详细信息; commit对应的基础tree不会显示
git verify-pack -v .git/objects/pack/pack-*.idx | sort -k 3 -g -r | head -n5 // 获取最大的5个对象

# -v 返回: SHA-1 | type | size | size-in-packfile | offset-in-packfile  

# -v 未分类的对象返回: SHA-1 | type | size | size-in-packfile | offset-in-packfile | depth | base-SHA-1
```

### 1.13.17. write-tree

## 1.14. Others

### 1.14.1. gitk

### 1.14.2. 标记

```js
HEAD 头指针
detached HEAD 分离头指针 // git checkout <commit-id>, 即 直接检出obj tree, 而非分支时
^   父
^^^ 父父父
~3  父父父
--  指定文件
```

## 1.15. ref

- <https://www.atlassian.com/git/tutorials/saving-changes/gitignore>