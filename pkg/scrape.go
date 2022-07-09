package pkg

import (
	"AV_Meta_Capture/internal"
	"log"
)

var sites = map[string]string{
	"avmoo": "https://www.avmoo.com",
}

var config = internal.Config{
	TargetFolder: "C:/Users/Yin/Desktop/视频处理/已发布",
}

func Scrape() {

	// 通过目录获取视频列表
	fileInfoSlice, err := internal.GetFileInfoFromDir(config.TargetFolder)
	if err != nil {
		log.Printf("获取视频列表失败, err: %s", err)
	}

	// 遍历视频列表
	for _, info := range *fileInfoSlice {
		metaTree, _ := AvmooWork(&info)
		log.Printf("%#v", metaTree)
	}
}

// 读取指定文件夹内的所有视频文件名(排除特殊目录)
// 将文件名保存到切片中
// 遍历文件名
// 去除多余字符/网址/广告
// 遍历刮削的网址
// 开始刮削
// 获取刮削结果 MovieMetaTree
// 生成 NFO 信息 NfoMetaTree
// 下载海报和背景墙
// 创建文件夹
// 移动文件到文件夹
// 移动文件夹到成功目录
