#! /bin/bash

# yum安装必要的基础程序,  其他用go或docker安装
# 修改系统配置 (用标记判定是否已经 >> 过)
# 是否强制重装
# 检查是否安装成功 输出所有ver
# 单步进行重装
# 交互模式

# append string, if the string not exists in the file.
appendNX() {
    str=$1
    file=$2

    if [ ! -f "${file}" ]; then
        echo "file not exists: ${file}"
        return
    fi
    if [ -z "$(grep "${str}" "${file}")" ]; then
        printf "\n" >> "${file}"
        echo "${str}" >> "${file}"
    fi
}

# const
{
    go_version=1.16.5
    docker_compose_version=1.29.2

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
    yum install -y epel-release
    yum install -y wget
    yum install -y git
    yum install -y svn
    yum install -y gcc
}

# oh-my-zsh
{
    yum install -y zsh
    # 安装并切换oh-my-zsh
    CHSH=yes RUNZSH=yes sh -c "$(curl -fsSL https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
    # 配置
    sed -i "s/^ZSH_THEME.*/ZSH_THEME=${zsh_theme}/g" "${HOME}"/.zshrc
    # 插件 
        # autojump
        git clone git://github.com/wting/autojump.git
        cd autojump && ./install.py
        appendNX "[[ -s /root/.autojump/etc/profile.d/autojump.sh ]] && source /root/.autojump/etc/profile.d/autojump.sh; autoload -U compinit && compinit -u" "${HOME}"/.zshrc
        # zsh-autosuggestions
        git clone https://github.com/zsh-users/zsh-autosuggestions "${ZSH_CUSTOM:-~/.oh-my-zsh/custom}"/plugins/zsh-autosuggestions
        # zsh-syntax-highlighting
        git clone https://github.com/zsh-users/zsh-syntax-highlighting.git "${ZSH_CUSTOM:-~/.oh-my-zsh/custom}"/plugins/zsh-syntax-highlighting

        # open above
        sed -i 's/^plugins=.*$/plugins=(git zsh-autosuggestions zsh-syntax-highlighting autojump)/g' "${HOME}"/.zshrc
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

    # 启动
    systemctl restart docker.service
    # 开机自启动
    systemctl enable docker.service
}

# docker-compose
{
    curl -L "https://github.com/docker/compose/releases/download/${docker_compose_version}/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
    chmod +x /usr/local/bin/docker-compose
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