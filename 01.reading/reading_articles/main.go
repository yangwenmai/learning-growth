package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"
	"github.com/tidwall/gjson"
)

func main() {
	startTime := time.Now()
	githubAccessToken := os.Getenv("GITHUB_ACCESS_TOKEN")
	// err := os.Remove("data.txt")
	// handleError(err, "remove data.txt")
	// err = os.Remove("README.md")
	// handleError(err, "remove README.md")
	ioutil.WriteFile("data.txt", []byte(""), 0644)
	ioutil.WriteFile("README.md", []byte(""), 0644)
	file, err := os.OpenFile("data.txt", os.O_WRONLY|os.O_APPEND, 0644)
	handleError(err, "open data.txt")
	file2, err := os.OpenFile("README.md", os.O_WRONLY|os.O_APPEND, 0644)
	handleError(err, "open README.md")
	defer file.Close()
	defer file2.Close()
	md := "# 阅读记录\n\n" +
		"\n我会将想看/看过的文章，同步推送到 [yangwenmai/reading Issues](https://github.com/yangwenmai/reading/issues) 中，如果你对他们有任何想法与意见，欢迎在 issue 中评论。\n\n"

	md += "## Reading 关键词词云\n\n"
	md += "![](reading_wordcloud_output.png)\n\n"

	md += "## 安装脚本\n\n"
	md += "1. [词云可视化：四行 Python 代码轻松上手到精通](https://github.com/TommyZihao/zihaowordcloud/)\n"
	md += "2. `export GITHUB_ACCESS_TOKEN=[Your Github Personal ACCESS_TOKEN]`\n"
	md += "3. `go run main.go`\n"
	md += "4. `python reading_wordcloud.py`\n\n"

	md += "## Reading List\n\n"
	md += "| Issues ID | 标题 | 创建时间 |\n"
	md += "|----|----|----|\n"
	for i := 1; i < 22; i++ {
		githubURL := fmt.Sprintf("https://api.github.com/repos/yangwenmai/reading/issues?access_token=%s&page=%d&per_page=50", githubAccessToken, i)
		fmt.Println("Run: ", githubURL)
		resp, err := http.Get(githubURL)
		handleError(err, "get github issues url:"+githubURL)
		defer func() {
			if resp != nil && resp.Body != nil {
				resp.Body.Close()
			}
		}()
		body, err := ioutil.ReadAll(resp.Body)
		handleError(err, "resp body read failed.")
		gj := gjson.Parse(string(body))
		length := gj.Array()
		if len(length) > 0 {
			ret, mdd := getBody(gj)
			ret = filterStopWords(ret)
			file.WriteString(ret)
			md += mdd
			file2.WriteString(md)
		}
		fmt.Println("current cost time: ", time.Now().Sub(startTime))
	}
	md += "\n\n"
	file2.WriteString(md)
	fmt.Println("total cost time: ", time.Now().Sub(startTime))
}

func filterStopWords(ret string) string {
	stopwords := []string{"微信扫一扫使用小程序", "允许", "取消", "知道了", "确定>", "Hi!", "相关阅读", "在看", " ", "我们", "一个", "可以", "这个", "那个", "自己", "就是", "问题", "如果", "用户", "可能", "些", "这些", "那些", "需要", "不同", "现在", "通过", "不是", "还是", "所以", "能", "因为", "使用", "他们", "大家", "应该", "或者", "时候", "然后", "一样", "成为", "其实", "开始"}

	for _, word := range stopwords {
		ret = strings.ReplaceAll(ret, word, "")
	}
	return ret
}

func handleError(err error, msg string) {
	if err != nil {
		fmt.Println("[ERROR] ", err, ", msg:", msg)
	}
}

func getBody(gj gjson.Result) (string, string) {
	result := ""
	md := ""
	gj.ForEach(func(key, value gjson.Result) bool {
		issueNumber := value.Get("number").Int()
		title := strings.TrimSpace(value.Get("title").String())
		title = strings.ReplaceAll(title, "|", "-")
		createdAt := value.Get("created_at").String()
		ret := value.Get("body").String()
		targetURL, err := getTargetURL(issueNumber, title, ret)
		if targetURL == "" {
			targetURL, err = getTargetURL2(issueNumber, title, ret)
			if targetURL == "" {
				if err != nil {
					fmt.Println(err)
				}
				return false
			}
			if err != nil {
				fmt.Println(err)
				return false
			}
		}
		if !strings.HasPrefix(targetURL, "http") {
			return false
		}
		md += fmt.Sprintf("| [#%d](https://github.com/yangwenmai/reading/issues/%d) | [%s](%s) | %s |\n",
			issueNumber, issueNumber, title, targetURL, getBeijingTime(createdAt))
		resp, err := doRequest(targetURL)
		handleError(err, "fetch origin URL: "+targetURL)
		if resp != nil {
			result += getHTMLText(resp)
		}
		return true
	})
	return result, md
}

func getBeijingTime(createdAt string) string {
	t, _ := time.Parse("2006-01-02T15:04:05Z", createdAt)
	loc, err := time.LoadLocation("Asia/Chongqing")
	handleError(err, "load time location failed.")
	t = t.In(loc)
	return t.Format("2006-01-02 15:04:05")
}

func getTargetURL2(issueNumber int64, title string, ret string) (string, error) {
	splitStr := "via Instapaper"
	if !strings.Contains(ret, splitStr) {
		return "", fmt.Errorf("Unsupport parse, issues id: https://github.com/yangwenmai/reading/issues/%d, title: %s\n", issueNumber, title)
	}
	idx := strings.Index(ret, splitStr)
	return strings.TrimSpace(ret[idx+len(splitStr):]), nil
}

func getTargetURL(issueNumber int64, title string, ret string) (string, error) {
	splitStr := "](http"
	if !strings.Contains(ret, splitStr) {
		return "", fmt.Errorf("Unsupport parse, issues id: https://github.com/yangwenmai/reading/issues/%d, title: %s\n", issueNumber, title)
	}
	idx := strings.Index(ret, splitStr)
	tmpStr := strings.TrimSpace(ret[idx+len(splitStr):])
	targetURL := ""
	if strings.HasSuffix(tmpStr, ")") {
		targetURL = "http" + tmpStr[:len(tmpStr)-1]
	} else {
		targetURL = "http" + tmpStr
	}
	if strings.HasSuffix(targetURL, ".pdf") {
		return "", fmt.Errorf("is pdf")
	}
	return targetURL, nil
}

func doRequest(crawlerUrl string) (*http.Response, error) {
	request, err := http.NewRequest("GET", crawlerUrl, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "request err with url:"+crawlerUrl)
	}

	// to avoid unsupported protocol scheme err
	if request.URL.Scheme == "" {
		request.URL.Scheme = "http"
	}

	request.Header = http.Header{"User-Agent": []string{"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36"}}
	crawlerClient := http.Client{Timeout: 15 * time.Second}
	resp, err := crawlerClient.Do(request)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func getHTMLText(resp *http.Response) string {
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	handleError(err, "reader resp body failed.")

	bodyContents := doc.Find("body").Each(func(_ int, selection *goquery.Selection) {
		selection.Find("script").Remove()
	})
	return strings.TrimSpace(bodyContents.Text())
}
