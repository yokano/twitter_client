// リクエスト URL のパターンに該当する処理をする。
// 必要に応じて Model と View を使って処理を進める。
package controller

import (
	"net/http"
)

// コントローラオブジェクト
type Controller struct {
}

// Controller オブジェクトを作成して返す
func NewController() *Controller {
	controller := new(Controller)
	return controller
}

// メソッドを http.HandleFunc() で呼び出し可能な関数型に変換して返す。
// http.HandleFunc() は引数として func(http.ResponseWriter, *http.Request) 型の関数しか渡せない。
// コントローラのメソッドをこの関数型で包むことで http.HandleFunc() から呼び出し可能にする。
func (this *Controller) GetHandler(callback func(this *Controller, w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		callback(this, w, r)
	}
}

// リクエスト URL に合わせて処理を振り分ける
func (this *Controller) Handle() {
	table := make(map[string]func(this *Controller, w http.ResponseWriter, r *http.Request), 0)
	
	// page
	table["/"] = (*Controller).Top
	
	// oauth
	table["/login_twitter"] = (*Controller).LoginTwitter
	table["/callback"] = (*Controller).CallbackTwitter
	
	// api
	table["/get_timeline"] = (*Controller).GetTimeline
	table["/get_account"] = (*Controller).GetAccount
	table["/logout"] = (*Controller).Logout
	
	for url, callback := range table {
		http.HandleFunc(url, this.GetHandler(callback))
	}
}
