package internal

import (
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

// InitColly 简单的初始化 Collector ，没有其他参数
func InitColly() *colly.Collector {
	c := colly.NewCollector()
	// 随机UA
	extensions.RandomUserAgent(c)
	// referer自动填写
	extensions.Referer(c)

	return c
}

// GetBodyHTML 爬取并返回页面的 body 内容
func GetBodyHTML(c *colly.Collector, domain string) (*colly.HTMLElement, error) {
	var body *colly.HTMLElement

	c.OnHTML("body", func(element *colly.HTMLElement) {
		body = element
	})

	// 请求页面
	err := c.Visit(domain)
	if err != nil {
		return nil, err
	}
	return body, nil
}
