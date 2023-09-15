package example

import (
	"bufio"
	"crawler/collect"
	"fmt"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

// tag v0.0.3
func mainV3() {
	url := "https://www.thepaper.cn/"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("fetch url error:%v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%v", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}

	numLinks := strings.Count(string(body), "<a")
	fmt.Printf("homepage has %d links!\n", numLinks)
}

// tag v0.0.4
func Fetch(url string) ([]byte, error) {

	resp, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error status code:%d", resp.StatusCode)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := collect.DeterminEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader, e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// tag v0.0.5
var headerRe = regexp.MustCompile(`<div class="news_li"[\s\S]*?<h2>[\s\S]*?<a.*?target="_blank">([\s\S]*?)</a>`)

func mainV5() {
	url := "https://www.thepaper.cn/"
	body, err := Fetch(url)

	if err != nil {
		fmt.Println("read content failed:%v", err)
		return
	}
	matches := headerRe.FindAllSubmatch(body, -1)
	for _, m := range matches {
		fmt.Println("fetch card news:", string(m[1]))
	}
}

// tag v0.0.6
//func mainV6() {
//	url := "https://www.thepaper.cn/"
//	body, err := Fetch(url)
//
//	if err != nil {
//		fmt.Println("read content failed:%v", err)
//		return
//	}
//	doc, err := htmlquery.Parse(bytes.NewReader(body))
//	if err != nil {
//		fmt.Println("htmlquery.Parse failed:%v", err)
//	}
//	nodes := htmlquery.Find(doc, `//div[@class="news_li"]/h2/a[@target="_blank"]`)
//
//	for _, node := range nodes {
//		fmt.Println("fetch card ", node.FirstChild.Data)
//	}
//}

// tag v0.0.9
//func main() {
//	url := "https://www.thepaper.cn/"
//	body, err := Fetch(url)
//
//	if err != nil {
//		fmt.Println("read content failed:%v", err)
//		return
//	}
//
//	// 加载HTML文档
//	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(body))
//	if err != nil {
//		fmt.Println("read content failed:%v", err)
//	}
//
//	doc.Find("div.news_li h2 a[target=_blank]").Each(func(i int, s *goquery.Selection) {
//		// 获取匹配标签中的文本
//		title := s.Text()
//		fmt.Printf("Review %d: %s\n", i, title)
//	})
//}
