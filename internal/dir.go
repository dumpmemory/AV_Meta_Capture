package internal

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// IsDir 判断所给路径是否是一个文件夹
func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

// IsExist 判断所给路径文件/文件夹是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// GetFileInfoFromDir 从目录中获取视频列表
func GetFileInfoFromDir(path string) (*[]FileInfo, error) {
	var infoSlice []FileInfo
	// 检查目录是否存在
	if !IsDir(path) {
		return &infoSlice, errors.New("未找到指定目录，请检查配置文件")
	}

	// 遍历目录中的文件
	fileInfoSlice, err := ioutil.ReadDir(path)
	if err != nil {
		errStr := fmt.Sprintf("遍历文件夹失败! err: %s", err)
		return &infoSlice, errors.New(errStr)
	}

	// 获取目录中的视频
	for _, info := range fileInfoSlice {
		// 获取后缀
		ext := filepath.Ext(info.Name())
		if ext == "" {
			break
		}

		// 获取视频格式
		fileName := info.Name()
		reg := regexp.MustCompile(".(mp4|wmv|avi|mkv)$")
		suffix := reg.FindString(fileName)

		// 获取文件名
		fileName = strings.Replace(fileName, suffix, "", -1)
		modTime := info.ModTime().Format("2006-01-02")

		infoSlice = append(infoSlice, FileInfo{
			Name:    fileName,
			ModTime: modTime,
		})
	}
	return &infoSlice, nil
}
