package controller

import (
	"net/http"
	"appengine"
	"fmt"
	. "server/lib"
	. "server/view"
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
		// ログイン成功
		params := make(map[string]string, 4)
		params["user_type"] = "Twitter"
		params["user_name"] = result["screen_name"]
		params["user_oauth_id"] = result["user_id"]
		params["user_pass"] = ""
		
		toURL := "https://api.twitter.com/1.1/statuses/user_timeline.json"
		c.Debugf("JOIN: %#v", Join("screen_name=", "yuta_okano"))
		result := oauth.Request(toURL, Join("screen_name=", "yuta_okano"))

		view.Login()
		fmt.Fprintf(w, result)
		fmt.Fprintf(w, Join("<br>screen_name=", params["user_name"]))
	} else {
		// ログイン失敗
		view.Login()
		fmt.Fprintf(w, "ログインに失敗しました")
	}
}
