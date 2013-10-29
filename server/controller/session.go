package controller

import (
	"net/http"
	"appengine"
	. "server/model"
	. "server/config"
	. "server/lib"
)

// 引数として渡されたユーザキーに該当するユーザのセッションを開始する。
// oauth_token, oauth_token_secret に関連付いたセッションIDを生成して memcache と cookie に保存する。
// ブラウザが提供する cookie に保存されたセッションIDを memcache に保存されたセッションIDを比較して認証する。
func (this *Controller) StartSession(w http.ResponseWriter, r *http.Request, token string, secret string) {
	c := appengine.NewContext(r)
	model := NewModel(c)
	sessionId := model.StartSession(token, secret)
	cookie := NewCookie("okatter", sessionId, HOSTNAME, "/", 1)
	http.SetCookie(w, cookie)
}

// ブラウザの Cookie に保存されているセッションIDを取得する。
// セッションが存在しない場合は空文字を返す。
func (this *Controller) GetSession(c appengine.Context, r *http.Request) string {
	var result string
	cookie, err := r.Cookie("okatter")
	if err == http.ErrNoCookie {
		result = ""
	} else if err != nil {
		result = ""
		c.Errorf(err.Error())
	} else {
		result = cookie.Value
	}
	return result
}

// セッションを終了する。
func (this *Controller) CloseSession(c appengine.Context, sessionId string, w http.ResponseWriter, r *http.Request) {
	model := NewModel(c)
	model.RemoveSession(sessionId)
	this.DeleteCookie(c, w)
}

// セッション ID が保存されたクッキーを削除する
func (this *Controller) DeleteCookie(c appengine.Context, w http.ResponseWriter) {
	cookie := NewCookie("okatter", "", HOSTNAME, "/", -1)
	http.SetCookie(w, cookie)
}
