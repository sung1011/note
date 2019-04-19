# git

## cmd
- merge, rebase, stash, cherry-pick, revert, blame, reset, tag, reflog, bisect...
## 版本撤销
-revert
    - 提交一个与指定commit内容相反的commit
    - 若在主分支revert一个功能分支， 则该功能分支无法重新merge到主分支， 可以用cherry-pick。
## 版本回溯
- 回溯指定版本 git checkout
    - `git checkout {commit_id} && git checkout -b {new_branch_name}`
- 回溯指定版本、回溯n个版本 git reset --hard
    - `git reset --hard[^回退上一版本|^^回退上两个版本|~n回退上n个版本| commit_id回退到某一版本] && git push --force`
## 版本整合
- rebase
    - 在未合并的分支, 合并分支中的多个commit为一个commit。