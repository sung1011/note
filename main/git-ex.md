# git 实战

## 迁移

```bash
    git clone --bare git://github.com/username/project.git # 克隆裸库(仅代码)
    git push --mirror git@gitcafe.com/username/newproject.git # 推送到新地址
```

## 回滚

1. 回滚指定版本 git checkout; 以新建分支回滚 (临时回滚)

   ```bash
      git checkout {commit_id} && git checkout -b {new_branch_name}
   ```

2. 回滚指定版本、n个版本 git reset --hard; 以主分支回滚 (永久回滚)

   ```bash
   git reset --hard [^回退上一版本|^^回退上两个版本|~n回退上n个版本|commit_id回退到某一版本] && git push -f
   ```

## 当前分支

```bash
current_branch=`git rev-parse --abbrev-ref HEAD 2> /dev/null`
```

## 错误的分支merge

1. 错误的将feature合入master, 未push

   1. 版本回退 `[master] git reset --hard origin/master`

2. 错误的将feature合入master, 并push

   1. 找到merge产生的commitID
   2. 撤销提交 `[master] git revert <merge commit> -m 1` // 产生revert的commitID
   3. 若需要取消上述撤销 `[master] git revert <revert commit> -m 1` 或 将feature的内容逐个cherry-pick到master `[master] git cherry-pick <feature commit>`

3. 错误的将带有feature的dev合入到master, 并push

   1. master撤销dev的所有内容 `[master] git revert <merge commit> -m 1`
   2. master保留feature内容(但不保留dev的a) `[master] git checkout <feature> -- <X files>; git add .;git commit` -- master已正常
   3. master合入dev(将revert带回dev) `[dev] git merge master` -- 此时dev中的a内容没有了, 期望dev有a
   4. 检出dev被撤销的文件(还原出a内容) `[dev] git checkout <merge commit> -- <X files>; git add .; git commit` -- dev已正常
   5. 若dev中有其他feature，需要类似【4】把这些被撤销feature内容还原出来

   ```bash
   # 正常情况下 dev 和 master 为平行关系，feature合入dev进行测试，合入master进行上线
    D-----E---X---F---a----- dev # 比master多一些脏提交a
             /                 \
            X feature           \  # 错误的将带有feature(X)的dev合入master
           /                     \
    D-----E---F------------------Xa master # feature错误的合入master 并且 dev的a错误的合进了master
   ```

### 查找大文件

```bash
# 获取最大的5个blob
git verify-pack -v .git/objects/pack/pack-*.idx | sort -k 3 -g -r | head -n5
# 通过blob获取文件名
git rev-list --objects --all | grep 33e65fc9089735578c9fd232637dde7d95db5d22
```

## 彻底删除某个文件（大文件、涉密文件）

```bash
git filter-branch --index-filter 'git rm --cached --ignore-unmatch < file >'

rm -rf .git/refs/original
git reflog expire --expire=now --all
git fsck --full --unreachable
git repack -A -d
git gc --aggressive --prune=now
git push --force
```

## git push -f 找回

TODO

### .gitignore 文件

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

## [git note](git.md)
