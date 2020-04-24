# svn

## 定义

---

## 实例

- svn log -l 10 -v {fileName}
- svn cat -r 4 {fileName} // 查看某文件版本4的内容
- svn diff -r 4:6 {fileName} // 比较某文件版本4与6的内容差异
- svn checkout {url}　{local path}  --username {} --password {} // 检出
- svn checkout -r {version} {url} {local path} // 检出某版本
- svn lock, unlock
