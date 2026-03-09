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

	//header
	//headers = {
	//	'accept': '*/*',
	//		'accept-language': 'en-US,en;q=0.9,en-GB;q=0.8,zh-CN;q=0.7,zh;q=0.6',
	//		'authorization': 'Bearer AAAAAAAAAAAAAAAAAAAAANRILgAAAAAAnNwIzUejRCOuH5E6I8xnZz4puTs%3D1Zv7ttfk8LF81IUq16cHjhLTvJu4FA33AGWWjCpTnA',
	//		'content-type': 'application/json',
	//		'priority': 'u=1, i',
	//		'referer': 'https://x.com/gcjpzx123/status/2029552095544901809',
	//		'sec-ch-ua': '"Not:A-Brand";v="99", "Google Chrome";v="145", "Chromium";v="145"',
	//		'sec-ch-ua-mobile': '?0',
	//		'sec-ch-ua-platform': '"macOS"',
	//		'sec-fetch-dest': 'empty',
	//		'sec-fetch-mode': 'cors',
	//		'sec-fetch-site': 'same-origin',
	//		'user-agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/145.0.0.0 Safari/537.36',
	//		'x-client-transaction-id': 'l2pt0KOfnkaN2s1ZjigfJVVnc0AK7rTubwX2dK5QBXMQyUn4Nw+6qPmMou1m5j6dm2kRyJJy8G//CdK9mRdKvF0O7LzylA',
	//		'x-csrf-token': 'bb67e64f4675e7592885bbacac6e7cbaaf8ce6155ba4f7683d2d0e7a5581877768ff150ba8fd4c6ff9550cf55d9b20ec4290be531a397980b20b066f000cb6b3b32309c8a53bbe50806604ab05ac8d7a',
	//		'x-twitter-active-user': 'yes',
	//		'x-twitter-auth-type': 'OAuth2Session',
	//		'x-twitter-client-language': 'zh-cn',
	//	# 'cookie': 'guest_id_marketing=v1%3A177306589292335312; guest_id_ads=v1%3A177306589292335312; guest_id=v1%3A177306589292335312; personalization_id="v1_ZhYTF4+MKum/bZbBRWtX5A=="; gt=2031011677500834269; __cuid=de3cc0d2817746b0a8e21cb49fc6ab23; g_state={"i_l":0,"i_ll":1773065906025,"i_b":"Zislfis2CC4+7INXL1IUHBOrofYRv0+Rjarssp+mC9I","i_e":{"enable_itp_optimization":0}}; kdt=44Nsxf6CQrk1YEIzsn8KGADuBe95GzFvKWFrNwr1; auth_token=b8528b5c097786837ef3958cfcbf3255101feb02; ct0=bb67e64f4675e7592885bbacac6e7cbaaf8ce6155ba4f7683d2d0e7a5581877768ff150ba8fd4c6ff9550cf55d9b20ec4290be531a397980b20b066f000cb6b3b32309c8a53bbe50806604ab05ac8d7a; att=1-y3LCfr4c6UEAtz1bucHTcEkpFgorQDUs249gMDhm; twid=u%3D1850829389724033024; __cf_bm=3hFRqD4PSTcpWmzHAXmO0jcbC.0PDGd8C2it8IAd0ZU-1773070861.7285812-1.0.1.1-1jz.G7U1FwcBZRLplmagqO0CYx2IMlT93AmfC0.enQPqUAb3yMQbZfYe1uo7TScdzifm1Mf0HnjV4NlO1Ojd9b6yGOvkGIx30rgTG6ezMojG5U7.2CnrQk_t5sw5rDli; lang=zh-CN',
	//}
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
		"x-client-transaction-id":   "l2pt0KOfnkaN2s1ZjigfJVVnc0AK7rTubwX2dK5QBXMQyUn4Nw+6qPmMou1m5j6dm2kRyJJy8G//CdK9mRdKvF0O7LzylA",
		"x-csrf-token":              "bb67e64f4675e7592885bbacac6e7cbaaf8ce6155ba4f7683d2d0e7a5581877768ff150ba8fd4c6ff9550cf55d9b20ec4290be531a397980b20b066f000cb6b3b32309c8a53bbe50806604ab05ac8d7a",
		"x-twitter-active-user":     "yes",
		"x-twitter-auth-type":       "OAuth2Session",
		"x-twitter-client-language": "zh-cn",
	})

	cookies := []*http.Cookie{
		{Name: "guest_id", Value: "v1%3A177306589292335312"},
		{Name: "guest_id_ads", Value: "v1%3A177306589292335312"},
		{Name: "guest_id_marketing", Value: "v1%3A177306589292335312"},
		{Name: "__cf_bm", Value: "3hFRqD4PSTcpWmzHAXmO0jcbC.0PDGd8C2it8IAd0ZU-1773070861.7285812-1.0.1.1-1jz.G7U1FwcBZRLplmagqO0CYx2IMlT93AmfC0.enQPqUAb3yMQbZfYe1uo7TScdzifm1Mf0HnjV4NlO1Ojd9b6yGOvkGIx30rgTG6ezMojG5U7.2CnrQk_t5sw5rDli"},
		{Name: "personalization_id", Value: "v1_ZhYTF4+MKum/bZbBRWtX5A=="},
		{Name: "gt", Value: "2031011677500834269"},
		{Name: "kdt", Value: "44Nsxf6CQrk1YEIzsn8KGADuBe95GzFvKWFrNwr1"},
		{Name: "auth_token", Value: "b8528b5c097786837ef3958cfcbf3255101feb02"},
		{Name: "ct0", Value: "bb67e64f4675e7592885bbacac6e7cbaaf8ce6155ba4f7683d2d0e7a5581877768ff150ba8fd4c6ff9550cf55d9b20ec4290be531a397980b20b066f000cb6b3b32309c8a53bbe50806604ab05ac8d7a"},
		{Name: "att", Value: "1-y3LCfr4c6UEAtz1bucHTcEkpFgorQDUs249gMDhm"},
		{Name: "lang", Value: "zh-CN"},
		{Name: "twid", Value: "u%3D1850829389724033024"},
	}

	request.SetCookies(cookies...)

	var data = map[string]string{
		"variables":    fmt.Sprintf(`{"focalTweetId":"%s","with_rux_injections":false,"rankingMode":"Relevance","includePromotedContent":true,"withCommunity":true,"withQuickPromoteEligibilityTweetFields":true,"withBirdwatchNotes":true,"withVoice":true}`, id),
		"features":     `{"rweb_video_screen_enabled":false,"payments_enabled":false,"profile_label_improvements_pcf_label_in_post_enabled":true,"rweb_tipjar_consumption_enabled":true,"verified_phone_label_enabled":false,"creator_subscriptions_tweet_preview_api_enabled":true,"responsive_web_graphql_timeline_navigation_enabled":true,"responsive_web_graphql_skip_user_profile_image_extensions_enabled":false,"premium_content_api_read_enabled":false,"communities_web_enable_tweet_community_results_fetch":true,"c9s_tweet_anatomy_moderator_badge_enabled":true,"responsive_web_grok_analyze_button_fetch_trends_enabled":false,"responsive_web_grok_analyze_post_followups_enabled":true,"responsive_web_jetfuel_frame":false,"responsive_web_grok_share_attachment_enabled":true,"articles_preview_enabled":true,"responsive_web_edit_tweet_api_enabled":true,"graphql_is_translatable_rweb_tweet_is_translatable_enabled":true,"view_counts_everywhere_api_enabled":true,"longform_notetweets_consumption_enabled":true,"responsive_web_twitter_article_tweet_consumption_enabled":true,"tweet_awards_web_tipping_enabled":false,"responsive_web_grok_show_grok_translated_post":false,"responsive_web_grok_analysis_button_from_backend":true,"creator_subscriptions_quote_tweet_preview_enabled":false,"freedom_of_speech_not_reach_fetch_enabled":true,"standardized_nudges_misinfo":true,"tweet_with_visibility_results_prefer_gql_limited_actions_policy_enabled":true,"longform_notetweets_rich_text_read_enabled":true,"longform_notetweets_inline_media_enabled":true,"responsive_web_grok_image_annotation_enabled":true,"responsive_web_enhance_cards_enabled":false}`,
		"fieldToggles": `{"withArticleRichContentState":true,"withArticlePlainText":false,"withGrokAnalyze":false,"withDisallowedReplyControls":false}`,
	}
	targetUrl := "https://x.com/i/api/graphql/vsCTCQrF8oqASUb-x2SBcg/TweetDetail"
	// 要把data，变成form表单的形式然后get出去
	//request.SetFormData(data)
	request.SetQueryParams(data)

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
