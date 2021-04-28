package logic

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
)

var Wg sync.WaitGroup

func GetUrlData(u string) ([]byte, error) {
	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func NewPathIfNotExists(path string) error {
	direx, err := PathExists(path)
	if err != nil {
		return err
	}
	if !direx {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			return err
		} else {
			fmt.Printf("mkdir %s success! \n", path)
		}
	}
	return nil
}

func DownImg(url, p string) {
	body, err := GetUrlData(url)
	if err != nil {
		fmt.Printf("get data by url:%s err:%s", url, err.Error())
		return
	}
	sr := fmt.Sprintf("%x", md5.Sum([]byte(url)))
	out, _ := os.Create(fmt.Sprintf("%s/img-%s%s", p, sr, path.Ext(url)))
	defer out.Close()
	_, err = io.Copy(out, bytes.NewReader(body))
	if err != nil {
		fmt.Println(fmt.Sprintf("down %s fail:%s", url, err.Error()))
		return
	}
	fmt.Println(fmt.Sprintf("down %s ok", url))
}
func GetBeginEnd() (int, int, error) {
	var begin, end int
	fmt.Println("---请输入开始页码：")
	_, err := fmt.Scanln(&begin)
	if err != nil {
		return 0, 0, err
	}

	fmt.Println("---请输入结束页码：")
	_, err = fmt.Scanln(&end)
	if err != nil {
		return 0, 0, err
	}

	if begin > end {
		return 0, 0, err
	}
	return begin, end, nil
}
