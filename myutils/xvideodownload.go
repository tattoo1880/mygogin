package myutils

import (
	"github.com/imroc/req/v3"
	"net/http"
	"regexp"
)

type GetXVideoDownload interface {
	GetXVideoDownloadUrl(url string) (string, error)
}

type XVideoDownload struct{}

func NewXVideoDownload() GetXVideoDownload {
	return &XVideoDownload{}
}

func (x *XVideoDownload) GetXVideoDownloadUrl(url string) (string, error) {
	client := req.C()
	request := client.R()

	request.SetHeaders(map[string]string{
		"Accept":                      "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7",
		"Accept-Language":             "zh-CN,zh;q=0.9",
		"Connection":                  "keep-alive",
		"Sec-Fetch-Dest":              "document",
		"Sec-Fetch-Mode":              "navigate",
		"Sec-Fetch-Site":              "none",
		"Sec-Fetch-User":              "?1",
		"Upgrade-Insecure-Requests":   "1",
		"User-Agent":                  "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/137.0.0.0 Safari/537.36",
		"device-memory":               "8",
		"sec-ch-ua":                   `"Google Chrome";v="137", "Chromium";v="137", "Not/A)Brand";v="24"`,
		"sec-ch-ua-arch":              `"x86"`,
		"sec-ch-ua-bitness":           `"64"`,
		"sec-ch-ua-full-version":      `"137.0.7151.105"`,
		"sec-ch-ua-full-version-list": `"Google Chrome";v="137.0.7151.105", "Chromium";v="137.0.7151.105", "Not/A)Brand";v="24.0.0.0"`,
		"sec-ch-ua-mobile":            "?0",
		"sec-ch-ua-model":             `""`,
		"sec-ch-ua-platform":          `"macOS"`,
		"sec-ch-ua-platform-version":  `"13.7.5"`,
		"viewport-width":              "1440",
	})
	//cookies = {
	//	'dscld': 'true',
	//		'session_ath': 'black',
	//		'last_views': '%5B%228097157-1750313230%22%5D',
	//		'session_token': 'a6dcc0ad2b71f71cnIwaObrD3OJpbmc8MV1zxmkmxfN5_Y69qjgbd2tkJZty0TzO9er3iMhdlojE04X4V_Y8k7yD0iAQycHQ3-NpvtzqJzBpZrleP4KEajZjWFiDg1oosnzfKCKpGH64qBjIT5L2zOu4GCkpfrOBPvJh0mz1MpZz_ImssX699IctA02YyLsVE_L6pTfSffaGRP0fxnblX8acA9XTr8jNg-NjKstf54PNFuGnR2NtiAGRpTwTYioWQA40zYRcZqRetdkwJ5IjSaPD9p-NCA-vJVC53Q%3D%3D',
	//}
	cookies := []*http.Cookie{
		{Name: "dscld", Value: "true"},
		{Name: "session_ath", Value: "black"},
		{Name: "last_views", Value: "%5B%228097157-1750313230%22%5D"},
		{Name: "session_token", Value: "a6dcc0ad2b71f71cnIwaObrD3OJpbmc8MV1zxmkmxfN5_Y69qjgbd2tkJZty0TzO9er3iMhdlojE04X4V_Y8k7yD0iAQycHQ3-NpvtzqJzBpZrleP4KEajZjWFiDg1oosnzfKCKpGH64qBjIT5L2zOu4GCkpfrOBPvJh0mz1MpZz_ImssX699IctA02YyLsVE_L6pTfSffaGRP0fxnblX8acA9XTr8jNg-NjKstf54PNFuGnR2NtiAGRpTwTYioWQA40zYRcZqRetdkwJ5IjSaPD9p-NCA-vJVC53Q%3D%3D"},
	}

	request.SetCookies(cookies...)

	resp, err := request.Get(url)
	if err != nil {
		return "", err
	}

	text := resp.String()

	//	使用正则提取 video_url "html5player.setVideoUrlHigh(' .+? ');
	re := `html5player.setVideoUrlHigh\('(.+?)'\);`
	reg := regexp.MustCompile(re)
	match := reg.FindStringSubmatch(text)

	if len(match) > 1 {
		println(match[1])
		return match[1], nil
	} else {
		return "", nil
	}

}
