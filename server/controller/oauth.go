package controller

import (
	"net/http"
	"appengine"
	"fmt"
	"strings"
	. "server/lib"
	. "server/view"
	config "server/config"
)

// Twitter ボタンが押された時の処理
func (this *Controller) LoginTwitter(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	oauth := NewOAuth1(c, fmt.Sprintf("http://okatter-client.appspot.com/callback"))
	result := oauth.RequestToken("https://api.twitter.com/oauth/request_token")
	oauth.Authenticate(w, r, "https://api.twitter.com/oauth/authenticate", result["oauth_token"])
}

// Twitter でユーザが許可をした時に呼び出されるコールバック関数
func (this *Controller) CallbackTwitter(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	token := r.FormValue("oauth_token")
	verifier := r.FormValue("oauth_verifier")
	
	oauth := NewOAuth1(c, "http://okatter-client.appspot.com/callback")
	result := oauth.ExchangeToken(token, verifier, "https://api.twitter.com/oauth/access_token")
	
	view := NewView(c, w)
	
	if result["oauth_token"] != "" {
		result["screen_name"] = strings.Trim(result["screen_name"], "\x00")

		// ログイン成功したらタイムライン読み出し
		params := make(map[string]string, 0)
		params["screen_name"] = result["screen_name"]
		params["oauth_token"] = result["oauth_token"]
		toURL := "https://api.twitter.com/1.1/statuses/user_timeline.json"
		oauth.Request("GET", toURL, params, "", []string{config.TWITTER_CONSUMER_SECRET, result["oauth_token_secret"]})

		view.TimeLine()
	} else {
		// ログイン失敗
		view.Login()
		fmt.Fprintf(w, "ログインに失敗しました")
	}
}
