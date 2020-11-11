#! /bin/bash

# yum安装必要的基础程序， 其他用go安装
# 修改系统配置 (用标记判定是否已经 >> 过)
# 是否强制重装
# 检查是否安装成功 输出所有ver
# 单步进行重装
# 交互模式

# ssh
echo ClientAliveInterval=3600 >> /etc/ssh/sshd_config
service sshd restart

yum update
# git
yum install -y git
# zsh
yum install -y zsh
# oh-my-zsh
sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
# docker
yum remove docker \
            docker-client \
            docker-client-latest \
            docker-common \
            docker-latest \
            docker-latest-logrotate \
            docker-logrotate \
            docker-engine
yum install -y yum-utils 
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo
yum install -y docker-ce docker-ce-cli containerd.io
# go
# wget https://golang.org/dl/go1.15.4.linux-amd64.tar.gz
wget https://studygolang.com/dl/golang/go1.15.4.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.4.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin

# 追加配置
appendConfig(){
echo 1
}

# 安装
# $1 ver
setup() {
    echo 
}