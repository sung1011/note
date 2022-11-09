# git 实战

## pull的内容涉及正在修改的文件

```bash
    # 工作区如果有 modified 的文件a, pull的内容恰巧包含a文件, 则会提示报错
    # 可以`git stash push` + `git pull` + `git stash pop`, 注意这有可能会冲突


    Updating e72ede4..bb7de2c
    error: Your local changes to the following files would be overwritten by merge:
        ccc
    Please commit your changes or stash them before you merge.
    Aborting
```

## 贡献度

```bash
# commit数量排名
git log --pretty='%aN' | sort | uniq -c | sort -k1 -n -r | head -n 3

# 指定用户的增删代码量
git log --author="{用户名}" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' 

# 所有用户的增删代码量
git log --format='%aN' | sort -u | while read name; do echo -en "$name\t"; git log --author="$name" --pretty=tformat: --numstat | awk '{ add += $1; subs += $2; loc += $1 - $2 } END { printf "added lines: %s, removed lines: %s, total lines: %s\n", add, subs, loc }' -; done

# 附加时间范围 git log --since="2018-03-01" --before="2019-01-09"
```


## 迁移

```bash
# 方法1
git clone --bare git://github.com/username/project.git # 克隆裸库(仅代码)

# 方法2
git push --mirror git@gitcafe.com/username/newproject.git # 推送到新地址
```

## 回滚


```bash
git checkout {commit_id} && git checkout -b {new_branch_name} #临时回滚; 回滚指定版本 && 新建分支

git reset --hard [^回退上一版本|^^回退上两个版本|~n回退上n个版本|commit_id回退到某一版本] && git push -f # 回退最近1/n个commit

git rebase -i HEAD [] && git push -f # 交互模式中指定删除某1/n个commit;
```

> 注意 `reset/ rebase` + `push -f` , 其他人pull后, 若本地有回滚前的内容, 则这些内容不会被删除, 而是在版本库中, 容易被重新带到远端

## 当前分支

```bash
git rev-parse --abbrev-ref HEAD 2> /dev/null
```

## 错误的分支merge

1. 错误的将feature合入master, 未push

   1. 版本回退 `[master] git reset --hard origin/master`

2. 错误的将feature合入master, 并push

   1. 找到merge产生的commitID
   2. 撤销提交 `[master] git revert <merge commit> -m 1` // 产生revert的commitID
   3. 若需要取消上述撤销 `[master] git revert <revert commit> -m 1` 或 将feature的内容逐个cherry-pick到master `[master] git cherry-pick <feature commit>`

3. 错误的将带有dev合入到master, 并push

```bash
# 正常情况下 dev 和 master 为平行关系, feature合入dev进行测试, 合入master进行上线
 D-----E---X---F---a----- dev # 比master多一些脏提交a
          /                 \
         X feature           \  # 错误的将带有feature(X)的dev合入master
        /                     \
 D-----E---F------------------Xa master # feature错误的合入master 并且 dev的a错误的合进了master

1. master撤销dev的所有内容 `[master] git revert <merge commit> -m 1`
2. master保留feature内容(但不保留dev的a) `[master] git checkout <feature> -- <X files>; git add .;git commit` -- master已正常
3. master合入dev(将revert带回dev) `[dev] git merge master` -- 此时dev中的a内容没有了, 期望dev有a
4. 检出dev被撤销的文件(还原出a内容) `[dev] git checkout <merge commit> -- <X files>; git add .; git commit` -- dev已正常
5. 若dev中有其他feature, 需要类似【4】把这些被撤销feature内容还原出来

# 污染严重的话, 可以废弃master 从上游分支重检出一个新的master' (比如main), 然后重新合并feature
```

### 查找大文件

```bash
# 获取最大的5个blob
git verify-pack -v .git/objects/pack/pack-*.idx | sort -k 3 -g -r | head -n5
# 通过blob获取文件名
git rev-list --objects --all | grep < oid(blob) >
```

## 彻底删除某个文件(大文件、涉密文件)

```bash
git filter-branch --force --index-filter \
  "git rm --cached --ignore-unmatch < file >" \
  --prune-empty --tag-name-filter cat -- --all
```

```bash
git filter-branch --index-filter 'git rm --cached --ignore-unmatch < file >'

rm -rf .git/refs/original
git reflog expire --expire=now --all
git fsck --full --unreachable
git repack -A -d
git gc --aggressive --prune=now
git push --force
```

## 哪些分支包含指定commit

```bash
git branch --contains < commitid >
```

## 导出某次commit的文件

```bash
git diff-tree -r --no-commit-id --name-only < oid(commit) > | xargs tar -rf mycommit.tar
```

## 文件创建时间

```bash
git log --pretty=format:"%ad" -- < file > | tail -1
```

## 分支信息 (包含创建时间)

```bash
git reflog show --date=iso < branch >
```

## 某次merge的内容

TODO

## 获取指定tree/blob被哪些commit引用了

TODO