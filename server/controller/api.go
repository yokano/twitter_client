package controller

import (
	"net/http"
	"fmt"
	"appengine"
	. "server/model"
	. "server/lib"
	config "server/config"
)

// アカウントデータの取得
func (this *Controller) GetTimeline(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	model := NewModel(c)
	
	sessionId := this.GetSession(c, r)
	oauthDatas := model.GetDatasFromSession(sessionId)
	
	params := make(map[string]string, 0)
	params["screen_name"] = r.FormValue("screen_name")
	params["oauth_token"] = oauthDatas["token"]

	oauth := NewOAuth1(c, "http://okatter-client.appspot.com/callback")
	timeline := oauth.Request("GET", "https://api.twitter.com/1.1/statuses/user_timeline.json", params, "", []string{config.TWITTER_CONSUMER_SECRET, oauthDatas["secret"]})
	
	c.Debugf("timeline: %s", timeline)
	
	fmt.Fprintf(w, timeline)
}

// タイムライン読み出し
func (this *Controller) GetAccount(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	model := NewModel(c)
	
	sessionId := this.GetSession(c, r)
	oauthDatas := model.GetDatasFromSession(sessionId)
	
	params := make(map[string]string, 0)
	params["oauth_token"] = oauthDatas["token"]
	
	oauth := NewOAuth1(c, "http://okatter-client.appspot.com/callback")
	account := oauth.Request("GET", "https://api.twitter.com/1.1/account/verify_credentials.json", params, "", []string{config.TWITTER_CONSUMER_SECRET, oauthDatas["secret"]})
	fmt.Fprintf(w, account)
}

// ログアウト
func (this *Controller) Logout(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	
	sessionId := this.GetSession(c, r)
	this.CloseSession(c, sessionId, w, r)
	
	http.Redirect(w, r, "/", 302)
}