package main

import (
	"backGroundSpider/logic"
	"fmt"
)

const (
	CrawlTypeNetbian = 1
	CrawlTypeDuitang = 2
)

func main() {
	var crawlType int

chooseCrawType:
	s :=
		`---选择要爬取的类型:
1:netbian横图
2:duitang长图
0:退出`
	fmt.Println(s)

	_, err := fmt.Scanln(&crawlType)
	if err != nil {
		fmt.Println("请输入合法的类型")
	}

	switch crawlType {
	case CrawlTypeNetbian:
		begin, end, err := logic.GetBeginEnd()
		if err != nil {
			fmt.Println("请输入合法的数字")
			return
		}
		logic.CrawNetbian(begin, end)
	case CrawlTypeDuitang:
		begin, end, err := logic.GetBeginEnd()
		if err != nil {
			fmt.Println("请输入合法的数字")
			return
		}
		logic.CrawDuitang(begin, end)
	case 0:
		fmt.Println("谢谢使用,再见!")
		return
	default:
		fmt.Println("暂不支持的类型")
		goto chooseCrawType
	}
}
