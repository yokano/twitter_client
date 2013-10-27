// 画面表示全般を行う。必要に応じてモデルからデータを取り出し整形する。
// HTML ファイルを出力するときは html ディレクトリからテンプレートを取り出して使う。
package view

import(
	"net/http"
	"html/template"
	"appengine"
	. "server/lib"
)

// 画面表示を行うオブジェクト
type View struct {
	c appengine.Context
	w http.ResponseWriter
}

// View オブジェクトを作成する。
func NewView(c appengine.Context, w http.ResponseWriter) *View {
	view := new(View)
	view.c = c
	view.w = w
	return view
}

// ログイン画面を表示する
func (this *View) Login() {
	t, err := template.ParseFiles("server/view/html/login.html")
	Check(this.c, err)
	t.Execute(this.w, nil)
}

// タイムライン画面を表示する
func (this *View) TimeLine() {
	t, err := template.ParseFiles("server/view/html/timeline.html")
	Check(this.c, err)
	t.Execute(this.w, nil)
}
