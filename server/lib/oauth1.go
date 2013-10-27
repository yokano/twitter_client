/**
 * Twitterとの通信
 * OAuth 1.0 Revision A を使う
 * @author y.okano
 * @file
 */
package lib

import (
	"encoding/base64"
	"strings"
	"strconv"
	"time"
	"fmt"
	"net/url"
	"appengine"
	"net/http"
	"crypto/hmac"
	"crypto/sha1"
	"sort"
	config "server/config"
)

/**
 * OAuth1.0aの通信を行うクラス
 * @class
 * @param {map[string]string} params oauthパラメータの配列
 * @param {appengine.Context} context コンテキスト
 */
type OAuth1 struct {
	params map[string]string
	context appengine.Context
}

/**
 * OAuthクラスのインスタンス化
 * @function
 * @params {appengine.Context} c コンテキスト
 * @params{appengine.Context} callback コールバックURL
 * @returns {*OAuth} OAuthインスタンス
 */
func NewOAuth1(c appengine.Context, callback string) *OAuth1 {
	params := make(map[string]string, 10)
	params["oauth_callback"] = callback
	params["oauth_consumer_key"] = config.TWITTER_CONSUMER_KEY
	params["oauth_signature_method"] = "HMAC-SHA1"
	params["oauth_version"] = "1.0"
	
	oauth := new(OAuth1)
	oauth.params = params
	oauth.context = c
	return oauth
}

/**
 * Twitter へリクエストトークンを要求する
 * @method
 * @memberof OAuth1
 * @param {string} targetUrl リクエスト要求先のURL
 * @returns {map[string]string} リクエスト結果
 */
func (this *OAuth1) RequestToken(targetUrl string) map[string]string {
	response := this.Request(targetUrl, "")
	datas := strings.Split(response, "&")
	result := make(map[string]string, len(datas))
	for i := 0; i < len(datas); i++ {
		data := strings.Split(datas[i], "=")
		result[data[0]] = data[1]
	}
	
	return result
}

/**
 * リクエストを送信してレスポンスを受信する
 * メソッドは POST 固定
 * @method
 * @memberof OAuth1
 * @param {string} targetUrl 送信先
 * @param {string} body リクエストボディ
 * @returns {string} レスポンス
 */
func (this *OAuth1) Request(targetUrl string, body string) string {
	// 認証用にパラメータを追加
	this.context.Debugf("body: %#v", body)
	p := strings.Split(body, "=")
	this.context.Debugf("p: %#v", p)
	if body != "" {
		this.params[p[0]] = p[1]
	}

	// リクエストごとに変わるパラメータを設定
	this.params["oauth_nonce"] = this.CreateNonce()
	this.params["oauth_timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	this.params["oauth_signature"] = this.CreateSignature(targetUrl)
	
	// リクエスト送信
	params := make(map[string]string, 1)
	params["Authorization"] = this.CreateHeader()
	response := Request(this.context, "POST", targetUrl, params, body)
	
	// レスポンスボディの読み取り
	result := make([]byte, 2048)
	response.Body.Read(result)
	
	return string(result)
}

/**
 * oauth_nonce を作成する
 * @method
 * @memberof OAuth1
 * @returns {string} 作成したoauth_nonce
 */
func (this *OAuth1) CreateNonce() string {
	nonce := ""
	for i := 0; i < 4; i++ {
		nonce = strings.Join([]string{nonce, string(GetRandomizedString())}, "")
	}
	this.context.Infof("nonce: %s", nonce)
	return nonce
}

/**
 * Aouthorization ヘッダを作成する
 * @method
 * @memberof OAuth1
 * @returns {string} ヘッダ
 */
func (this *OAuth1) CreateHeader() string {
	params := make([]string, 0)
	for key, val := range this.params {
		if key == "screen_name" {
			continue
		}
		key = url.QueryEscape(key)
		val = url.QueryEscape(val)
		set := fmt.Sprintf(`%s="%s"`, key, val)
		params = append(params, set)
	}
	header := strings.Join(params, ", ")
	header = fmt.Sprintf("OAuth %s", header)
	return header
}

/**
 * oauth_signature を作成する
 * @method
 * @memberof OAuth1
 * @param {string} targetUrl リクエスト送信先のURL
 * @returns {string} oauth_signature
 */
func (this *OAuth1) CreateSignature(targetUrl string) string {
	keys := make([]string, 0)
	for key := range this.params {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	
	params := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := this.params[key]
		params[i] = fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(val))
		this.context.Debugf("Signature:%#v", params[i])
	}
	paramString := strings.Join(params, "&")
	baseString := fmt.Sprintf("POST&%s&%s", url.QueryEscape(targetUrl), url.QueryEscape(paramString))
	
	signatureKey := fmt.Sprintf("%s&", url.QueryEscape(config.TWITTER_CONSUMER_SECRET))
	hash := hmac.New(sha1.New, []byte(signatureKey))
	hash.Write([]byte(baseString))
	signature := hash.Sum(nil)
	return base64.StdEncoding.EncodeToString(signature)
}

/**
 * 認証ページヘリダイレクトする
 * @memberof OAuth1
 * @method
 * @param {http.ResponseWriter} w 応答先
 * @param {*http.Request} r リクエスト
 * @param {string} targetUrl リダイレクト先
 * @param {string} token 未認証リクエストトークン
 */
func (this *OAuth1) Authenticate(w http.ResponseWriter, r *http.Request, targetUrl string, token string) {
	to := fmt.Sprintf("?oauth_token=%s", token)
	to = strings.Join([]string{targetUrl, to}, "")
	http.Redirect(w, r, to, 302)
}

/**
 * リクエストトークンをアクセストークンに変換する
 * @memberof OAuth1
 * @method
 * @param {string} token リクエストトークン
 * @param {string} verifier 認証データ
 * @param {string} targetUrl リクエストの送信先
 * @returns {map[string]string} アクセストークンとユーザデータ
 */
func (this *OAuth1) ExchangeToken(token string, verifier string, targetUrl string) map[string]string {
	this.params["oauth_token"] = token
	body := fmt.Sprintf("oauth_verifier=%s", verifier)
	response := this.Request(targetUrl, body)
	
	datas := strings.Split(response, "&")
	result := make(map[string]string, len(datas))
	for i := 0; i < len(datas); i++ {
		data := strings.Split(datas[i], "=")
		result[data[0]] = data[1]
	}
	return result
}
