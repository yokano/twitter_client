package model

import (
	"fmt"
	"encoding/json"
	"appengine/memcache"
	. "server/lib"
)

// セッションを開始する。サーバの memcache 内にセッションIDと oauth_token, oauth_secret の対応を保存している。
// 関数はセッションIDを返す。
func (this *Model) StartSession(token string, secret string) string {
	sessionId := ""
	for i := 0; i < 4; i++ {
		sessionId = fmt.Sprintf("%s%s", sessionId, GetRandomizedString())
	}
	
	data := make(map[string]string, 0)
	data["token"] = token
	data["secret"] = secret
	
	encodedData, err := json.Marshal(data)
	item := &memcache.Item {
		Key: sessionId,
		Value: encodedData,
	}
	err = memcache.Set(this.c, item)
	Check(this.c, err)
	
	return sessionId
}

// memcache から指定されたセッションIDのセッション情報を削除する
func (this *Model) RemoveSession(sessionId string) {
	err := memcache.Delete(this.c, sessionId)
	Check(this.c, err)
}

// memcache から指定されたセッションIDに関連づいている oauth_token, oauth_secret を返す
func (this *Model) GetDatasFromSession(sessionId string) map[string]string {
	item, err := memcache.Get(this.c, sessionId)
	Check(this.c, err)
	if err != nil {
		this.c.Warningf("セッションIDに関連付けられたユーザデータが存在しませんでした: sessionId:%s", sessionId)
	}
	data := make(map[string]string, 0)
	err = json.Unmarshal(item.Value, &data)
	Check(this.c, err)
	return data
}

// セッションのクリア
func (this *Model) ResetSession() {
	err := memcache.Flush(this.c)
	Check(this.c, err)
}