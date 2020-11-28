package service

import (
	"fmt"
	"os"

	"github.com/sung1011/meepo/util"
)

// Nginx _
type Nginx struct {
	Ver              string
	DownLoadURL      string
	DownLoadPath     string
	DownLoadPathTemp string
	Crack            string
	Configure        []string
}

// NewNginx _
func NewNginx() *Nginx {
	ver := "1.18.0"
	// 裂缝 meepo喜欢将他的敌人困在裂缝中 包安装的目录
	crack := "/usr/local/crack/"
	return &Nginx{
		Ver:              ver,
		DownLoadURL:      fmt.Sprintf("http://nginx.org/download/nginx-%s.tar.gz", ver),
		DownLoadPath:     fmt.Sprintf("/tmp/nginx-%s.tar.gz", ver),
		DownLoadPathTemp: fmt.Sprintf("/tmp/nginx-%s.tar.gz.temp", ver),
		Crack:            crack,
		Configure:        []string{"--prefix=/opt/ngx", "--http-log-path=/opt/access.log"},
	}
}

// BeginSetup _
func (ngx *Nginx) BeginSetup() (err error) {
	// 文件已经下载 则直接返回
	_, statErr := os.Stat(ngx.DownLoadPath)
	if statErr == nil || os.IsExist(statErr) {
		return
	}
	err = util.DownLoad(ngx.DownLoadURL, ngx.DownLoadPathTemp)
	if err != nil {
		return
	}
	err = os.Rename(ngx.DownLoadPathTemp, ngx.DownLoadPath)
	if err != nil {
		return
	}
	return
}

// DoSetup _
func (ngx *Nginx) DoSetup() (err error) {
	err = util.UnGzip(ngx.DownLoadPath, ngx.Crack)
	if err != nil {
		return
	}
	err = os.Chdir(ngx.Crack + "nginx-" + ngx.Ver)
	if err != nil {
		return
	}
	util.RunCmd("./configure", ngx.Configure...)
	util.RunCmd("make")
	util.RunCmd("make", "install")
	return
}

// EndSetup _
func (ngx *Nginx) EndSetup() (err error) {
	// TODO 启动
	return nil
}
