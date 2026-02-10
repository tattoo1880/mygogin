package myutils

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/imroc/req/v3"
	"github.com/tidwall/gjson"
)

type GetListUtils interface {
	GetList(id string) string
}

type getListUtilsImpl struct{}

func NewGetListUtils() GetListUtils {
	return &getListUtilsImpl{}
}

func handlerError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func (g *getListUtilsImpl) GetList(id string) string {
	client := req.C()
	request := client.R()

	// 设置请求头

	request.SetHeaders(map[string]string{
		"accept":                    "*/*",
		"accept-language":           "zh-CN,zh;q=0.9",
		"authorization":             "Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA",
		"content-type":              "application/json",
		"priority":                  "u=1, i",
		"referer":                   "https://x.com/Kunluntalk/status/1812256811199922205",
		"sec-ch-ua":                 `"Not/A)Brand";v="8", "Chromium";v="126", "Google Chrome";v="126"`,
		"sec-ch-ua-mobile":          "?0",
		"sec-ch-ua-platform":        "macOS",
		"sec-fetch-dest":            "empty",
		"sec-fetch-mode":            "cors",
		"sec-fetch-site":            "same-origin",
		"user-agent":                "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36",
		"x-client-transaction-id":   "g0u4vt+Ko5s391U5LENI0nRzSYAEBA0FshF7ipCTFWfM6TyLegvRo0sSKwNT5QA6TNIXuIbmM0K0tipwpOhaVXyUGBN5gA",
		"x-csrf-token":              "9d2cd6b8f29f24503bcf8a0677a02a46d61d4f6c34e522268d5764bfdb3f38d79f69b24c8637f1e3f7909b0b093842e8d3c57eedfab08ccff586f0e991145fbd3698395aa576180ce81b6cd5063c41e3",
		"x-twitter-active-user":     "yes",
		"x-twitter-auth-type":       "OAuth2Session",
		"x-twitter-client-language": "zh-cn",
		"x-xp-forwarded-for":        "ff261f1be977c77bf684384b042e08be8dd2289e34dd88c7ca2d55f2d67581c406580ff0e4c22fad8e254d17238830fa0fb7f24fa78523f35e338d0a994df9419cf36a468f37b48463a985ed83f6425077f62d14fd80a3c6fb643e6cacf7da08d97d2bc8d9f62ce7ef8f6e9276042e766926acaee1525d8182fc73af1560cd0787522279fcc165fde3a46ff61a09ea8d957c9a60fc677b40fbff13abb4dda6b4afb5fec0cf017d87aa91686f01b910bc7581f664dccfd878781317834b81c12ef896772a4e4210c2e477ccebdf99ddcc52759fb10d69ae0e53ac5bc02df243c4654136e0bb98f34d5911c05357a3ce8c4f89bb8b1d8d02a770b9521f87925ac9c2",
	})

	// 设置Cookies
	cookies := []*http.Cookie{
		{Name: "guest_id", Value: "v1%3A177071483672131597"},
		{Name: "guest_id_ads", Value: "v1%3A174990001013364461"},
		{Name: "guest_id_marketing", Value: "v1%3A174990001013364461"},
		{Name: "__cf_bm", Value: "1uek3OYTtepgGLcUMBaLHl62ncNu6pBMPPN8vFpAJWQ-1749900010-1.0.1.1-0g4_5xyhLEzXxMIVj9PN0oygPfaydhwkhA7IicacfC2dQeUSNkpnruU.kbZ9IlxrUf0pWSXbLSfZkCAkHfs4GKloQ0EHVUXiANndlbDJqk8"},
		{Name: "personalization_id", Value: "v1_H6K4LwMccNB+x5dN/aXIXg=="},
		{Name: "gt", Value: "1933846922655334517"},
		{Name: "kdt", Value: "c3CfYlYOf9Q7Wnm5WLWNCimzFE7W31KemKWBHtlJ"},
		{Name: "auth_token", Value: "c6d4545e0a6e5bf0455f92eb29f03a87b6c40eaf"},
		{Name: "ct0", Value: "9d2cd6b8f29f24503bcf8a0677a02a46d61d4f6c34e522268d5764bfdb3f38d79f69b24c8637f1e3f7909b0b093842e8d3c57eedfab08ccff586f0e991145fbd3698395aa576180ce81b6cd5063c41e3"},
		{Name: "att", Value: "1-yDmgHZVUZCcDKf4gGI7b9RLWQBtkYPxNzzrp3Gdj"},
		{Name: "lang", Value: "en"},
		{Name: "twid", Value: "u%3D1850829389724033024"},
	}

	request.SetCookies(cookies...)

	var data = map[string]string{
		"variables":    fmt.Sprintf(`{"focalTweetId":"%s","with_rux_injections":false,"rankingMode":"Relevance","includePromotedContent":true,"withCommunity":true,"withQuickPromoteEligibilityTweetFields":true,"withBirdwatchNotes":true,"withVoice":true}`, id),
		"features":     `{"rweb_video_screen_enabled":false,"payments_enabled":false,"profile_label_improvements_pcf_label_in_post_enabled":true,"rweb_tipjar_consumption_enabled":true,"verified_phone_label_enabled":false,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"premium_content_api_read_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"responsive_web_grok_analyze_button_fetch_trends_enabled":false,"responsive_web_grok_analyze_post_followups_enabled":true,"responsive_web_jetfuel_frame":false,"responsive_web_grok_share_attachment_enabled":true,"articles_preview_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"responsive_web_grok_show_grok_translated_post":false,"responsive_web_grok_analysis_button_from_backend":true,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"responsive_web_grok_image_annotation_enabled":true,"responsive_web_enhance_cards_enabled":false}`,
		"fieldToggles": `{"withArticleRichContentState":true,"withArticlePlainText":false,"withGrokAnalyze":false,"withDisallowedReplyControls":false}`,
	}
	targetUrl := "https://x.com/i/api/graphql/8IPrg-fiWPM4p735QRfGqA/TweetDetail"
	// 要把data，变成form表单的形式然后get出去
	request.SetFormData(data)

	resp, err := request.Get(targetUrl)
	if err != nil {
		handlerError(err)
	}

	fmt.Println("请求成功，数据为：", resp.String())

	// ! todo 使用正则将所有的url提取出来
	var resultList []string
	re := regexp.MustCompile(`"url"\s*:\s*"([^"]+\.mp4[^"]*)"`)

	matches := re.FindAllStringSubmatch(resp.String(), -1)
	for _, match := range matches {
		if len(match) > 1 {
			resultList = append(resultList, match[1])
		}
	}

	if len(resultList) == 0 {
		// var result map[string]interface{}
		// err = resp.Unmarshal(&result)
		// if err != nil {
		// 	handlerError(err)
		// }
		//
		// log.Println("没有匹配到视频链接，尝试从其他字段获取")
		//
		// obj1 := result["data"].(map[string]interface{})["threaded_conversation_with_injections_v2"].(map[string]interface{})["instructions"].([]interface{})[0].(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["itemContent"].(map[string]interface{})["tweet_results"].(map[string]interface{})["result"].(map[string]interface{})["card"].(map[string]interface{})["legacy"].(map[string]interface{})["binding_values"].([]interface{})[0].(map[string]interface{})["value"].(map[string]interface{})["string_value"]
		// // ! 打印obj1的类型
		// fmt.Printf("obj1 type: %T\n", obj1)
		// var jsonData map[string]interface{}
		// err = json.Unmarshal([]byte(obj1.(string)), &jsonData)
		// if err != nil {
		// 	handlerError(err)
		// }
		// targetid := jsonData["component_objects"].(map[string]interface{})["media_1"].(map[string]interface{})["data"].(map[string]interface{})["id"].(string)
		// targetUrllist := jsonData["media_entities"].(map[string]interface{})[targetid].(map[string]interface{})["video_info"].(map[string]interface{})["variants"]
		// cardlist := []string{}
		// for _, v := range targetUrllist.([]interface{}) {
		// 	if v.(map[string]interface{})["content_type"].(string) == "video/mp4" {
		// 		cardlist = append(cardlist, v.(map[string]interface{})["url"].(string))
		// 	}
		// }
		// return cardlist

		// 如果没有匹配到视频链接，尝试从其他字段获取

		var results []gjson.Result
		jsonData := resp.String()
		result := gjson.Parse(jsonData)
		findAllByKey(result, "string_value", &results)
		if len(results) == 0 {
			log.Println("没有匹配到视频链接")
			return ""
		}
		fmt.Println("找到了string_value字段，数量为：", len(results))
		for _, res := range results {
			var innerResults []gjson.Result
			innerResult := gjson.Parse(res.String())
			findAllByKey(innerResult, "url", &innerResults)
			for _, innerRes := range innerResults {
				if matched, _ := regexp.MatchString(`\.mp4`, innerRes.String()); matched {
					resultList = append(resultList, innerRes.String())
				}
			}
			fmt.Println("当前string_value字段中找到的url数量为：", len(innerResults))
			fmt.Println(innerResults)
			if len(innerResults) != 0 {
				// ! 将结果返回
				for _, innerRes := range innerResults {
					if matched, _ := regexp.MatchString(`\.mp4`, innerRes.String()); matched {
						resultList = append(resultList, innerRes.String())
					}
				}
				break
			}
		}

	}

	bestUrl := getBestQualityURL(resultList)
	fmt.Println(bestUrl)
	fmt.Println(bestUrl)
	fmt.Println(bestUrl)
	fmt.Println(bestUrl)
	fmt.Println(bestUrl)
	fmt.Println(bestUrl)
	fmt.Println(bestUrl)
	return bestUrl
}

func findAllByKey(result gjson.Result, key string, out *[]gjson.Result) {
	if result.IsObject() {
		result.ForEach(func(k, v gjson.Result) bool {
			if k.String() == key {
				*out = append(*out, v)
			}
			findAllByKey(v, key, out)
			return true
		})
	} else if result.IsArray() {
		result.ForEach(func(_, v gjson.Result) bool {
			findAllByKey(v, key, out)
			return true
		})
	}
}

func getBestQualityURL(urls []string) string {
	re := regexp.MustCompile(`(\d{2,4})x(\d{2,4})`)
	maxPixels := 0
	bestURL := ""
	for _, url := range urls {
		match := re.FindStringSubmatch(url)
		if len(match) == 3 {
			w, _ := strconv.Atoi(match[1])
			h, _ := strconv.Atoi(match[2])
			pixels := w * h
			if pixels > maxPixels {
				maxPixels = pixels
				bestURL = url
			}
		}
	}
	return bestURL
}
