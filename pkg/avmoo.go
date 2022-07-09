package pkg

import (
	"AV_Meta_Capture/internal"
	"errors"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strconv"
	"strings"
)

type Avmoo internal.MovieMetaTree

func (tree *Avmoo) GetRootDom(s string, c *colly.Collector) (*colly.HTMLElement, error) {
	var root *colly.HTMLElement

	// 获取 avmoo 代理地址
	home, err := internal.GetBodyHTML(c, "https://www.avmoo.com")
	if err != nil {
		return nil, err
	}
	host, hostExists := home.DOM.Find("div.text h4:first-of-type a").Attr("href")
	if !hostExists {
		return nil, errors.New("未找到 Avmoo 代理地址，请检查设置")
	}

	// 搜索番号，查询结果
	searchStr := fmt.Sprintf("%s/cn/search/%s", host, s)
	element, err := internal.GetBodyHTML(c, searchStr)
	if err != nil {
		return nil, err
	}

	numberExists := false
	element.DOM.Find("#waterfall .item").Each(func(i int, selection *goquery.Selection) {
		href, exists := selection.Find("a.movie-box").Attr("href")
		if !exists {
			err = errors.New("未获取到页面链接，退出流程")
			return
		}
		// 获取页面识别码
		number := selection.Find("date:first-of-type").Text()
		// 判断页面识别码和需要刮削的番号是否一致
		if internal.CodeToContentId(s) == internal.CodeToContentId(number) {
			numberExists = true
			targetElement, errStr := internal.GetBodyHTML(c, href)
			if errStr != nil {
				err = errStr
				return
			}
			root = targetElement
		}
	})

	if !numberExists {
		return nil, errors.New(fmt.Sprintf("[%s] 未找到, 站点: [Avmoo], 链接: %s", s, searchStr))
	}

	if err != nil {
		return nil, err
	}

	return root, nil
}

func (tree *Avmoo) GetMetaTree(c *colly.HTMLElement, s string) (*internal.MovieMetaTree, error) {
	var studio, set, genre []string
	var actors []internal.Actor

	root := c.DOM
	container := root.Find(`.container`)
	divInfo := root.Find(`div.info`)
	avatar := root.Find("div#avatar-waterfall")

	// 识别码
	number := divInfo.Find(`span:contains("识别码:")`).Siblings().Text()

	// 检查号码是否匹配
	if internal.CodeToContentId(number) != internal.CodeToContentId(s) {
		errStr := fmt.Sprintf("请求爬取的号码和页面内号码不一致!\n请求号码: %s\n页面号码: %s", s, number)
		return &internal.MovieMetaTree{}, errors.New(errStr)
	}

	// 标题
	title := strings.TrimSpace(container.Find("h3").Text())
	// 背景墙
	fanArt, _ := container.Find(".screencap img").Attr("src")
	// 发行日期
	premiered := divInfo.Find(`span:contains("发行时间:")`).Parent().Text()
	// 电影时长
	runtime := divInfo.Find(`span:contains("长度:")`).Parent().Text()
	// 导演
	director := divInfo.Find(`span:contains("导演:")`).Parent().Text()

	// 制作商
	divInfo.Find(`p:contains("制作商:")`).Next().Find("a").Each(func(i int, selection *goquery.Selection) {
		studio = append(studio, selection.Text())
	})
	// 发行商
	divInfo.Find(`p:contains("发行商:")`).Next().Find("a").Each(func(i int, selection *goquery.Selection) {
		studio = append(studio, selection.Text())
	})
	// 系列
	divInfo.Find(`p:contains("系列:")`).Next().Find("a").Each(func(i int, selection *goquery.Selection) {
		set = append(set, selection.Text())
	})
	// 类别
	divInfo.Find(`p:contains("类别:")`).Next().Find("a").Each(func(i int, selection *goquery.Selection) {
		genre = append(genre, selection.Text())
	})
	// 演员
	avatar.Find("a.avatar-box").Each(func(i int, selection *goquery.Selection) {
		name := selection.Find("span").Text()
		actor := internal.Actor{
			Name:  name,
			Role:  name,
			Order: strconv.Itoa(i),
		}
		actors = append(actors, actor)
	})

	// 清洗多余数据
	premiered = strings.Trim(strings.Replace(premiered, "发行时间:", "", -1), " ")
	director = strings.Trim(strings.Replace(director, "导演:", "", -1), " ")
	runtime = strings.Replace(runtime, "长度:", "", -1)
	runtime = strings.Trim(strings.Replace(runtime, "分钟", "", -1), " ")

	metaTree := &Avmoo{
		Number:    number,
		Title:     title,
		Premiered: premiered,
		Runtime:   runtime,
		Director:  director,
		Outline:   "",
		Thumb:     "",
		FanArt:    fanArt,
		DateAdded: "",
		Studio:    studio,
		Set:       set,
		Genre:     genre,
		Actor:     actors,
	}

	return (*internal.MovieMetaTree)(metaTree), nil
}

func AvmooWork(info *internal.FileInfo) (*internal.MovieMetaTree, error) {
	var avmoo Avmoo
	c := internal.InitColly()
	number := info.Name

	dom, err := avmoo.GetRootDom(number, c)
	if err != nil {
		return &internal.MovieMetaTree{}, err
	}

	metaTree, err := avmoo.GetMetaTree(dom, number)
	if err != nil {
		return &internal.MovieMetaTree{}, err
	}

	metaTree.DateAdded = info.ModTime

	return metaTree, nil
}
