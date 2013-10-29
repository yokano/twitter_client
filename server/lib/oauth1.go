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
	oauth := new(OAuth1)
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
	params := make(map[string]string, 0)
	params["oauth_callback"] = config.TWITTER_CALLBACK_URL
	response := this.Request("POST", targetUrl, make(map[string]string), "", []string{config.TWITTER_CONSUMER_SECRET, ""})
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
 * @param {string} method POSTかGET
 * @param {string} targetUrl 送信先
 * @param {string} params パラメータ
 * @param {string} body リクエストボディ
 * @param {string} secret 暗号鍵（ConsumerSecret と OAuth Token Secret）
 * @returns {string} レスポンス
 */
func (this *OAuth1) Request(method string, targetUrl string, params map[string]string, body string, secret []string) string {
	oauthParams := make(map[string]string, 0)
	
	for key, val := range params {
		oauthParams[key] = val
	}
	
	oauthParams["oauth_consumer_key"] = config.TWITTER_CONSUMER_KEY
	oauthParams["oauth_signature_method"] = "HMAC-SHA1"
	oauthParams["oauth_version"] = "1.0"
	oauthParams["oauth_nonce"] = this.CreateNonce()
	oauthParams["oauth_timestamp"] = strconv.Itoa(int(time.Now().Unix()))
	oauthParams["oauth_signature"] = this.CreateSignature(method, targetUrl, oauthParams, secret)
	
	for key, val := range params {
		oauthParams[key] = val
	}
	
	// リクエスト送信
	httpParams := make(map[string]string, 1)
	httpParams["Authorization"] = this.CreateHeader(oauthParams)

	var response *http.Response
	if method == "GET" {
		query := make(map[string]string, 0)
		for key, val := range params {
			if key != "oauth_token" {
				query[key] = val
			}
		}
		response = Get(this.context, targetUrl, query, httpParams)
	} else {
		response = Request(this.context, method, targetUrl, httpParams, body)
	}
	
	// レスポンスボディの読み取り
	result := make([]byte, 1024 * 1024)
	response.Body.Read(result)
	resultString := string(result)
	resultString = strings.Trim(resultString, "\x00")
	
	return resultString
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
	return nonce
}

/**
 * Aouthorization ヘッダを作成する
 * @method
 * @memberof OAuth1
 * @param {map[string]string} oauthParams OAuthパラメータ
 * @returns {string} ヘッダ
 */
func (this *OAuth1) CreateHeader(oauthParams map[string]string) string {
	headerParams := make([]string, 0)
	for key, val := range oauthParams {
		if key == "screen_name" {
			continue
		}
		key = url.QueryEscape(key)
		val = url.QueryEscape(val)
		set := fmt.Sprintf(`%s="%s"`, key, val)
		headerParams = append(headerParams, set)
	}
	header := strings.Join(headerParams, ", ")
	header = fmt.Sprintf("OAuth %s", header)
	return header
}

/**
 * oauth_signature を作成する
 * @method
 * @memberof OAuth1
 * @param {string} method メソッド
 * @param {string} targetUrl リクエスト送信先のURL
 * @param {map[string]string} oauthParams パラメータ
 * @param {string} secret 暗号化鍵(Consumer Secret と OAuth Token Secret)
 * @returns {string} oauth_signature
 */
func (this *OAuth1) CreateSignature(method string, targetUrl string, oauthParams map[string]string, secret []string) string {
	keys := make([]string, 0)
	for key := range oauthParams {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	
	params := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		val := oauthParams[key]
		params[i] = fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(val))
	}
	paramString := strings.Join(params, "&")
	baseString := fmt.Sprintf("%s&%s&%s", method, url.QueryEscape(targetUrl), url.QueryEscape(paramString))
	
	signatureKey := fmt.Sprintf("%s&%s", url.QueryEscape(secret[0]), url.QueryEscape(secret[1]))
	
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
	params := make(map[string]string, 0)
	params["oauth_token"] = token
	params["oauth_token"] = params["oauth_token"]
	
	body := fmt.Sprintf("oauth_verifier=%s", verifier)
	response := this.Request("POST", targetUrl, params, body, []string{config.TWITTER_CONSUMER_SECRET, ""})
	
	datas := strings.Split(response, "&")
	result := make(map[string]string, len(datas))
	for i := 0; i < len(datas); i++ {
		data := strings.Split(datas[i], "=")
		result[data[0]] = data[1]
	}
	return result
}
