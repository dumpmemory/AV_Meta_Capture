package internal

import "github.com/gocolly/colly"

// GetMetaTree 这个接口定义了每个网站的爬虫需要实现的方法
type GetMetaTree interface {
	GetRootDom(s string, c *colly.Collector) (*colly.HTMLElement, error) // GetRootDom 通过番号获取页面 Body 数据
	GetMetaTree(c *colly.HTMLElement, s string) (*MovieMetaTree, error)  // GetMetaTree 通过 Body 获取 MovieMetaTree
}
