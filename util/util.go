package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type DownLoadReader struct {
	io.Reader
	name  string
	total int64
	cur   int64
}

func (r *DownLoadReader) Read(p []byte) (n int, err error) {
	n, err = r.Reader.Read(p)
	r.cur += int64(n)
	fmt.Printf("\r [%s] %v/%v 进度: %.2f%%", r.name, r.cur, r.total, float64(r.cur*10000/r.total)/100)
	return
}

// DownLoad 下载
func DownLoad(url string, file string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	f, err := os.Create(file)
	if err != nil {
		return
	}
	defer f.Close()

	reader := &DownLoadReader{
		name:   file,
		Reader: resp.Body,
		total:  resp.ContentLength,
	}

	_, err = io.Copy(f, reader)
	return
}
