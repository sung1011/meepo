package service

import (
	"log"

	"github.com/sung1011/meepo/util"
)

// Nginx _
type Nginx struct {
	DownLoadURL string
}

// NewNginx _
func NewNginx() *Nginx {
	return &Nginx{
		DownLoadURL: "http://nginx.org/download/nginx-1.18.0.tar.gz",
	}
}

// BeforeSetup _
func (ngx *Nginx) BeforeSetup() {
	err := util.DownLoad(ngx.DownLoadURL, "/tmp/nginx.tar.gz")
	if err != nil {
		log.Panic(err)
	}
}

// DoSetup _
func (ngx *Nginx) DoSetup() {
}

// AfterSetup _
func (ngx *Nginx) AfterSetup() {
}
