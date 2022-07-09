package internal

import (
	"path"
	"path/filepath"
	"strings"
)

// CodeToContentId 删除 "-" -> 去空格 -> 英文小写
func CodeToContentId(code string) string {
	return strings.ToLower(strings.Trim(strings.Replace(code, "-", "", -1), " "))
}

func FilterFilename(fileName string, filter string) string {
	// 获取正确文件名
	fileName = filepath.Base(strings.ToLower(fileName))
	// 删除扩展名
	fileName = strings.TrimSuffix(fileName, path.Ext(fileName))
	// 转换过滤规则为数组
	filters := strings.Split(filter, "||")
	// 循环过滤
	for _, f := range filters {
		// 过滤
		fileName = strings.ReplaceAll(fileName, f, "")
	}
	// 将所有 . 替换为 -
	fileName = strings.ReplaceAll(fileName, ".", "-")
	// 过滤空格
	fileName = strings.TrimSpace(fileName)

	return fileName
}
