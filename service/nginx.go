package service

import (
	"os"

	"github.com/sung1011/meepo/util"
)

// Nginx _
type Nginx struct {
	DownLoadURL     string
	DownLoadPath    string
	DownLoadingPath string
	OPT             string
}

// NewNginx _
func NewNginx() *Nginx {
	return &Nginx{
		DownLoadURL:     "http://nginx.org/download/nginx-1.18.0.tar.gz",
		DownLoadingPath: "/tmp/nginx.tar.gz.meepo",
		DownLoadPath:    "/tmp/nginx.tar.gz",
		OPT:             "/opt/",
	}
}

// BeginSetup _
func (ngx *Nginx) BeginSetup() (err error) {
	// 文件已经下载 直接返回
	_, statErr := os.Stat(ngx.DownLoadPath)
	if statErr == nil || os.IsExist(statErr) {
		return
	}
	err = util.DownLoad(ngx.DownLoadURL, ngx.DownLoadingPath)
	if err != nil {
		return
	}
	err = os.Rename(ngx.DownLoadingPath, ngx.DownLoadPath)
	if err != nil {
		return
	}
	return
}

// DoSetup _
func (ngx *Nginx) DoSetup() (err error) {
	// TODO 解压 编译
	err = util.UnGzip(ngx.DownLoadPath, ngx.OPT)
	return
}

// EndSetup _
func (ngx *Nginx) EndSetup() (err error) {
	// TODO 启动
	return nil
}
