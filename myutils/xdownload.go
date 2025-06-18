package myutils

import (
	"encoding/json"
	"fmt"
	"github.com/imroc/req/v3"
	"log"
	"net/http"
)

type GetListUtils interface {
	GetList(id string) []string
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

func (g *getListUtilsImpl) GetList(id string) []string {
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
		"x-client-transaction-id":   "neq8NAn0WD0jNMClKOp1AoqDRBEMmocz5y4CH2x97E6PdtASQu/sQJ4W0PQO8M+VXodlYJ7KSRAd+4Zi+REYzzWlg+rjng",
		"x-csrf-token":              "4ae004fe95a13ee8cc45769bfc42f777ff3731f9b7de80d3af0b469cc2ac47eed91f84d7d10a2b875e12b8066d7142b5c0fd794311f1c34f98388cbfe16462e1fb661f761e2cd1f15afddd3a8112d2fe",
		"x-twitter-active-user":     "yes",
		"x-twitter-auth-type":       "OAuth2Session",
		"x-twitter-client-language": "zh-cn",
		"x-xp-forwarded-for":        "ff261f1be977c77bf684384b042e08be8dd2289e34dd88c7ca2d55f2d67581c406580ff0e4c22fad8e254d17238830fa0fb7f24fa78523f35e338d0a994df9419cf36a468f37b48463a985ed83f6425077f62d14fd80a3c6fb643e6cacf7da08d97d2bc8d9f62ce7ef8f6e9276042e766926acaee1525d8182fc73af1560cd0787522279fcc165fde3a46ff61a09ea8d957c9a60fc677b40fbff13abb4dda6b4afb5fec0cf017d87aa91686f01b910bc7581f664dccfd878781317834b81c12ef896772a4e4210c2e477ccebdf99ddcc52759fb10d69ae0e53ac5bc02df243c4654136e0bb98f34d5911c05357a3ce8c4f89bb8b1d8d02a770b9521f87925ac9c2",
	})

	// 设置Cookies
	cookies := []*http.Cookie{
		{Name: "guest_id", Value: "v1%3A174990001013364461"},
		{Name: "guest_id_ads", Value: "v1%3A174990001013364461"},
		{Name: "guest_id_marketing", Value: "v1%3A174990001013364461"},
		{Name: "__cf_bm", Value: "1uek3OYTtepgGLcUMBaLHl62ncNu6pBMPPN8vFpAJWQ-1749900010-1.0.1.1-0g4_5xyhLEzXxMIVj9PN0oygPfaydhwkhA7IicacfC2dQeUSNkpnruU.kbZ9IlxrUf0pWSXbLSfZkCAkHfs4GKloQ0EHVUXiANndlbDJqk8"},
		{Name: "personalization_id", Value: "v1_H6K4LwMccNB+x5dN/aXIXg=="},
		{Name: "gt", Value: "1933846922655334517"},
		{Name: "kdt", Value: "c3CfYlYOf9Q7Wnm5WLWNCimzFE7W31KemKWBHtlJ"},
		{Name: "auth_token", Value: "903daf0b5421b5d1932777277976e84f7282aaa5"},
		{Name: "ct0", Value: "4ae004fe95a13ee8cc45769bfc42f777ff3731f9b7de80d3af0b469cc2ac47eed91f84d7d10a2b875e12b8066d7142b5c0fd794311f1c34f98388cbfe16462e1fb661f761e2cd1f15afddd3a8112d2fe"},
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

	// fmt.Println("请求成功，数据为：", resp.String())

	var result = map[string]interface{}{}
	err = resp.Unmarshal(&result)
	if err != nil {
		handlerError(err)
	}

	list := result["data"].(map[string]interface{})["threaded_conversation_with_injections_v2"].(map[string]interface{})["instructions"].([]interface{})[0].(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["itemContent"].(map[string]interface{})["tweet_results"].(map[string]interface{})["result"].(map[string]interface{})["legacy"].(map[string]interface{})["entities"].(map[string]interface{})["media"]

	if list == nil {
		obj1 := result["data"].(map[string]interface{})["threaded_conversation_with_injections_v2"].(map[string]interface{})["instructions"].([]interface{})[0].(map[string]interface{})["entries"].([]interface{})[0].(map[string]interface{})["content"].(map[string]interface{})["itemContent"].(map[string]interface{})["tweet_results"].(map[string]interface{})["result"].(map[string]interface{})["card"].(map[string]interface{})["legacy"].(map[string]interface{})["binding_values"].([]interface{})[0].(map[string]interface{})["value"].(map[string]interface{})["string_value"]
		// ! 打印obj1的类型

		var jsonData map[string]interface{}
		err = json.Unmarshal([]byte(obj1.(string)), &jsonData)

		if err != nil {
			handlerError(err)
		}

		targetid := jsonData["component_objects"].(map[string]interface{})["media_1"].(map[string]interface{})["data"].(map[string]interface{})["id"].(string)

		targetUrllist := jsonData["media_entities"].(map[string]interface{})[targetid].(map[string]interface{})["video_info"].(map[string]interface{})["variants"]

		cardlist := []string{}
		for _, v := range targetUrllist.([]interface{}) {
			if v.(map[string]interface{})["content_type"].(string) == "video/mp4" {
				cardlist = append(cardlist, v.(map[string]interface{})["url"].(string))
			}
		}

		return cardlist

	}

	op := list.([]interface{})[0].(map[string]interface{})["video_info"].(map[string]interface{})["variants"]
	// 取出op中的所有url
	resultList := []string{}
	for _, v := range op.([]interface{}) {
		// fmt.Println(v.(map[string]interface{})["url"])
		if v.(map[string]interface{})["content_type"].(string) == "video/mp4" {
			resultList = append(resultList, v.(map[string]interface{})["url"].(string))
		}

	}
	return resultList
}
