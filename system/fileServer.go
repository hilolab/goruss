package system

import (
	"path"
	"strings"
)

type FileServer struct {
	StaticDir map[string]string
}

func NewFileServer() *FileServer {
	return &FileServer{
		StaticDir: make(map[string]string),
	}
}

//设置静态资源目录
func (f *FileServer) SetStaticDir(url, dir string) bool {
	if len(url) == 0 || len(dir) == 0 {
		return false
	}
	url = strings.Trim(url, "/")
	f.StaticDir[url] = strings.TrimRight(dir, "/")

	return true
}

//判断是否是静态资源目录
func (f *FileServer) IsStaticDir(url string) (string, bool) {
	//防止../
	url = path.Clean(url)
	url = strings.Trim(url, "/")
	//判断是否是ico或rebots
	if url == "favicon.ico" || url == "robots.txt" {
		return url, true
	}

	for k, v := range f.StaticDir {
		index := strings.Index(url, k)
		if index == 0 {
			return v + "/" + url[index:], true
		}
	}

	return "", false
}
