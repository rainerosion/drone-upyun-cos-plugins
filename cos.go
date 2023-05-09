package main

import (
	"github.com/upyun/go-sdk/v3/upyun"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var (
		operator       = GetEnv("up_operator")
		bucket         = GetEnv("up_bucket")
		password       = GetEnv("up_password")
		localBasePath  = GetEnv("local_base_path")
		remoteBasePath = GetEnv("remote_base_path")
	)
	yun := upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   bucket,
		Operator: operator,
		Password: password,
	})
	var fileList []string
	var err error = nil
	filepath.WalkDir(localBasePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			Exit(err)
		}
		// 添加文件路径到文件列表
		if IsFile(path) {
			fileList = append(fileList, path)
		}
		return nil
	})
	// 断点续传 文件大于 10M 才会分片
	resume := &upyun.MemoryRecorder{}
	// 若设置为 nil，则为正常的分片上传
	yun.SetRecorder(resume)
	separator := "/"
	for _, path := range fileList {
		if strings.HasPrefix(path, "/") {
			separator = ""
		}
		remotePath := remoteBasePath + separator + path
		err = yun.Put(&upyun.PutObjectConfig{
			Path:      remotePath,
			LocalPath: path,
		})
		if err != nil {
			log.Printf("local file: %v remote file: %v upload status: fail errmsg:%v\n", path, remotePath, err.Error())
		} else {
			log.Printf("local file: %v remote file: %v upload status: success\n", path, remotePath)
		}
	}
}

func Exit(err error) {
	log.Println("error => " + err.Error())
	os.Exit(1)
}

// IsDir 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// GetEnv drone settings下的变量将会转换为PLUGIN_开头的环境变量
func GetEnv(name string) string {
	return os.Getenv("PLUGIN_" + strings.ToUpper(name))
}
