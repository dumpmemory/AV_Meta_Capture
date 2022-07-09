package internal

import "encoding/xml"

type FileInfo struct {
	Name    string
	ModTime string
}

// Actor 是演员信息的结构体
type Actor struct {
	Name  string `xml:"name,omitempty"`
	Role  string `xml:"role,omitempty"`
	Order string `xml:"order,omitempty"`
	Thumb string `xml:"thumb,omitempty"`
}

// MovieMetaTree 是电影信息树的结构体
type MovieMetaTree struct {
	Number    string   // 号码
	Title     string   // 标题
	Premiered string   // 发行日期
	Runtime   string   // 时长
	Director  string   // 导演
	Outline   string   // 大纲
	Thumb     string   // 海报
	FanArt    string   // 背景墙
	DateAdded string   // 添加时间
	Studio    []string // 制造商
	Set       []string // 系列
	Genre     []string // 流派
	Actor     []Actor  // 演员
}

// NfoMetaTree 是 NFO 信息树的结构体.
// 参考: https://kodi.wiki/view/NFO_files/Movies
type NfoMetaTree struct {
	XMLName       xml.Name `xml:"movie"`                   // 根节点名称
	Number        string   `xml:"number,omitempty"`        // 号码
	Title         string   `xml:"title,omitempty"`         // 标题
	SortTitle     string   `xml:"sorttitle,omitempty"`     // 排序标题
	OriginalTitle string   `xml:"originaltitle,omitempty"` // 原始标题
	Premiered     string   `xml:"premiered,omitempty"`     // 发行日期
	Year          string   `xml:"year,omitempty"`          // 发行年份
	Runtime       string   `xml:"runtime,omitempty"`       // 时长
	Director      string   `xml:"director,omitempty"`      // 导演
	Outline       string   `xml:"outline,omitempty"`       // 大纲
	Plot          string   `xml:"plot,omitempty"`          // 情节
	Thumb         string   `xml:"thumb,omitempty"`         // 海报
	FanArt        string   `xml:"fanart,omitempty"`        // 背景墙
	Country       string   `xml:"country,omitempty"`       // 国家
	DateAdded     string   `xml:"dateadded,omitempty"`     // 添加时间
	Studio        []string `xml:"studio,omitempty"`        // 制造商
	Set           []string `xml:"set,omitempty"`           // 系列
	Genre         []string `xml:"genre,omitempty"`         // 流派
	Tag           []string `xml:"tag,omitempty"`           // 标签
	Actor         []Actor  `xml:"actor,omitempty"`         // 演员
}

type Config struct {
	TargetFolder string `json:"targetFolder"`
}
