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
}

// NewNginx _
func NewNginx() *Nginx {
	return &Nginx{
		DownLoadURL:     "http://nginx.org/download/nginx-1.18.0.tar.gz",
		DownLoadingPath: "/tmp/nginx.tar.gz.meepo",
		DownLoadPath:    "/tmp/nginx.tar.gz",
	}
}

// BeforeSetup _
func (ngx *Nginx) BeforeSetup() (err error) {
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
	return nil
}

// AfterSetup _
func (ngx *Nginx) AfterSetup() (err error) {
	// TODO 启动
	return nil
}
