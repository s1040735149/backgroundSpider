package logic

import (
	"fmt"
	"path"
)

const NetBianPath = "./netbianimg"

//爬取横图
func CrawNetbian(begin, end int) {
	err := NewPathIfNotExists(NetBianPath)
	if err != nil {
		fmt.Printf("mkdir err %s", err.Error())
		return
	}
	for i := begin; i <= end; i++ {
		Wg.Add(1)

		url := fmt.Sprintf("http://www.netbian.com/index_%d.htm", i)
		if i == 1 {
			url = "http://www.netbian.com/index.htm"
		}
		go CrawNetbianHandle(url)
	}
	Wg.Wait()
}

func CrawNetbianHandle(u string) {
	defer Wg.Done()
	data, err := GetUrlData(u)
	if err != nil {
		fmt.Printf("get data by url:%s err:%s", u, err.Error())
		return
	}
	sdata := string(data)

	list := CrawNetbianGetEachUrl(sdata)

	for _, v := range list {
		CrawNetbianEachHandle(v)
	}
}
func CrawNetbianGetEachUrl(sdata string) []string {
	f := false
	list := make([]string, 0)
	stack := ""
	for k := 0; k < len(sdata); {
		v := sdata[k]
		if string(v) == "<" && k+15 < len(sdata) && sdata[k:k+15] == "<a href=\"/desk/" {
			f = true
			k += 15
			continue
		}
		if string(v) == "." && f {
			f = false
			list = append(list, stack)
			stack = ""
		}
		if f {
			stack += string(v)
		}
		k++
	}
	return list
}
func CrawNetbianEachHandle(u string) {
	url := fmt.Sprintf("http://www.netbian.com/desk/%s-1920x1080.htm", u)
	data, err := GetUrlData(url)
	if err != nil {
		fmt.Printf("get data by url:%s err:%s", u, err.Error())
		return
	}
	sdata := string(data)
	f := false
	list := make([]string, 0)
	stack := ""
	for k := 0; k < len(sdata); {
		v := sdata[k]
		if string(v) == "<" && k+9 < len(sdata) && sdata[k:k+9] == "<a href=\"" {
			f = true
			k += 9
			continue
		}
		if string(v) == "\"" && f {
			f = false
			list = append(list, stack)
			stack = ""
		}
		if f {
			stack += string(v)
		}
		k++
	}
	for _, v := range list {
		if path.Ext(v) == ".jpg" {
			DownImg(v, NetBianPath)
		}
	}
}
