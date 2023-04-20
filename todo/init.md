# init

## brew

```js
    # install
        autojump
        bat
        exa
        gvm
        fzf
        nvm
        ncdu
        pidof
        zoxide // zoxide is a smarter cd command, inspired by z and autojump.
        git-flow
```

## iterm2

```js
    # Theme
        Preference.Appearance.General.Theme = Minimal
    # 按键映射
        Preferences.Profiles.Keys = Natural Text Editing
    # 文字大小
        Preferences.Profiles.Text.Font = 14
    # Colors
        Preferences.Profiles.Colors.Color-Presets = Tango Dark
    # Unlimited scrollback
        Preferences.Profiles.Terminal.Unlimited-scrollback
    # dedicated window
        Perferences.Keys.Hotkey.Create-a-Dedicated-Hotkey-Window.Double-tap-key = ^Control
```

## ssh

```js
    # 生成.ssh 公钥私钥
        ssh-keygen -t rsa -C "tickles@xxx"
```

## .zshrc

```js
    # env
        export GO111MODULE="on" # go

    # theme
        ZSH_THEME="arrow"

    # proxy
        export no_proxy="localhost,127.0.0.1,localaddress,.localdomain.com"
        export http_proxy="http://127.0.0.1:7890"
        export https_proxy=$http_proxy

    # alias
        alias gacp='git add .;git commit -m \"ig:\";git push'
        alias glacp='git pull;git add .;git commit -m \"ig:\";git push'
        alias dk='docker'
        alias dki='docker image'
        alias dkc='docker container'
        alias dkcp='docker-compose'
        alias ks='kubectl'
        alias mnk='minikube'
        alias vf='vim $(fzf)'
        alias readlink='greadlink'
        alias cat='bat --paging=never --plain'
        alias cd='z' # zoxide

    # zoxide
        eval "$(zoxide init zsh)"
```

## vim

```js
    # install spf13-vim
        curl https://j.mp/spf13-vim3 -L > spf13-vim.sh && sh spf13-vim.sh

    # plugin
        git clone https://github.com/zsh-users/zsh-autosuggestions ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-autosuggestions # install
        plugins=( [plugins...] zsh-autosuggestions) # .zshrc 

        git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting # install
        plugins=( [plugins...] zsh-syntax-highlighting) # .zshrc

    # config .vimrc.local
        set shell=bash\ -i
        filetype plugin on
        syntax on
        set nospell
        set encoding=utf-8
        set wrap

        "tagbar
        "autocmd VimEnter * nested :TagbarOpen

        " syntastic
        let g:syntastic_always_populate_loc_list = 0
        let g:syntastic_auto_loc_list = 0
        let g:syntastic_check_on_open = 0
        let g:syntastic_check_on_wq = 0

        " colorscheme
        "let g:solarized_termcolors=16
        color fx "Load a colorscheme

        set foldmethod=manual "不折叠 另需要手动修改PIV/syntax/php.vim
        let g:DisableAutoPHPFolding = 1
        let g:PIVAutoClose = 1

        set nobackup " 不备份

        " keybindings
        " - map operate
        let mapleader = ","
        map <leader>b :buffers<cr>
        map <leader>1 :b1<cr>
        map <leader>2 :b2<cr>
        map <leader>3 :b3<cr>
        map <leader>4 :b4<cr>
        map <leader>5 :b5<cr>
        map <leader>6 :b6<cr>
        map <leader>d :bd<cr>
        map <leader>. :bn<cr>
        map <leader>m :bp<cr>
        map <leader>w :w<cr>
        map <leader>W :wq<cr>
        map <leader>q :q<cr>
        map <leader>Q :q!<cr>
        map <leader>f :Ack
        " - map git
        map <leader>gll :Git pull<cr>
        map <leader>ga :Git add .<cr>
        map <leader>gacp :!git pull;git add .;git commit -m 'ig:';git push<cr>
        " - map run script
        map <Leader>l :%!python3<CR> "运行py3
        map <Leader>p :%!php<CR> "运行php
        " - map format
        map <leader>= gg=G'' 
```

## app

```js
    google Chrome
    iterm2
    vscode
    mpv
    迅雷
    wechat
    eudic 欧路词典
    百度云盘 BaiduNetdisk
    dash
    snipaste
        - ctrl+cmd+s
        - ctrl+cmd+p
    docker
    qq音乐
    raycast
```