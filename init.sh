#! /bin/bash

# ==================== const

# 程序目录
CRACK_PATH="/usr/local/crack/"
# 源码目录
CRACK_PKG_PATH="/usr/local/crack/pkg/"
# bin目录
CRACK_BIN_PATH="/usr/local/bin/"

# go version
GO_VER="1.15.6"

# ==================== func

function isCmdExist() {
	local cmd="$1"
  	if [ -z "$cmd" ]; then
		echo "Usage isCmdExist yourCmd"
		return 1
	fi

    if ! which "$cmd" >/dev/null 2>&1;
    then
		return 0
	fi

	return 2
}

# ==================== install

# -------- dir

mkdir -p "$CRACK_PATH"
mkdir -p "$CRACK_PKG_PATH"

# -------- go

if ! isCmdExist go; then
    # download golang
    curl -sL https://golang.org/dl/go"$GO_VER".linux-amd64.tar.gz > /tmp/go"$GO_VER".tar.gz
    # install golang
    tar -zxvf /tmp/go"$GO_VER".tar.gz -C "$CRACK_PATH"
    # link bin
    ln -s "$CRACK_PATH"go/bin/go "$CRACK_BIN_PATH"
fi

# goenv  
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

# -------- git

if ! isCmdExist git; then
    # install
    yum install -y git
fi

# -------- meepo

# install
go get -u -v github.com/sung1011/meepo
