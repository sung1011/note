#! /bin/bash

# yum安装必要的基础程序， 其他用go或docker安装
# 修改系统配置 (用标记判定是否已经 >> 过)
# 是否强制重装
# 检查是否安装成功 输出所有ver
# 单步进行重装
# 交互模式


# const
{
    go_version=1.16.5

    zsh_theme=arrow
}

# ssh
{
    sed -i 's/^#ClientAliveInterval.*$/ClientAliveInterval 3600/g' /etc/ssh/sshd_config
    service sshd restart
}

# installer
{
    yum update -y
    yum install -y wget
    yum install -y git
}

# oh-my-zsh
{
    yum install -y zsh
    # 安装并切换oh-my-zsh
    CHSH=yes RUNZSH=yes sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
    # 配置
    sed -i "s/^ZSH_THEME.*/ZSH_THEME=${zsh_theme}/g" "${HOME}"/.zshrc
}


# docker
{
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
}

# go
{
    # download
    # wget https://golang.org/dl/go${go_version}.linux-amd64.tar.gz
    wget https://studygolang.com/dl/golang/go${go_version}.linux-amd64.tar.gz
    tar -C /usr/local -xzf go${go_version}.linux-amd64.tar.gz

    # $PATH
    if [ -z "$(grep "PATH=\$PATH:/usr/local/go/bin" "${HOME}"/.zshrc)" ]; then
        printf "\n" >> "${HOME}"/.zshrc
        echo "export PATH=\$PATH:/usr/local/go/bin" >> "${HOME}"/.zshrc
    fi
}

# nginx